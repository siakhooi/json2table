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
		title := column.Title
		if len(title) > column.Width {
			title = title[:column.Width]
		}
		fmt.Printf("%*s ", -column.Width, title)
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

func optimizeSpec(spec *Spec) {
	for i, column := range spec.Columns {
		if column.MinWidth > 0 && column.Width < column.MinWidth {
			spec.Columns[i].Width = column.MinWidth
		}
		if column.MaxWidth > 0 && column.Width > column.MaxWidth {
			spec.Columns[i].Width = column.MaxWidth
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
	optimizeSpec(spec)

	printHeader(spec.Columns)

	for _, item := range dataArray {
		for _, column := range spec.Columns {
			value, err := jsonpath.Get(column.Path, item)
			if err != nil {
				value = nil
			}
			valueStr := fmt.Sprintf("%v", value)
			if len(valueStr) > column.Width {
				valueStr = valueStr[:column.Width]
			}
			fmt.Printf("%*s ", -column.Width, valueStr)
		}
		fmt.Println("")
	}
	return nil
}
