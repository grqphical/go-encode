package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-encode",
	Short: "A tool to encode/decode text in different formats",
	Long: `A tool to encode/decode text in different formats such as base64, base32, hex, binary and many more. Use the list to command to view all
    available formats.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	encodeCmd.Flags().StringP("text", "t", "", "Text to encode. Leave blank to read from STDIN")
	encodeCmd.Flags().StringP("format", "f", "", "Encoding format to use")
	encodeCmd.MarkFlagRequired("format")

	rootCmd.AddCommand(decodeCmd)

	decodeCmd.Flags().StringP("text", "t", "", "Text to encode. Leave blank to read from STDIN")
	decodeCmd.Flags().StringP("format", "f", "", "Encoding format to use")
	decodeCmd.MarkFlagRequired("format")

	rootCmd.AddCommand(listCmd)
}
