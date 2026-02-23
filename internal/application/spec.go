/*
Package application run the application
*/
package application

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Column represents a column specification
type Column struct {
	Path  string `json:"path" validate:"required"`
	Title string `json:"title"`

	Width int
}

// Spec represents the specification structure
type Spec struct {
	DataPath string   `json:"dataPath"`
	Columns  []Column `json:"columns" validate:"required,min=1,dive"`
}

func (c *Column) setDefaults() {
	if c.Title == "" {
		parts := strings.Split(c.Path, ".")
		if len(parts) > 1 {
			c.Title = parts[len(parts)-1]
		} else {
			c.Title = c.Path
		}
	}
	c.Width = len(c.Title)
}
func (s *Spec) setDefaults() {
	if s.DataPath == "" {
		s.DataPath = "$"
	}
	for i := range s.Columns {
		s.Columns[i].setDefaults()
	}
}

// ParseAndValidateSpec parses JSON data into a Spec and validates it
func ParseAndValidateSpec(data []byte) (*Spec, error) {
	var spec Spec

	// Unmarshal JSON into spec struct
	err := json.Unmarshal(data, &spec)
	if err != nil {
		return nil, fmt.Errorf("error parsing spec: %w", err)
	}
	spec.setDefaults()

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
func ReadSpec(specFile string) ([]byte, error) {
	// Read the spec file
	data, err := os.ReadFile(specFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read spec file: %w", err)
	}
	return data, nil
}

// ReadParseValidateSpec reads a spec file, parses it, and validates it
func ReadParseValidateSpec(specFile string, envSpec string) (*Spec, error) {
	data := []byte(envSpec)
	if envSpec == "" {
		var err error
		data, err = ReadSpec(specFile)
		if err != nil {
			return nil, err
		}
	}
	spec, err := ParseAndValidateSpec(data)
	if err != nil {
		return nil, err
	}

	return spec, nil
}
