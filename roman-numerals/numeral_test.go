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
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"19 gets converted to XIX", 19, "XIX"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
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
