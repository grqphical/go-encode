package decoder

type Decoder interface {
	DecodeString(string) (string, error)
}

func GetDecoders() map[string]Decoder {
	var encoders map[string]Decoder = map[string]Decoder{}

	encoders["base64"] = &Base64Decoder{}
	encoders["ascii85"] = &Ascii85Decoder{}

	return encoders
}
