package decoder

import (
	"encoding/hex"
)

type HexDecoder struct {
	Decoder
}

func (d *HexDecoder) DecodeString(data string) (string, error) {
	decodedData, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}

	return string(decodedData), nil
}
