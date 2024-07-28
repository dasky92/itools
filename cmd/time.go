package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	sourceFormat, targetFormat string
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "time converter.",
	Long:  `Convert bewteen timestamp and datetime string.`,
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			ts     int64
			err    error
			source string
			target time.Time
		)

		source = args[0]
		// convert timestamp to datetime format
		ts, err = strconv.ParseInt(source, 10, 64)
		if err == nil {
			if len(source) != 10 && len(source) != 13 && len(source) != 16 {
				cobra.CheckErr("timestamp format: " + source)
			}
			switch len(source) {
			case 10:
				target = time.Unix(ts, 0)
			case 13:
				target = time.UnixMilli(ts)
			case 16:
				target = time.UnixMicro(ts)
			default:
				cobra.CheckErr("timestamp format: " + source)
			}
			fmt.Printf("dateformat: %s\n", target.Format(time.RFC3339))
			os.Exit(0)
		}
		// convert datetime format to timeistamp
		target, err = time.Parse(time.RFC3339, source)
		cobra.CheckErr(err)

		fmt.Printf("timestamp: %d\n", target.UnixMilli())
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)

	// Now rfc3339 only is supported.
	timeCmd.Flags().StringVarP(&sourceFormat, "sourceFormat", "s", "rfc3339", "")
	timeCmd.Flags().StringVarP(&targetFormat, "targetFormat", "t", "rfc3339", "")
}

// AppendNano append given int64 to 16 bits.
func AppendNano(sec int64) (int64, error) {
	secStr := strconv.FormatInt(sec, 10)
	if len(secStr) < 0 || len(secStr) > 16 {
		return sec, errors.New("timestamp format error")
	} else if len(secStr) == 16 {
		return sec, nil
	} else {
		appendBit := 16 - len(secStr)
		return sec << appendBit, nil
	}
}