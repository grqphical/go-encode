package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/grqphical/go-encode/internal/decoder"
)

func TestBase64Decoder(t *testing.T) {
	decoder := decoder.Base64Decoder{}

	decodedData, _ := decoder.DecodeString("SGVsbG8sIFdvcmxkIQ==")

	assert.Equal(t, "Hello, World!", decodedData)
}

func TestBase32Decoder(t *testing.T) {
	decoder := decoder.Base32Decoder{}

	decodedData, _ := decoder.DecodeString("JBSWY3DPFQQFO33SNRSCC===")

	assert.Equal(t, "Hello, World!", decodedData)
}

func TestAscii85Decoder(t *testing.T) {
	decoder := decoder.Ascii85Decoder{}

	decodedData, _ := decoder.DecodeString("87cURD_*#4DfTZ)+T")

	assert.Equal(t, "Hello, World!\x00\x00\x00\x00", decodedData)
}

func TestHexDecoder(t *testing.T) {
	decoder := decoder.HexDecoder{}

	decodedData, _ := decoder.DecodeString("48656c6c6f2c20576f726c6421")

	assert.Equal(t, "Hello, World!", decodedData)
}

func TestBinaryDecoder(t *testing.T) {
	decoder := decoder.BinaryDecoder{}

	decodedData, _ := decoder.DecodeString("01001000011001010110110001101100011011110010110000100000010101110110111101110010011011000110010000100001")

	assert.Equal(t, "Hello, World!", decodedData)
}
