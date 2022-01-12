package confindex

import (
	"errors"
	"fmt"
)

type Confidentiality string

const (
	Internal     Confidentiality = "internal"
	Confidential Confidentiality = "confidential"
	Secret       Confidentiality = "secret"
)

func ToConfidentiality(str string) (Confidentiality, error) {
	switch str {
	case string(Internal):
		fallthrough
	case string(Confidential):
		fallthrough
	case string(Secret):
		return Confidentiality(str), nil
	default:
		return "invalid", errors.New(fmt.Sprintf("Invalid confidentiality %s", str))
	}
}

func (c Confidentiality) ToScriptKey() string {
	switch c {
	case Internal:
		return "int"
	case Confidential:
		return "restricted"
	case Secret:
		return "secret"
	default:
		return "invalid"
	}
}
