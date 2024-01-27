package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/grqphical/go-encode/internal/encoder"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all valid encoding formats",
	Long:  `Lists all valid encoding formats`,
	Run:   listFormats,
}

func listFormats(cmd *cobra.Command, args []string) {
	encoders := encoder.GetEncoders()

	for name := range encoders {
		fmt.Println(name)
	}
}
