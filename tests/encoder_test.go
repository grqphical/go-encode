package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/grqphical/go-encode/internal/encoder"
)

func TestBase64Encoder(t *testing.T) {
	text := "SGVsbG8sIFdvcmxkIQ=="

	encoder := encoder.Base64Encoder{}

	assert.Equal(t, text, encoder.EncodeString("Hello, World!"))
}

func TestBase32Encoder(t *testing.T) {
	text := "JBSWY3DPFQQFO33SNRSCC==="

	encoder := encoder.Base32Encoder{}

	assert.Equal(t, text, encoder.EncodeString("Hello, World!"))
}

func TestAscii85Encoder(t *testing.T) {
	text := "87cURD_*#4DfTZ)+TMKB"

	encoder := encoder.Ascii85Encoder{}

	assert.Equal(t, text, encoder.EncodeString("Hello, World!"))
}

func TestHexEncoder(t *testing.T) {
	text := "48656c6c6f2c20576f726c6421"

	encoder := encoder.HexEncoder{}

	assert.Equal(t, text, encoder.EncodeString("Hello, World!"))
}
