package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/grqphical/go-encode/internal/decoder"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes text from a certain format",
	Long: `Decodes text from a certain format. 

    For Example:
    go-encode decode -t "Hello, World!" -f base64`,
	Run: decodeText,
}

func decodeText(cmd *cobra.Command, args []string) {
	text, _ := cmd.Flags().GetString("text")
	format, _ := cmd.Flags().GetString("format")

	var textToEncode string
	var err error

	if text == "" {
		textToEncode, err = readFromStdin()
		if err != nil {
			if err == ErrNoInput {
				cmd.PrintErrln(err.Error())
				os.Exit(1)
			}
		}
	} else {
		textToEncode = text
	}

	decoders := decoder.GetDecoders()

	decoder, ok := decoders[format]
	if !ok {
		cmd.PrintErrln("invalid format provided")
		os.Exit(1)
	}

	result, err := decoder.DecodeString(textToEncode)
	if err != nil {
		cmd.PrintErrln("invalid data provided")
		os.Exit(1)
	}
	fmt.Println(result)
}
