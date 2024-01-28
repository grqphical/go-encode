package encoder

import (
	"encoding/base32"
)

type Base32Encoder struct {
	Encoder
}

func (e *Base32Encoder) EncodeString(data string) string {
	encodedData := base32.StdEncoding.EncodeToString([]byte(data))
	return encodedData
}
