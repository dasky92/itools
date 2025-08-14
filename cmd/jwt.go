package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cobra"
)

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt [token]",
	Short: "Parse and display JWT token information",
	Long:  `Parse a JWT token and display its header, payload, and signature information.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tokenString := args[0]
		parts := strings.Split(tokenString, ".")

		if len(parts) != 3 {
			fmt.Println("Error: Invalid JWT token format. Expected 3 parts separated by dots.")
			return
		}

		// Parse without validating the signature
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// We're only parsing, not validating, so we return nil
			return nil, nil
		})

		// Print header
		fmt.Println("Header:")
		headerJSON, err := json.MarshalIndent(token.Header, "", "  ")
		if err != nil {
			fmt.Printf("  Error parsing header: %v\n", err)
		} else {
			fmt.Println(string(headerJSON))
		}

		// Print claims (payload)
		fmt.Println("\nPayload:")
		claimsJSON, err := json.MarshalIndent(token.Claims, "", "  ")
		if err != nil {
			fmt.Printf("  Error parsing payload: %v\n", err)
		} else {
			fmt.Println(string(claimsJSON))
		}

		// Print signature info
		fmt.Println("\nSignature:")
		fmt.Printf("  %s\n", parts[2])
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)
}