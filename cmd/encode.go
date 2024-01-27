package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/grqphical/go-encode/internal/encoder"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode text to a certain format",
	Long: `Encodes text to a certain format. 

    For Example:
    go-encode encode -t "Hello, World!" -f base64`,
	Run: encodeText,
}

var ErrNoInput error = errors.New("no input data provided")

// reads data from stdin if it has been piped to the process
// returns an error if it wasn't
func readFromStdin() (string, error) {
	stat, _ := os.Stdin.Stat()
	// Make sure we are reading from a piped stdin and not the terminal so program doesnt hang
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ := io.ReadAll(os.Stdin)
		return string(bytes), nil
	} else {
		return "", ErrNoInput
	}
}

func encodeText(cmd *cobra.Command, args []string) {
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

	encoders := encoder.GetEncoders()

	encoder, ok := encoders[format]
	if !ok {
		cmd.PrintErrln("invalid format provided")
		os.Exit(1)
	}

	result := encoder.EncodeString(textToEncode)
	fmt.Println(result)
}
