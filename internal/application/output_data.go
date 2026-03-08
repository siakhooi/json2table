/*
Package application run the application
*/
package application

import (
	"fmt"

	"github.com/PaesslerAG/jsonpath"
)

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
			fullValue := fmt.Sprintf("%v", value)
			prefix, printValue, suffix := FormatAlignedTextParts(fullValue, column.Width, column.Align)
			urlPath := column.URLPath
			if urlPath != "" {
				urlValue, err := jsonpath.Get(urlPath, item)
				if err == nil && urlValue != nil {
					urlStr := fmt.Sprintf("%v", urlValue)
					printValue = GetLink(printValue, urlStr)
				}
			}
			coloredPrintValue := GetColored(fullValue, printValue, column.Color)
			fmt.Printf("%s%s%s ", prefix, coloredPrintValue, suffix)
		}
		fmt.Println("")
	}

}
