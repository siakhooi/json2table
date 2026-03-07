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
		prefix, printValue, suffix := FormatAlignedTextParts(fmt.Sprintf("%v", title), column.Width, column.Align)

		fmt.Printf("%s%s%s ", prefix, printValue, suffix)

	}
	fmt.Println("")
}

func analyseData(spec *Spec, dataArray []interface{}) {
	for _, item := range dataArray {
		for i, column := range spec.Columns {
			for _, path := range column.Path {
				value, err := jsonpath.Get(path, item)
				if err != nil {
					continue
				}
				if value == nil {
					continue
				}
				valueStr := fmt.Sprintf("%v", value)
				if len(valueStr) > column.Width {
					spec.Columns[i].Width = len(valueStr)
				}
				break
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

func printData(dataArray []interface{}, spec *Spec) {
	for _, item := range dataArray {
		for _, column := range spec.Columns {
			var value interface{}
			for _, path := range column.Path {
				v, err := jsonpath.Get(path, item)
				if err == nil && v != nil {
					value = v
					break
				}
			}
			valueStr := fmt.Sprintf("%v", value)
			prefix, printValue, suffix := FormatAlignedTextParts(valueStr, column.Width, column.Align)
			urlPath := column.URLPath
			if urlPath != "" {
				urlValue, err := jsonpath.Get(urlPath, item)
				if err == nil && urlValue != nil {
					urlStr := fmt.Sprintf("%v", urlValue)
					printValue = GetLink(printValue, urlStr)
				}
			}
			coloredPrintValue := GetColored(printValue, column.Color)
			fmt.Printf("%s%s%s ", prefix, coloredPrintValue, suffix)
		}
		fmt.Println("")
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
	printData(dataArray, spec)

	return nil
}
