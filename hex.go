package colors

import (
	"log"
	"strconv"
)

// Hex is an alias for hexadecimal values
type Hex struct {
	Hex string
}

// ToRgb performs a conversion from HEX to RGB
func (h *Hex) ToRgb() *Rgb {
	if _, s := checkHex(h.Hex); s != "" {
		log.Fatal("Error with HEX color: ", s)
	}
	r, _ := strconv.ParseInt(h.Hex[1:3], 16, 64)
	g, _ := strconv.ParseInt(h.Hex[3:5], 16, 64)
	b, _ := strconv.ParseInt(h.Hex[5:7], 16, 64)
	return &Rgb{int(r), int(g), int(b)}
}

// ToHsv performs a conversion from HEX to HSV
func (h *Hex) ToHsv() *Hsv {
	if _, s := checkHex(h.Hex); s != "" {
		log.Fatal("Error with HEX color: ", s)
	}
	rgb := h.ToRgb()
	return rgb.ToHsv()
}

// ToHsl performs a conversion from HEX to HSL
func (h *Hex) ToHsl() *Hsl {
	if _, s := checkHex(h.Hex); s != "" {
		log.Fatal("Error with HEX color: ", s)
	}
	rgb := h.ToRgb()
	return rgb.ToHsl()
}

// ToCmyk performs a conversion from HEX to CMYK
func (h *Hex) ToCmyk() *Cmyk {
	if _, s := checkHex(h.Hex); s != "" {
		log.Fatal("Error with HEX color: ", s)
	}
	rgb := h.ToRgb()
	return rgb.ToCmyk()
}
