/*
Package application handles spec file operations
*/
package application

import (
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
