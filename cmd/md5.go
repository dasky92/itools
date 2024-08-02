package cmd

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	isFile bool
)

// md5Cmd represents the md5 command
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "md5sum",
	Long:  `md5 sum for string.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !isFile {
			res := md5.Sum([]byte(args[0]))
			fmt.Printf("md5sum: %x\n", res)
			os.Exit(0)
		}
		f, err := os.Open(args[0])
		cobra.CheckErr(err)
		defer f.Close()

		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			cobra.CheckErr(err)
		}

		fmt.Printf("md5sum: %x\n", h.Sum(nil))
	},
}

func init() {
	rootCmd.AddCommand(md5Cmd)
	md5Cmd.Flags().BoolVarP(&isFile, "isfile", "f", false, "if the inptut is filepath ?")
}
