package format

import "testing"

func TestFormatQR(t *testing.T) {
	// Test cases for FormatQR function
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"00020101021126400010vn.zalopay0115rN4omo4IIXVKkVF020300238620010A00000072701320006970454011899ZP23356M058497860208QRIBFTTA5204739953037045802VN63046DF0",
			`00 02 01
01 02 11
26 40
. . . 00 10 vn.zalopay
. . . 01 15 rN4omo4IIXVKkVF
. . . 02 03 002
38 62
. . . 00 10 A000000727
. . . 01 32
. . . . . . 00 06 970454
. . . . . . 01 18 99ZP23356M05849786
. . . 02 08 QRIBFTTA
52 04 7399
53 03 704
58 02 VN
63 04 6DF0
`, nil},
	}

	for _, test := range tests {
		result, err := FormatQR(test.input)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("expected error\n %v, got\n %v", test.err, err)
			continue
		}
		if result != test.expected {
			t.Errorf("expected\n %s, got\n %s", test.expected, result)
		}
	}
}
