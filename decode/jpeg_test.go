package decode

import (
	"image/color"
	"testing"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		in   string
		want color.Color
	}{
		{"../testimages/black-50x50.jpg", color.RGBA{0, 0, 0, 255}},
		{"../testimages/white-50x50.jpg", color.RGBA{255, 255, 255, 255}},
	}
	for _, c := range cases {
		got, _ := AvgColorFromFile(c.in)
		if got != c.want {
			t.Errorf("AvgColorFromFile(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AvgColorFromFile("../testimages/black-50x50.jpg")
		AvgColorFromFile("../testimages/white-50x50.jpg")
		AvgColorFromFile("../testimages/barracks.jpg")
	}
}
