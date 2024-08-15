package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// sqlcompileCmd represents the sqlcompile command
var sqlcompileCmd = &cobra.Command{
	Use:   "sqlcompile",
	Short: "Pre compile sql",
	Long: `Replace placeholder in sql with real parameters.
First positional flags is sql or it's path, Second positional flags is params or it's path.

For example:
select id,name from t where id=:id, {"id": 1} =>
select id,name from t where id=1`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var sql string
		var data []byte
		var params map[string]interface{}
		var err error

		// Parse sql
		data, err = os.ReadFile(args[0])
		if err != nil {
			sql = args[0]
		} else {
			sql = string(data)
		}

		// Parse parameters
		var d json.Decoder
		data, err = os.ReadFile(args[1])
		if err != nil {
			d = *json.NewDecoder(strings.NewReader(args[1]))
		} else {
			d = *json.NewDecoder(bytes.NewReader(data))
		}
		d.UseNumber()
		err = d.Decode(&params)
		cobra.CheckErr(err)

		// Replace params
		for k, v := range params {

			sql = strings.Replace(sql, fmt.Sprintf(":%s", k), convertValue(v), -1)
		}
		fmt.Println(sql)
	},
}

func init() {
	rootCmd.AddCommand(sqlcompileCmd)
}

func convertValue(value interface{}) string {
	var processValue string

	v := reflect.ValueOf(value)

	switch v.Kind() {
	// Compatible with int64, float64
	// If not, 1234567 would be converted to 1.234567+e6
	case reflect.String:
		x, ok := value.(json.Number)
		if !ok {
			processValue = fmt.Sprintf("'%s'", v)
			return processValue
		}
		i, err := x.Int64()
		if err != nil {
			f, err := x.Float64()
			if err != nil {
				processValue = fmt.Sprintf("'%f'", f)
			} else {
				processValue = fmt.Sprintf("%f", f)
			}
		} else {
			processValue = fmt.Sprintf("%v", i)
		}
	case reflect.Array, reflect.Slice:
		var list []string
		realValue := value.([]interface{})
		for i := 0; i < len(realValue); i++ {
			list = append(list, convertValue(realValue[i]))
		}
		processValue = "(" + strings.Join(list, ",") + ")"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		x, _ := value.(int)
		processValue = strconv.Itoa(x)
	case reflect.Int64:
		processValue = fmt.Sprintf("%v", value)
	case reflect.Float32, reflect.Float64:
		processValue = fmt.Sprintf("%v", v.Float())
	case reflect.Bool:
		processValue = fmt.Sprintf("%v", v.Bool())
	default:
		processValue = fmt.Sprintf("%s", value)
	}
	return processValue
}
