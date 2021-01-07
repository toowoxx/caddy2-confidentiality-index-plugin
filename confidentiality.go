package confindex

import (
	"errors"
	"fmt"
)

type Confidentiality string

const (
	Internal Confidentiality = "internal"
	Restricted Confidentiality = "restricted"
	Confidential Confidentiality = "confidential"
)

func ToConfidentiality(str string) (Confidentiality, error) {
	switch str {
	case string(Internal): fallthrough
	case string(Restricted): fallthrough
	case string(Confidential):
		return Confidentiality(str), nil
	default:
		return "invalid", errors.New(fmt.Sprintf("Invalid confidentiality %s", str))
	}
}


func (c Confidentiality) ToScriptKey() string {
	switch c {
	case Internal:
		return "int"
	case Restricted:
		return "rest"
	case Confidential:
		return "conf"
	default:
		return "invalid"
	}
}
