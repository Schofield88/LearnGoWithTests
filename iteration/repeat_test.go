package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeats things 3 times", func(t *testing.T) {
		repeated := Repeat("a", 3)
	    expected := "aaa"

	    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
	}
	} )


	t.Run("Repeats things 6 times", func(t *testing.T) {
		repeated := Repeat("b", 6)
	    expected := "bbbbbb"

	if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
	}
	})
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
