package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lengthCmd represents the length command
var lenCmd = &cobra.Command{
	Use:   "len",
	Short: "calculate string length",
	Long: `Calculate string length.

For example:
input : hello world
output: 11`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Length: %d\n", len(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(lenCmd)
}
