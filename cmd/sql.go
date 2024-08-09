package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	sqlFilePath  string
	templateMode string
)

const (
	seleonMode   string = ":"
	percentMode  string = "%"
	mPercentMode string = "%%"
)

// sqlCmd represents the sql command
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "connect sql template with params",
	Long:  `connect sql template with params that would be a map[string]inteface{}`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if templateMode == percentMode && len(args) == 0 {
			cobra.CheckErr("'%' mode require position args")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sql called")
	},
}

func init() {
	rootCmd.AddCommand(sqlCmd)
	sqlCmd.Flags().StringVarP(&sqlFilePath, "sqlpath", "p", "", "sql template file path")
	sqlCmd.Flags().StringVarP(&sqlFilePath, "mode", "m", seleonMode, `sql template mode,
':' => ':param1',
'%' => '%s'
'%%'=> '%(params1)'
`)
}
