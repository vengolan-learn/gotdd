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
		{"3 gets converted to II", 2, "III"},
		{"4 gets converted to II", 2, "IV"},
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
