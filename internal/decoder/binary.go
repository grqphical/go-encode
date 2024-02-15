package decoder

import (
	"strconv"
	"strings"
)

type BinaryDecoder struct {
	Decoder
}

func (d *BinaryDecoder) DecodeString(data string) (string, error) {
	var asciiChars []string

	data = strings.ReplaceAll(data, " ", "")

	for i := 0; i < len(data); i += 8 {
		binaryDigit := data[i : i+8]
		decimalVal, _ := strconv.ParseInt(binaryDigit, 2, 0)
		asciiChars = append(asciiChars, string(decimalVal))
	}

	result := ""
	for _, char := range asciiChars {
		result += char
	}

	return result, nil
}
