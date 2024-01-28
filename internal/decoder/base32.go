package decoder

import (
	"encoding/base32"
)

type Base32Decoder struct {
	Decoder
}

func (d *Base32Decoder) DecodeString(data string) (string, error) {
	decodedData, err := base32.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	return string(decodedData), nil
}
