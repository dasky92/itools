package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	prefix string
	indent string
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json formatter",
	Long:  `json formatter`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			obj  interface{}
			data []byte
			err  error
		)
		data = []byte(args[0])
		err = json.Unmarshal(data, &obj)
		cobra.CheckErr(err)

		data, err = json.MarshalIndent(obj, prefix, indent)
		cobra.CheckErr(err)

		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	jsonCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "json prefix character")
	jsonCmd.Flags().StringVarP(&indent, "indent", "i", "  ", "json indent character")
}
