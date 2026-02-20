/*
Package application run the application
*/
package application

import (
	"fmt"

	"github.com/PaesslerAG/jsonpath"
)

// PrintTable prints the spec and JSON data in a pretty format
func PrintTable(spec *Spec, fullData interface{}) error {

	dataPath := spec.DataPath
	data, err := jsonpath.Get(dataPath, fullData)
	if err != nil {
		return fmt.Errorf("error selecting data with jsonpath: %w", err)
	}

	//iterate data as array and print each item
	if dataArray, ok := data.([]interface{}); ok {
		for _, item := range dataArray {
			for _, column := range spec.Columns {
				value, err := jsonpath.Get(column.Path, item)
				if err != nil {
					return fmt.Errorf("error selecting column data with jsonpath: %w", err)
				}
				fmt.Printf("%v ", value)
			}
			fmt.Println("")
		}
		return nil
	}

	return fmt.Errorf("data selected with jsonpath is not an array")
}
