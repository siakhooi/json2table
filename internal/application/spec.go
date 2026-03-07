/*
Package application run the application
*/
package application

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Column represents a column specification
type Column struct {
	Path     StringList    `json:"path" validate:"required"`
	Title    string        `json:"title"`
	MinWidth int           `json:"minWidth" validate:"min=0,ltefield=MaxWidth"`
	MaxWidth int           `json:"maxWidth" validate:"min=0,gtefield=MinWidth"`
	Align    Alignment     `json:"align" validate:"omitempty,oneof=left right center"`
	URLPath  string        `json:"urlPath"`
	Color    TextColorSpec `json:"color"`

	Width int
}

// Spec represents the specification structure
type Spec struct {
	DataPath string   `json:"dataPath"`
	Columns  []Column `json:"columns" validate:"required,min=1,dive"`
}

func (c *Column) setDefaults() {
	if c.Title == "" {
		parts := strings.Split(c.Path[0], ".")
		if len(parts) > 1 {
			c.Title = parts[len(parts)-1]
		} else {
			c.Title = c.Path[0]
		}
	}
	c.Width = len(c.Title)
	if c.MaxWidth == 0 {
		c.MaxWidth = math.MaxInt
	}
	if c.Align == "" {
		c.Align = DefaultAlignment
	}
	if len(c.Color.Default) == 0 {
		c.Color = DefaultTextColor
	}
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

	return validate.Struct(spec)
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

func validateSpecFileValues(spec *Spec) error {
	for i, col := range spec.Columns {
		if col.Align != "" && !isValidAlignment(col.Align) {
			return fmt.Errorf("column %d: invalid align value: %q", i, col.Align)
		}

		for _, c := range col.Color.Default {
			if !isValidTextColor(c) {
				return fmt.Errorf("column %d: invalid color value: %q", i, c)
			}
		}
	}
	return nil
}

// ReadParseValidateSpec reads a spec file, parses it, and validates it
func ReadParseValidateSpec(specFile, envSpec string) (*Spec, error) {
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

	if err := validateSpecFileValues(spec); err != nil {
		return nil, err
	}

	return spec, nil
}
