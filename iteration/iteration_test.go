package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := strings.Repeat("a", 5000)
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("alonger stringstill longer to check the efficiency of buf")

	}
}
func BenchmarkRepeatBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatBuf("alonger stringstill longer to check the efficiency of buf")
	}
}
