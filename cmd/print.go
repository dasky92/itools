/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print",
	Long:  `Print string to os.stdout with \n, it would be more beautify`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		segments := strings.Split(args[0], "\\n")
		for _, v := range segments {
			fmt.Println(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
