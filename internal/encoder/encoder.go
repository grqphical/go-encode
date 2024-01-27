package encoder

type Encoder interface {
	EncodeString(string) string
}

func GetEncoders() map[string]Encoder {
	var encoders map[string]Encoder = map[string]Encoder{}

	encoders["base64"] = &Base64Encoder{}
	encoders["ascii85"] = &Ascii85Encoder{}
	encoders["hex"] = &HexEncoder{}

	return encoders
}
