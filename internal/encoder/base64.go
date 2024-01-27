package encoder

import (
	"encoding/base64"
)

type Base64Encoder struct {
	Encoder
}

func (e *Base64Encoder) EncodeString(data string) string {
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	return encodedData
}
