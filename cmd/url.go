package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

var decode bool
var encode bool

// urlencodeCmd
var urlencodeCmd = &cobra.Command{
	Use:   "urlencode [string]",
	Short: "Encode a string to URL-safe format",
	Long:  `Encode a string to URL-safe format.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		target := url.QueryEscape(source)
		fmt.Println(target)
	},
}

// urldecodeCmd
var urldecodeCmd = &cobra.Command{
	Use:   "urldecode [string]",
	Short: "Decode a URL-encoded string",
	Long:  `Decode a URL-encoded string.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		target, err := url.QueryUnescape(source)
		cobra.CheckErr(err)
		fmt.Println(target)
	},
}

// urlparseCmd
var urlparseCmd = &cobra.Command{
	Use:   "urlparse [url]",
	Short: "Parse a URL and print its query parameters",
	Long:  `Parse a URL and print its query parameters, one per line.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rawurl := args[0]
		u, err := url.Parse(rawurl)
		cobra.CheckErr(err)
		fmt.Printf("Scheme:   %s\n", u.Scheme)
		fmt.Printf("Host:     %s\n", u.Host)
		fmt.Printf("Path:     %s\n", u.Path)
		fmt.Printf("Fragment: %s\n", u.Fragment)
		fmt.Printf("Query:\n")
		values := u.Query()
		if len(values) == 0 {
			fmt.Println("  (none)")
		} else {
			for key, vals := range values {
				for _, v := range vals {
					fmt.Printf("  %s = %s\n", key, v)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(urlencodeCmd)
	rootCmd.AddCommand(urldecodeCmd)
	rootCmd.AddCommand(urlparseCmd)
}
