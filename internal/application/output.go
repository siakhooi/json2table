/*
Package application run the application
*/
package application

import (
	"fmt"

	"github.com/PaesslerAG/jsonpath"
)

func selectDataArray(dataPath string, fullData interface{}) ([]interface{}, error) {
	data, err := jsonpath.Get(dataPath, fullData)
	if err != nil {
		return nil, fmt.Errorf("error selecting data with jsonpath: %w", err)
	}

	if dataArray, ok := data.([]interface{}); ok {
		return dataArray, nil
	}

	return nil, fmt.Errorf("data selected with jsonpath is not an array")
}

func printHeader(columns []Column) {
	for _, column := range columns {
		fmt.Printf("%*s ", -column.Width, column.Title)
	}
	fmt.Println("")
}

func analyseData(spec *Spec, dataArray []interface{}) {
	for _, item := range dataArray {
		for i, column := range spec.Columns {
			value, err := jsonpath.Get(column.Path, item)
			if err != nil {
				continue
			}
			valueStr := fmt.Sprintf("%v", value)
			if len(valueStr) > column.Width {
				spec.Columns[i].Width = len(valueStr)
			}
		}
	}
}

// PrintTable prints JSON data in tabular format based on the provided specification
func PrintTable(spec *Spec, fullData interface{}) error {
	dataArray, err := selectDataArray(spec.DataPath, fullData)
	if err != nil {
		return err
	}
	analyseData(spec, dataArray)

	printHeader(spec.Columns)

	for _, item := range dataArray {
		for _, column := range spec.Columns {
			value, err := jsonpath.Get(column.Path, item)
			if err != nil {
				value = nil
			}
			fmt.Printf("%*v ", -column.Width, value)
		}
		fmt.Println("")
	}
	return nil
}
