package colors

import "testing"

func TestRgbToHex(t *testing.T) {
	tests := []struct {
		input *Rgb
		want  *Hex
	}{
		{&Rgb{251, 252, 250}, &Hex{"#fbfcfa"}},
		{&Rgb{0, 0, 0}, &Hex{"#000000"}},
		{&Rgb{255, 255, 255}, &Hex{"#ffffff"}},
		{&Rgb{250, 22, 25}, &Hex{"#fa1619"}},
		{&Rgb{2, 52, 250}, &Hex{"#0234fa"}},
	}

	for _, test := range tests {
		if *test.input.ToHex() != *test.want {
			t.Errorf("%v should be %v, found %v", test.input, test.want, test.input.ToHex())
		}
	}
}

func TestRgbToHsv(t *testing.T) {
	tests := []struct {
		input *Rgb
		want  *Hsv
	}{
		{&Rgb{251, 252, 250}, &Hsv{90, 1, 99}},
		{&Rgb{0, 0, 0}, &Hsv{0, 0, 0}},
		{&Rgb{255, 255, 255}, &Hsv{0, 0, 100}},
		{&Rgb{250, 22, 25}, &Hsv{359, 91, 98}},
		{&Rgb{2, 52, 250}, &Hsv{228, 99, 98}},
	}

	for _, test := range tests {
		if *test.input.ToHsv() != *test.want {
			t.Errorf("%v should be %v, found %v", test.input, test.want, test.input.ToHsv())
		}
	}
}

func TestRgbToHsl(t *testing.T) {
	tests := []struct {
		input *Rgb
		want  *Hsl
	}{
		{&Rgb{251, 252, 250}, &Hsl{90, 25, 98}},
		{&Rgb{0, 0, 0}, &Hsl{0, 0, 0}},
		{&Rgb{255, 255, 255}, &Hsl{0, 0, 100}},
		{&Rgb{250, 22, 25}, &Hsl{359, 96, 53}},
		{&Rgb{2, 52, 250}, &Hsl{228, 98, 49}},
	}

	for _, test := range tests {
		if *test.input.ToHsl() != *test.want {
			t.Errorf("%v should be %v, found %v", test.input, test.want, test.input.ToHsl())
		}
	}
}

func TestRgbToCmyk(t *testing.T) {
	tests := []struct {
		input *Rgb
		want  *Cmyk
	}{
		{&Rgb{251, 252, 250}, &Cmyk{0, 0, 1, 1}},
		{&Rgb{0, 0, 0}, &Cmyk{0, 0, 0, 100}},
		{&Rgb{255, 255, 255}, &Cmyk{0, 0, 0, 0}},
		{&Rgb{250, 22, 25}, &Cmyk{0, 91, 90, 2}},
		{&Rgb{2, 52, 250}, &Cmyk{99, 79, 0, 2}},
	}

	for _, test := range tests {
		if *test.input.ToCmyk() != *test.want {
			t.Errorf("%v should be %v, found %v", test.input, test.want, test.input.ToCmyk())
		}
	}
}

func TestHexToRgb(t *testing.T) {
	tests := []struct {
		want  *Rgb
		input *Hex
	}{
		{&Rgb{251, 252, 250}, &Hex{"#fbfcfa"}},
		{&Rgb{0, 0, 0}, &Hex{"#000000"}},
		{&Rgb{255, 255, 255}, &Hex{"#ffffff"}},
		{&Rgb{250, 22, 25}, &Hex{"#fa1619"}},
		{&Rgb{2, 52, 250}, &Hex{"#0234fa"}},
	}

	for _, test := range tests {
		if *test.input.ToRgb() != *test.want {
			t.Errorf("%v should be %v, found %v", test.input, test.want, test.input.ToRgb())
		}
	}
}

func TestHsvToRgb(t *testing.T) {
	tests := []struct {
		want  *Rgb
		input *Hsv
	}{
		{&Rgb{251, 252, 250}, &Hsv{90, 1, 99}},
		{&Rgb{0, 0, 0}, &Hsv{0, 0, 0}},
		{&Rgb{255, 255, 255}, &Hsv{0, 0, 100}},
		{&Rgb{250, 22, 26}, &Hsv{359, 91, 98}},
		{&Rgb{2, 52, 250}, &Hsv{228, 99, 98}},
	}

	for _, test := range tests {
		if *test.input.ToRgb() != *test.want {
			t.Errorf("%v should be %v, found %v", test.input, test.want, test.input.ToRgb())
		}
	}
}

func TestHslToRgb(t *testing.T) {
	tests := []struct {
		want  *Rgb
		input *Hsl
	}{
		{&Rgb{250, 251, 249}, &Hsl{90, 25, 98}},
		{&Rgb{0, 0, 0}, &Hsl{0, 0, 0}},
		{&Rgb{255, 255, 255}, &Hsl{0, 0, 100}},
		{&Rgb{250, 20, 24}, &Hsl{359, 96, 53}},
		{&Rgb{2, 51, 247}, &Hsl{228, 98, 49}},
	}

	for _, test := range tests {
		if *test.input.ToRgb() != *test.want {
			t.Errorf("%v should be %v, found %v", test.input, test.want, test.input.ToRgb())
		}
	}
}

func TestCmykToRgb(t *testing.T) {
	tests := []struct {
		want  *Rgb
		input *Cmyk
	}{
		{&Rgb{252, 252, 250}, &Cmyk{0, 0, 1, 1}},
		{&Rgb{0, 0, 0}, &Cmyk{0, 0, 0, 100}},
		{&Rgb{255, 255, 255}, &Cmyk{0, 0, 0, 0}},
		{&Rgb{250, 22, 25}, &Cmyk{0, 91, 90, 2}},
		{&Rgb{2, 52, 250}, &Cmyk{99, 79, 0, 2}},
	}

	for _, test := range tests {
		if *test.input.ToRgb() != *test.want {
			t.Errorf("%v should be %v, found %v", test.input, test.want, test.input.ToRgb())
		}
	}
}
