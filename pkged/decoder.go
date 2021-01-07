package pkged

import (
	"encoding/base64"
	"errors"
	"fmt"
)

func GetText(file string) (string, error) {
	encoded, ok := fileMap[file]
	if !ok {
		return "", errors.New(fmt.Sprintf("file %s does not exist in binary", file))
	}
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
