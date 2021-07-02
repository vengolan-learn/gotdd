package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://facebook.com"
	fastURL := "htp://www.quii.co.uk"

	want := fastURL
	got := Racer(slowURL, fastURL)
	if got != want {
		t.Errorf("want %q, but got %q", want, got)
	}

}
