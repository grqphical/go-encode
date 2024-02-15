package encoder

import "fmt"

type BinaryEncoder struct {
	Encoder
}

func (e *BinaryEncoder) EncodeString(data string) string {
	var result string

	for _, char := range data {
		result += fmt.Sprintf("%08b", char)
	}

	return result
}
