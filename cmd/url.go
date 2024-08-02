package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var decode bool
var encode bool

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "url process tool",
	Long:  `decode & encode.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if decode && encode {
			cobra.CheckErr("Only one of decode, encode is allowed to specificy")
		}
		source := args[0]
		if decode {
			target, err := url.QueryUnescape(source)
			cobra.CheckErr(err)
			fmt.Println(target)
			os.Exit(0)
		}
		target := url.QueryEscape(source)
		fmt.Println(target)
	},
}

func init() {
	rootCmd.AddCommand(urlCmd)

	urlCmd.Flags().BoolVarP(&decode, "decode", "d", false, "string decode to safe string")
	urlCmd.Flags().BoolVarP(&encode, "encode", "e", false, "encode to safe string")
}
