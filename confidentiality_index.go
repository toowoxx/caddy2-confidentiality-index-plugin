package confindex

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/pkg/errors"
	"github.com/toowoxx/caddy2-html-injection-plugin"
	"go.uber.org/zap"

	"github.com/toowoxx/caddy2-confidentiality-index-plugin/embedded"
)

var tpl *template.Template

type TemplateOptions struct {
	Confidentiality        string
	ConfidentialityDocLink string
}

func init() {
	caddy.RegisterModule(Middleware{})
	httpcaddyfile.RegisterHandlerDirective("confidentiality", parseCaddyfile)
	httpcaddyfile.RegisterGlobalOption("confidentiality_doc_link", unmarshalDocLink)

	var err error
	tpl, err = template.New("html").Parse(embedded.ConfidentialityIndexHtml)
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not parse template"))
	}
}

type Middleware struct {
	ConfidentialityStr string `json:"confidentiality"`
	Confidentiality    Confidentiality

	injectedWriter      *injection.InjectedWriter
	injectionMiddleware *injection.Middleware
	logger              *zap.Logger
}

var documentationLink string

func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.confidentiality",
		New: func() caddy.Module { return new(Middleware) },
	}
}

func (m *Middleware) Provision(ctx caddy.Context) error {
	m.logger = ctx.Logger(m)
	m.logger.Info("Provisioning confidentiality index plugin",
		zap.String("ConfidentialityStr", m.ConfidentialityStr),
		zap.String("documentationLink", documentationLink),
	)
	m.injectionMiddleware = &injection.Middleware{
		ContentType: "text/html",
		// Not necessary as we will be using a custom LineHandler
		Inject: "",
		Before: "</body>",
		Logger: m.logger,
	}
	return m.injectionMiddleware.Provision(ctx)
}

func (m *Middleware) Validate() error {
	confidentiality, err := ToConfidentiality(m.ConfidentialityStr)
	if err != nil {
		return errors.Wrap(err, "confidentiality index validation error")
	}
	m.Confidentiality = confidentiality
	return m.injectionMiddleware.Validate()
}

func (m Middleware) HandleLine(line string) (string, error) {
	if strings.Contains(line, m.injectedWriter.M.Before) {
		var strBuf strings.Builder
		err := tpl.Execute(&strBuf, TemplateOptions{
			Confidentiality:        m.Confidentiality.ToScriptKey(),
			ConfidentialityDocLink: documentationLink,
		})
		if err != nil {
			return "", errors.Wrap(err, "could not execute template")
		}
		textToInject := strBuf.String()
		textToInject = m.injectedWriter.HandleCSPForText(textToInject)

		return strings.Replace(line, m.injectedWriter.M.Before, textToInject+m.injectedWriter.M.Before, 1), nil
	}
	return line, nil
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	if m.injectionMiddleware.ShouldBypassForRequest(w, r) {
		return next.ServeHTTP(w, r)
	}
	r.Header.Set("Accept-Encoding", "identity")
	injectedWriter := injection.CreateInjectedWriter(w, r, m.injectionMiddleware)

	m.injectedWriter = injectedWriter
	injectedWriter.LineHandler = m
	err := next.ServeHTTP(injectedWriter, r)
	if err != nil {
		return err
	}
	if err := injectedWriter.Flush(); err != nil {
		return err
	}
	return nil
}

func (m *Middleware) unmarshalConfidentialityDir(d *caddyfile.Dispenser) error {
	if !d.Next() {
		return d.Err("expected token following confidentiality")
	}
	if !d.NextArg() {
		return d.Err("expected argument for confidentiality")
	}
	value := d.Val()
	if len(value) == 0 {
		return d.Err("expected confidentiality classification after 'confidentiality' token")
	}
	m.ConfidentialityStr = value
	return nil
}

func unmarshalDocLink(d *caddyfile.Dispenser, _ interface{}) (interface{}, error) {
	if !d.Next() {
		return nil, d.Err("expected token following confidentiality_doc_link")
	}
	if !d.NextArg() {
		return nil, d.Err("expected argument for confidentiality_doc_link")
	}
	value := d.Val()
	if len(value) == 0 {
		return nil, d.Err("expected string after 'confidentiality_doc_link' token")
	}
	documentationLink = value
	return value, nil
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.unmarshalConfidentialityDir(h.Dispenser)
	return m, err
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ injection.LineHandler       = (*Middleware)(nil)
)
