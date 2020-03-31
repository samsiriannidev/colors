package colors

import (
	"log"
	"math"
)

// Cmyk represents the CMYK model.
type Cmyk struct {
	Cyan    int
	Magenta int
	Yellow  int
	Black   int
}

// ToRgb performs a conversion from CMYK to RGB
func (cmyk *Cmyk) ToRgb() *Rgb {
	if _, s := check(cmyk, cmyk.Cyan, cmyk.Magenta, cmyk.Yellow, cmyk.Black); s != "" {
		log.Fatal("Error with CMYK color: ", s)
	}
	return &Rgb{
		int(math.Round(255 * (1 - float64(cmyk.Cyan)/100) * (1 - float64(cmyk.Black)/100))),
		int(math.Round(255 * (1 - float64(cmyk.Magenta)/100) * (1 - float64(cmyk.Black)/100))),
		int(math.Round(255 * (1 - float64(cmyk.Yellow)/100) * (1 - float64(cmyk.Black)/100))),
	}
}

// ToHsv performs a conversion from CMYK to HSV
func (cmyk *Cmyk) ToHsv() *Hsv {
	if _, s := check(cmyk, cmyk.Cyan, cmyk.Magenta, cmyk.Yellow, cmyk.Black); s != "" {
		log.Fatal("Error with CMYK color: ", s)
	}
	rgb := cmyk.ToRgb()
	return rgb.ToHsv()
}

// ToHsl performs a conversion from CMYK to HSL
func (cmyk *Cmyk) ToHsl() *Hsl {
	if _, s := check(cmyk, cmyk.Cyan, cmyk.Magenta, cmyk.Yellow, cmyk.Black); s != "" {
		log.Fatal("Error with CMYK color: ", s)
	}
	rgb := cmyk.ToRgb()
	return rgb.ToHsl()
}

// ToHex performs a conversion from CMYK to HEX
func (cmyk *Cmyk) ToHex() *Hex {
	if _, s := check(cmyk, cmyk.Cyan, cmyk.Magenta, cmyk.Yellow, cmyk.Black); s != "" {
		log.Fatal("Error with CMYK color: ", s)
	}
	rgb := cmyk.ToRgb()
	return rgb.ToHex()
}
