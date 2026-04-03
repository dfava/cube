// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package internal

import "fmt"

var printInColors bool

type Color int

const (
	zero   Color = iota // the "zero color"
	green               // green
	white               // white
	orange              // orange
	red                 // red
	yellow              // yellow
	blue                // blue
)

func (c Color) Abs() Color {
	if c < 0 {
		return -c
	}
	return c
}

// using ANSI escape codes for colors
func (c Color) String() string {
	var names [7]string
	if printInColors {
		names = [...]string{" ", "\033[32mg\033[0m", "\033[37mw\033[0m", "\033[35mo\033[0m", "\033[31mr\033[0m", "\033[33my\033[0m", "\033[34mb\033[0m"} // no orange, using magenta instead
	} else {
		names = [...]string{" ", "g", "w", "o", "r", "y", "b"}
	}
	var str string
	if c < 0 {
		str = "-"
		c = -c
	}
	return str + names[c]
}

var stringToColor = map[string]Color{
	" ":                zero,
	"g":                green,
	"w":                white,
	"o":                orange,
	"r":                red,
	"y":                yellow,
	"b":                blue,
	"\033[32mg\033[0m": green,
	"\033[37mw\033[0m": white,
	"\033[35mo\033[0m": orange,
	"\033[31mr\033[0m": red,
	"\033[33my\033[0m": yellow,
	"\033[34mb\033[0m": blue,
}

func ParseColor(str string) (Color, error) {
	c, ok := stringToColor[str]
	if !ok {
		return c, fmt.Errorf("ParseColor %s", str)
	}
	return c, nil
}

func PrintInColors(b bool) {
	printInColors = b
}
