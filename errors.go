package colors

import (
	"fmt"
	"regexp"
)

func check100(n int) (bool, string) {
	if n < 0 || n > 100 {
		return false, fmt.Sprintf("%d out of range [0, 100]\n", n)
	}
	return true, ""
}

func check360(n int) (bool, string) {
	if n < 0 || n > 359 {
		return false, fmt.Sprintf("%d out of range [0, 359]\n", n)
	}
	return true, ""
}

func check255(n int) (bool, string) {
	if n < 0 || n > 255 {
		return false, fmt.Sprintf("%d out of range [0, 255]\n", n)
	}
	return true, ""
}

func checkHex(v string) (bool, string) {
	ok, _ := regexp.Match(`^#([A-Fa-f0-9]{6})$`, []byte(v))
	if !ok {
		return ok, fmt.Sprintf("%s not correct!", v)
	}
	return ok, ""
}

func check(color interface{}, values ...int) (bool, string) {
	msg := ""
	switch color.(type) {
	case *Rgb:
		if ok, msg := check255(values[0]); !ok {
			return false, fmt.Sprintf("Red -> %s\n", msg)
		}
		if ok, msg := check255(values[1]); !ok {
			return false, fmt.Sprintf("Green -> %s\n", msg)
		}
		if ok, msg := check255(values[2]); !ok {
			return false, fmt.Sprintf("Blue -> %s\n", msg)
		}
	case *Hsv:
		if ok, msg := check360(values[0]); !ok {
			return false, fmt.Sprintf("Hue -> %s\n", msg)
		}
		if ok, msg := check100(values[1]); !ok {
			return false, fmt.Sprintf("Saturation -> %s\n", msg)
		}
		if ok, msg := check100(values[2]); !ok {
			return false, fmt.Sprintf("Value -> %s\n", msg)
		}
	case *Hsl:
		if ok, msg := check360(values[0]); !ok {
			return false, fmt.Sprintf("Hue -> %s\n", msg)
		}
		if ok, msg := check100(values[1]); !ok {
			return false, fmt.Sprintf("Saturation -> %s\n", msg)
		}
		if ok, msg := check100(values[2]); !ok {
			return false, fmt.Sprintf("Lightness -> %s\n", msg)
		}
	case *Cmyk:
		if ok, msg := check100(values[0]); !ok {
			return false, fmt.Sprintf("Cyan -> %s\n", msg)
		}
		if ok, msg := check100(values[1]); !ok {
			return false, fmt.Sprintf("Magenta -> %s\n", msg)
		}
		if ok, msg := check100(values[2]); !ok {
			return false, fmt.Sprintf("Yellow -> %s\n", msg)
		}
		if ok, msg := check100(values[3]); !ok {
			return false, fmt.Sprintf("Black -> %s\n", msg)
		}
	}
	return true, msg
}
