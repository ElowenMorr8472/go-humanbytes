package humanbytes

import "testing"

func TestFormat(t *testing.T) {
	cases := map[float64]string{
		512:            "512B",
		1024:           "1.0KB",
		1024 * 1024 * 3: "3.0MB",
	}
	for in, want := range cases {
		if got := Format(in); got != want {
			t.Errorf("Format(%v) = %q, want %q", in, got, want)
		}
	}
}

func TestParseRoundTrip(t *testing.T) {
	for _, n := range []float64{512, 1024, 1024 * 1024 * 3} {
		got, err := Parse(Format(n))
		if err != nil {
			t.Fatalf("Parse(Format(%v)) errored: %v", n, err)
		}
		if got != n {
			t.Errorf("round trip of %v gave %v", n, got)
		}
	}
	if _, err := Parse("abc"); err == nil {
		t.Error("Parse should reject input without a number")
	}
}
