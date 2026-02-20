/*
Package application handles spec file operations
*/
package application

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// ReadData reads a data file
func ReadData(dataFilePath string) ([]byte, error) {
	if dataFilePath == "-" {
		// read from stdin
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("error reading stdin: %w", err)
		}
		return data, nil
	}
	// Check if file is readable
	_, err := os.Open(dataFilePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	// Read file contents
	data, err := os.ReadFile(dataFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return data, nil

}

// ParseData parses JSON data into an interface{}
func ParseData(data []byte) (interface{}, error) {
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}
	return jsonData, nil
}

// ReadParseData reads a data file and parses it into an interface{}
func ReadParseData(dataFilePath string) (interface{}, error) {
	data, err := ReadData(dataFilePath)
	if err != nil {
		return nil, err
	}
	return ParseData(data)
}
