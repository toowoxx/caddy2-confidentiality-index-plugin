ARG CADDY_VERSION=2.6.4
ARG REPOSITORY_NAME=toowoxx/caddy2-confidentiality-index-plugin
FROM caddy:${CADDY_VERSION}-builder as builder

ADD . /src/plugin

RUN xcaddy build \
    --with github.com/${REPOSITORY_NAME}=/src/plugin

FROM caddy:${CADDY_VERSION}
COPY --from=builder /usr/bin/caddy /usr/bin/caddy
