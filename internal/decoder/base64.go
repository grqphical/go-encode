package decoder

import (
	"encoding/base64"
)

type Base64Decoder struct {
	Decoder
}

func (d *Base64Decoder) DecodeString(data string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	return string(decodedData), nil
}
