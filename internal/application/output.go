/*
Package application handles spec file operations
*/
package application

import (
	"encoding/json"
	"fmt"
)

// PrintTable prints the spec and JSON data in a pretty format
func PrintTable(spec *Spec, data interface{}) error {
	// Pretty print spec JSON
	prettyJSON, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		return fmt.Errorf("error formatting spec file: %w", err)
	}

	fmt.Println(string(prettyJSON))

	// Pretty print JSON
	prettyJSON1, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error formatting JSON: %w", err)
	}

	fmt.Println(string(prettyJSON1))

	return nil
}
