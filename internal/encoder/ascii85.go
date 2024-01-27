package encoder

import (
	"encoding/ascii85"
)

type Ascii85Encoder struct {
	Encoder
}

func (e *Ascii85Encoder) EncodeString(data string) string {
	maxSize := ascii85.MaxEncodedLen(len(data))

	var dest []byte = make([]byte, maxSize)
	ascii85.Encode(dest, []byte(data))

	return string(dest)
}
