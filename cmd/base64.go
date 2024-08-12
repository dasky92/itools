package cmd

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "base64 encode & decode",
	Long:  `It is just basic base64 encode and decode.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !encode && !decode {
			cobra.CheckErr("Please use -e/-d")
		}
		var target string
		if encode {
			originStr := args[0]
			target = base64.StdEncoding.EncodeToString([]byte(originStr))
			fmt.Printf("base64 string: %s\n", target)
			os.Exit(0)
		}
		processStr := args[0]
		originStr, err := base64.StdEncoding.DecodeString(processStr)
		cobra.CheckErr(err)

		fmt.Printf("origin string: %s\n", string(originStr))
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)

	base64Cmd.Flags().BoolVarP(&decode, "decode", "d", false, "string decode to origin")
	base64Cmd.Flags().BoolVarP(&encode, "encode", "e", false, "string encode to base64 string")
}
