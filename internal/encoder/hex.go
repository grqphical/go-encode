package encoder

import (
	"fmt"
)

type HexEncoder struct {
	Encoder
}

func (e *HexEncoder) EncodeString(data string) string {
	var result string
	for _, r := range data {
		result += fmt.Sprintf("%x", r)
	}

	return result
}
