package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var outputFormats = []struct {
	label  string
	format string
}{
	{"RFC3339", time.RFC3339},
	{"ISO8601", "2006-01-02T15:04:05Z07:00"},
	{"ANSIC", time.ANSIC},
	{"UnixDate", time.UnixDate},
	{"RFC822", time.RFC822},
	{"RFC822Z", time.RFC822Z},
	{"RFC850", time.RFC850},
	{"RFC1123", time.RFC1123},
	{"RFC1123Z", time.RFC1123Z},
	{"YYYY-MM-DD", "2006-01-02"},
	{"YYYY/MM/DD", "2006/01/02"},
	{"YYYY-MM-DD HH:MM:SS", "2006-01-02 15:04:05"},
	{"YYYY/MM/DD HH:MM:SS", "2006/01/02 15:04:05"},
	{"s", ""},
	{"ms", ""},
	{"us", ""},
	{"ns", ""},
}

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "time converter.",
	Long:  `Convert bewteen timestamp and datetime string.`,
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			ts        int64
			err       error
			source    string
			target    time.Time
			onlyLabel string
		)
		onlyFormat, _ := cmd.Flags().GetString("format")
		if onlyFormat != "" {
			onlyLabel = onlyFormat
		}
		if len(args) == 0 {
			now := time.Now()
			printTimeFormat(now, onlyLabel)
			os.Exit(0)
		} else {
			source = args[0]
			ts, err = strconv.ParseInt(source, 10, 64)
		}

		// convert timestamp to datetime format
		if err == nil {
			switch len(source) {
			case 10:
				target = time.Unix(ts, 0)
			case 13:
				target = time.UnixMilli(ts)
			case 16:
				target = time.UnixMicro(ts)
			case 19:
				target = time.Unix(0, ts)
			default:
				cobra.CheckErr("timestamp format: " + source)
			}
			fmt.Println("Parsed time from timestamp:")
			printTimeFormat(target, onlyLabel)
			os.Exit(0)
		}

		var parseErr error
		for _, outputFormat := range outputFormats {
			target, parseErr = time.Parse(outputFormat.label, source)
			if parseErr == nil {
				break
			}
		}
		if parseErr != nil {
			cobra.CheckErr("unrecoginzed tiem format: " + source)
		}

		fmt.Println("Parsed time from string:")
		printTimeFormat(target, onlyLabel)
		os.Exit(0)
	},
}

func init() {
	var formatOptions []string
	for _, f := range outputFormats {
		if f.label != "" {
			formatOptions = append(formatOptions, f.label)
		}
		if f.format != "" && f.format != f.label {
			formatOptions = append(formatOptions, f.format)
		}
	}
	timeCmd.Flags().StringP("format", "f", "", fmt.Sprintf(
		"only output the specified format label or layout, options: [%s]",
		formatOptionsString(formatOptions),
	))
	rootCmd.AddCommand(timeCmd)
}

// formatOptionsString 用于帮助信息展示
func formatOptionsString(opts []string) string {
	seen := make(map[string]struct{})
	var uniq []string
	for _, o := range opts {
		if _, ok := seen[o]; !ok {
			uniq = append(uniq, o)
			seen[o] = struct{}{}
		}
	}
	return strings.Join(uniq, " | ")
}

// printTimeFormat 输出时间
func printTimeFormat(t time.Time, onlyFormat string) {
	for _, f := range outputFormats {
		if onlyFormat != "" && f.label != onlyFormat && f.format != onlyFormat {
			continue
		}
		switch f.label {
		case "s":
			fmt.Printf("  %-19s: %d\n", f.label, t.Unix())
		case "ms":
			fmt.Printf("  %-19s: %d\n", f.label, t.UnixMilli())
		case "us":
			fmt.Printf("  %-19s: %d\n", f.label, t.UnixMicro())
		case "ns":
			fmt.Printf("  %-19s: %d\n", f.label, t.UnixNano())
		default:
			fmt.Printf("  %-19s: %s\n", f.label, t.Format(f.format))
		}
	}
}
