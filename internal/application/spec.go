/*
Package application handles spec file operations
*/
package application

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)

// Column represents a column specification
type Column struct {
	Path string `json:"path" validate:"required"`
}

// Spec represents the specification structure
type Spec struct {
	Columns []Column `json:"columns" validate:"required,min=1,dive"`
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

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(spec)

	return err
}

// ReadSpec reads a spec file
func ReadSpec(validatedSpecFile string) ([]byte, error) {
	// Read the spec file
	data, err := os.ReadFile(validatedSpecFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read spec file: %w", err)
	}
	return data, nil
}
