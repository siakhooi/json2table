/*
Package application run the application
*/
package application

import (
	"fmt"
	"os"
)

// ValidateArgs of cli
func ValidateArgs(args []string) (string, error) {
	if len(args) == 0 {
		// check if stdin is piped
		fi, err := os.Stdin.Stat()
		if err != nil {
			return "", fmt.Errorf("cannot stat stdin: %w", err)
		}
		// If stdin is not a character device, treat as piped input
		if (fi.Mode() & os.ModeCharDevice) == 0 {
			return "-", nil
		}

		return "", fmt.Errorf("inputJsonFile is required")
	}

	if len(args) > 1 {
		return "", fmt.Errorf("too many arguments")
	}
	filename := args[0]

	return filename, nil
}
