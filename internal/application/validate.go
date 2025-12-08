/*
Package application run the application
*/
package application

import "fmt"

// ValidateArgs of cli
func ValidateArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("inputJsonFile is required")
	}

	if len(args) > 1 {
		return "", fmt.Errorf("too many arguments")
	}
	filename := args[0]

	return filename, nil
}
