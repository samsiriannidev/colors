package colors

import "math"

// Hsv represents the HSV model
type Hsv struct {
	Hue        int
	Saturation int
	Value      int
}

// ToRgb performs a convertion from HSV to RGB
func (hsv *Hsv) ToRgb() *Rgb {
	var r, g, b float64

	s := float64(hsv.Saturation) / 100.0
	v := float64(hsv.Value) / 100.0
	c := v * s
	m := v - c
	x := c * (1 - (math.Abs(math.Mod(float64(hsv.Hue)/60, 2) - 1)))
	h := hsv.Hue

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

// ToHsl performs a conversion from HSV to HSL
func (hsv *Hsv) ToHsl() *Hsl {
	rgb := hsv.ToRgb()
	return rgb.ToHsl()
}

// ToCmyk performs a conversion from HSV to CMYK
func (hsv *Hsv) ToCmyk() *Cmyk {
	rgb := hsv.ToRgb()
	return rgb.ToCmyk()
}

// ToHex performs a conversion from HSV to HEX
func (hsv *Hsv) ToHex() *Hex {
	rgb := hsv.ToRgb()
	return rgb.ToHex()
}
