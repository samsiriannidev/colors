# colors
A golang module to work with colors.
Thanks to this module you'll be able to perform conversions between the most used color models: **RGB, HSL, HSV, CMYK, HEX**

## Prerequisites

This module requires [Golang](https://golang.org) on your computer before you can start using it!
For more info about Golang installation, [go here!](https://golang.org/dl/)

## Installation
Open the terminal and digit:

`go get https://github.com/samsiriannidev/colors`

## Usage
The package is mainly composed by four structures: `Rgb`, `Cmyk`, `Hsl`, `Hsv`, `Hex`

### Rgb

`Rgb` has three fields: 
  * `Red`, that has a range of [0, 255]
  * `Green`, that has a range of [0, 255]
  * `Blue`, that has a range of [0, 255]

### Hsl

`Hsl` has three fields: 
  * `Hue`, that has a range of [0, 359]
  * `Saturation`, that has a range of [0, 100]
  * `Lightness`, that has a range of [0, 100]
  
### Hsv

`Hsv` has three fields: 
  * `Hue`, that has a range of [0, 359]
  * `Saturation`, that has a range of [0, 100]
  * `Value`, that has a range of [0, 100]
  
### Cmyk

`Cmyk` has three fields: 
  * `Cyan`, that has a range of [0, 100]
  * `Magenta`, that has a range of [0, 100]
  * `Yellow`, that has a range of [0, 100]
  * `Black`, that has a range of [0, 100]

#### Examples

```Go
package main

import (
  "fmt"
  "github.com/samsiriannidev/colors"
)

func main() {
  rgb := colors.Rgb{220, 23, 34}
  fmt.Printf("Red: %d\n", rgb.Red)
  
  cmyk := rgb.ToCmyk()   // ToCmyk() returns *Cmyk
  hsv := cmyk.ToHsv()    // ToHsv() returns *Hsv
}
```

```Go
package main

import (
  "fmt"
  "github.com/samsiriannidev/colors"
)

func main() {
  rgbs := []colors.Rgb{
		colors.Rgb{23, 32, 200},
		colors.Rgb{222, 21, 100},
		colors.Rgb{100, 55, 245},
	}

	for _, rgb := range rgbs {
		fmt.Printf("HEX: %s\n", rgb.ToHex().Hex)
	}
}
```

## Authors

  * **Sam Sirianni**, visit [www.samsirianni.com](https://samsirianni.com)
