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

		return "", fmt.Errorf("dataFile is required")
	}

	if len(args) > 1 {
		return "", fmt.Errorf("too many arguments")
	}
	filename := args[0]

	return filename, nil
}

// ValidateSpecFile validates the spec file option
// Requires either the specFile parameter or JSON2TABLE_SPEC_FILE environment variable to be set
func ValidateSpecFile(specFile string) (string, error) {
	// If specFile is provided via CLI flag, use it
	if specFile != "" {
		return specFile, nil
	}

	// Try to read from environment variable
	envSpecFile := os.Getenv("JSON2TABLE_SPEC_FILE")
	if envSpecFile != "" {
		return envSpecFile, nil
	}

	// Neither CLI flag nor environment variable provided - spec is mandatory
	return "", fmt.Errorf("spec is mandatory: provide -s/--spec flag or set JSON2TABLE_SPEC_FILE environment variable")
}
