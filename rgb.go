package colors

import (
	"bytes"
	"fmt"
	"math"
)

// Rgb represents the RGB model
type Rgb struct {
	Red   int
	Green int
	Blue  int
}

// ToHex performs a conversion from RGB model to Hexadecimal one.
func (rgb *Rgb) ToHex() *Hex {
	var hex bytes.Buffer
	hex.WriteByte('#')
	hex.WriteString(fmt.Sprintf("%.2x", rgb.Red))
	hex.WriteString(fmt.Sprintf("%.2x", rgb.Green))
	hex.WriteString(fmt.Sprintf("%.2x", rgb.Blue))
	return &Hex{hex.String()}
}

// ToHsv performs a conversion from RGB model to HSV one.
func (rgb *Rgb) ToHsv() *Hsv {
	var h, s, v float64

	r1 := float64(rgb.Red) / 255.0
	g1 := float64(rgb.Green) / 255.0
	b1 := float64(rgb.Blue) / 255.0
	cmax := math.Max(r1, math.Max(g1, b1))
	cmin := math.Min(r1, math.Min(g1, b1))
	delta := cmax - cmin

	if delta == 0 {
		h = 0
	} else if cmax == r1 {
		h = math.Round(60 * (g1 - b1) / delta)
	} else if cmax == g1 {
		h = math.Round(60 * (2 + ((b1 - r1) / delta)))
	} else if cmax == b1 {
		h = math.Round(60 * (4 + ((r1 - g1) / delta)))
	}

	if h < 0 {
		h += 360
	}

	if cmax == 0 {
		s = 0
	} else {
		s = math.Round(delta / cmax * 100)
	}

	v = math.Round(cmax * 100)

	return &Hsv{int(h), int(s), int(v)}
}

// ToHsl performs a conversion from RGB to HSL
func (rgb *Rgb) ToHsl() *Hsl {
	var h, s, l float64

	r1 := float64(rgb.Red) / 255.0
	g1 := float64(rgb.Green) / 255.0
	b1 := float64(rgb.Blue) / 255.0
	cmax := math.Max(r1, math.Max(g1, b1))
	cmin := math.Min(r1, math.Min(g1, b1))

	l = math.Round((cmin + cmax) * 50)

	if cmax == cmin {
		s = 0
	} else {
		if l/100 < 0.5 {
			s = math.Round((cmax - cmin) / (cmax + cmin) * 100)
		} else {
			s = math.Round((cmax - cmin) / (2 - cmax - cmin) * 100)
		}
	}

	if cmax == cmin {
		h = 0
	} else if cmax == r1 {
		h = math.Round(60 * (g1 - b1) / (cmax - cmin))
	} else if cmax == g1 {
		h = math.Round(60 * (2 + (b1-r1)/(cmax-cmin)))
	} else if cmax == b1 {
		h = math.Round(60 * (4 + (r1-g1)/(cmax-cmin)))
	}

	if h < 0 {
		h += 360
	}

	return &Hsl{int(h), int(s), int(l)}
}

// ToCmyk performs a conversion from RGB to CMYK
func (rgb *Rgb) ToCmyk() *Cmyk {
	var c, m, y, k float64

	r1 := float64(rgb.Red) / 255.0
	g1 := float64(rgb.Green) / 255.0
	b1 := float64(rgb.Blue) / 255.0

	k = 1 - math.Max(r1, math.Max(g1, b1))

	if k == 1 {
		c, m, y = 0, 0, 0
	} else {
		c = math.Round(100 * (1 - r1 - k) / (1 - k))

		m = math.Round(100 * (1 - g1 - k) / (1 - k))

		y = math.Round(100 * (1 - b1 - k) / (1 - k))
	}

	return &Cmyk{int(c), int(m), int(y), int(math.Round(k * 100))}
}
