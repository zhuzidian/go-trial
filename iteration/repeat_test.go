package iteration

import "testing"

func TestRepeat(t *testing.T) {

	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("期待 %q 实际 %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
