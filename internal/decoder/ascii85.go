package decoder

import (
	"encoding/ascii85"
)

type Ascii85Decoder struct {
	Decoder
}

func (d *Ascii85Decoder) DecodeString(data string) (string, error) {
	dest := make([]byte, len(data))

	_, _, err := ascii85.Decode(dest, []byte(data), true)
	if err != nil {
		return "", err
	}

	return string(dest), nil
}
