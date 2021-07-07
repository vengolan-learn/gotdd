package numeral

import "testing"

func TestRomanNumeral(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      int
		Roman       string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to II", 3, "III"},
		{"4 gets converted to II", 4, "IV"},
		{"5 gets converted to II", 5, "V"},
		{"6 gets converted to II", 6, "VI"},
		{"7 gets converted to II", 7, "VII"},
		{"8 gets converted to II", 8, "VIII"},
		{"9 gets converted to II", 9, "IX"},
		{"14 gets converted to II", 14, "XIV"},
		{"18 gets converted to II", 18, "XVIII"},
		{"19 gets converted to II", 19, "XIX"},
		{"20 gets converted to II", 20, "XX"},
		{"39 gets converted to II", 39, "XXXIX"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			want := test.Roman
			got := ConvertToRoman(test.Arabic)
			if got != want {
				t.Errorf("want %q, but got %q", want, got)
			}
		})
	}

}
