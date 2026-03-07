package application

import "testing"

func TestFormatAlignedTextParts(t *testing.T) {
	tests := []struct {
		name       string
		value      string
		width      int
		align      Alignment
		wantPrefix string
		wantValue  string
		wantSuffix string
	}{
		{
			name:       "left align pads suffix",
			value:      "ab",
			width:      5,
			align:      AlignLeft,
			wantPrefix: "",
			wantValue:  "ab",
			wantSuffix: "   ",
		},
		{
			name:       "right align pads prefix",
			value:      "ab",
			width:      5,
			align:      AlignRight,
			wantPrefix: "   ",
			wantValue:  "ab",
			wantSuffix: "",
		},
		{
			name:       "center align splits padding",
			value:      "ab",
			width:      5,
			align:      AlignCenter,
			wantPrefix: " ",
			wantValue:  "ab",
			wantSuffix: "  ",
		},
		{
			name:       "center align odd padding with longer text",
			value:      "cat",
			width:      8,
			align:      AlignCenter,
			wantPrefix: "  ",
			wantValue:  "cat",
			wantSuffix: "   ",
		},
		{
			name:       "truncates when value exceeds width",
			value:      "abcdef",
			width:      4,
			align:      AlignLeft,
			wantPrefix: "",
			wantValue:  "abcd",
			wantSuffix: "",
		},
		{
			name:       "exact width has no padding",
			value:      "abcd",
			width:      4,
			align:      AlignRight,
			wantPrefix: "",
			wantValue:  "abcd",
			wantSuffix: "",
		},
		{
			name:       "zero width yields empty value",
			value:      "abcd",
			width:      0,
			align:      AlignLeft,
			wantPrefix: "",
			wantValue:  "",
			wantSuffix: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotPrefix, gotValue, gotSuffix := FormatAlignedTextParts(tc.value, tc.width, tc.align)

			if gotPrefix != tc.wantPrefix || gotValue != tc.wantValue || gotSuffix != tc.wantSuffix {
				t.Fatalf(
					"FormatAlignedTextParts(%q, %d, %q) = (%q, %q, %q), want (%q, %q, %q)",
					tc.value,
					tc.width,
					tc.align,
					gotPrefix,
					gotValue,
					gotSuffix,
					tc.wantPrefix,
					tc.wantValue,
					tc.wantSuffix,
				)
			}
		})
	}
}
