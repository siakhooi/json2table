/*
Package application handles spec file operations
*/
package application

import (
	"encoding/json"
	"fmt"
	"os"
)

// Column represents a column specification
type Column struct {
	Path string `json:"path"`
}

// Spec represents the specification structure
type Spec struct {
	Columns []Column `json:"columns"`
}

// ParseAndValidateSpec parses JSON data into a Spec and validates it
func ParseAndValidateSpec(data []byte) (*Spec, error) {
	var spec Spec

	// Unmarshal JSON into spec struct
	err := json.Unmarshal(data, &spec)
	if err != nil {
		return nil, fmt.Errorf("error parsing spec: %w", err)
	}

	// Validate minimum structure
	if err := ValidateSpec(&spec); err != nil {
		return nil, err
	}

	return &spec, nil
}

// ValidateSpec validates that the spec meets minimum requirements
func ValidateSpec(spec *Spec) error {
	if spec == nil {
		return fmt.Errorf("spec is nil")
	}

	if len(spec.Columns) == 0 {
		return fmt.Errorf("spec must contain at least one column")
	}

	for i, col := range spec.Columns {
		if col.Path == "" {
			return fmt.Errorf("column %d: path is required", i)
		}
	}

	return nil
}

// ReadSpec reads a spec file, validates it, and returns the parsed Spec
func ReadSpec(specFile string) (*Spec, error) {
	// Validate and get the spec file path
	validatedSpecFile, err := ValidateSpecFile(specFile)
	if err != nil {
		return nil, err
	}

	// Read the spec file
	data, err := os.ReadFile(validatedSpecFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read spec file: %w", err)
	}

	// Parse and validate the spec
	return ParseAndValidateSpec(data)
}
