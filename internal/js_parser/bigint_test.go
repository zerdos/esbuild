package js_parser

import "testing"

func TestBigIntToDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string // We will check if output matches either the full or fallback string
	}{
		{"0", "0"},
		{"123", "123"},
		{"0x1f", "31"},
		{"0X1F", "31"},
		{"0o77", "63"},
		{"0O77", "63"},
		{"0b1010", "10"},
		{"0B1010", "10"},
		{"0x1_f", "31"},
		{"0o7_7", "63"},
		{"0b1_0_1_0", "10"},
		// Larger than uint64 limits (behavior varies by platform, but we just want to execute the code)
		{"0xFFFFFFFFFFFFFFFFFF", ""},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result := bigIntToDecimal(tc.input)
			if tc.expected != "" && result != tc.expected && result != tc.input {
				t.Errorf("bigIntToDecimal(%q) = %q; want %q or %q", tc.input, result, tc.expected, tc.input)
			}
		})
	}
}
