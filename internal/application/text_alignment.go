/*
Package application run the application
*/
package application

import "fmt"

// Alignment represents text alignment in a column
type Alignment string

const (
	// AlignLeft aligns text to the left
	AlignLeft Alignment = "left"
	// AlignRight aligns text to the right
	AlignRight Alignment = "right"
	// AlignCenter aligns text to the center
	AlignCenter Alignment = "center"
)

// DefaultAlignment is the default text alignment (left)
var DefaultAlignment = AlignLeft

// FormatAlignedTextParts formats the given value according to the specified width and alignment, returning the prefix, formatted value, and suffix for proper alignment in a table cell.
func FormatAlignedTextParts(value string, width int, align Alignment) (string, string, string) {
	shortvalue := value
	if len(value) > width {
		shortvalue = value[:width]
	}
	prefix := ""
	suffix := ""
	if len(shortvalue) < width {
		padding := width - len(shortvalue)
		switch align {
		case AlignRight:
			prefix = fmt.Sprintf("%*s", padding, "")
		case AlignCenter:
			leftPad := padding / 2
			rightPad := padding - leftPad
			prefix = fmt.Sprintf("%*s", leftPad, "")
			suffix = fmt.Sprintf("%*s", rightPad, "")
		default: // AlignLeft
			suffix = fmt.Sprintf("%*s", padding, "")
		}
	}

	return prefix, shortvalue, suffix
}

func isValidAlignment(align Alignment) bool {
	validAlignments := map[Alignment]bool{
		AlignLeft:   true,
		AlignRight:  true,
		AlignCenter: true,
	}
	return validAlignments[align]
}
