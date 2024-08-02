package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// unicodeCmd represents the unicode command
var unicodeCmd = &cobra.Command{
	Use:   "unicode",
	Short: "unicode process tool",
	Long:  `switch unicode to string each other.`,
	Run: func(cmd *cobra.Command, args []string) {
		if decode && encode {
			cobra.CheckErr("Only one of decode, encode is allowed to specificy")
		}
		source := args[0]
		if encode {
			target := strconv.QuoteToASCII(source)
			fmt.Println(target)
			os.Exit(0)
		}
		target, err := strconv.Unquote("\"" + source + "\"")
		cobra.CheckErr(err)
		fmt.Println(target)
	},
}

func init() {
	rootCmd.AddCommand(unicodeCmd)

	unicodeCmd.Flags().BoolVarP(&decode, "decode", "d", false, "string decode to safe string")
	unicodeCmd.Flags().BoolVarP(&encode, "encode", "e", false, "encode to safe string")
}
