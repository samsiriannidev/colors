package colors

import "math"

// Hsl represents the HSL model
type Hsl struct {
	Hue        int
	Saturation int
	Lightness  int
}

// ToRgb performs a conversion from HSL to RGB
func (hsl *Hsl) ToRgb() *Rgb {
	var r, g, b float64

	s := float64(hsl.Saturation) / 100.0
	l := float64(hsl.Lightness) / 100.0
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(float64(hsl.Hue)/60, 2)-1))
	m := l - c/2
	h := hsl.Hue

	if h >= 0 && h < 60 {
		r, g, b = c, x, 0
	} else if h >= 60 && h < 120 {
		r, g, b = x, c, 0
	} else if h >= 120 && h < 180 {
		r, g, b = 0, c, x
	} else if h >= 180 && h < 240 {
		r, g, b = 0, x, c
	} else if h >= 240 && h < 300 {
		r, g, b = x, 0, c
	} else if h >= 300 && h < 360 {
		r, g, b = c, 0, x
	}

	return &Rgb{int(math.Round((r + m) * 255)), int(math.Round((g + m) * 255)), int(math.Round((b + m) * 255))}
}

// ToHsv performs a conversion from HSL to HSV
func (hsl *Hsl) ToHsv() *Hsv {
	rgb := hsl.ToRgb()
	return rgb.ToHsv()
}

// ToCmyk performs a conversion from HSL to CMYK
func (hsl *Hsl) ToCmyk() *Cmyk {
	rgb := hsl.ToRgb()
	return rgb.ToCmyk()
}

// ToHex performs a conversion from HSL to HEX
func (hsl *Hsl) ToHex() *Hex {
	rgb := hsl.ToRgb()
	return rgb.ToHex()
}
