// +build linux unix darwin

package log

import (
	"fmt"
)

const (
	BLACK       = 30
	DARK_BLUE   = 34
	DEEP_GREEN  = 32
	DARK_INDIGO = 36
	DARK_RED    = 31
	DARK_PINK   = 35
	DARK_YELLOW = 33
	LIGHT_WHITE = 37

	GRAY   = 38
	BLUE   = 42
	GREEN  = 40
	INDIGO = 44
	RED    = 39
	PINK   = 43
	YELLOW = 41
	WHITE  = 45
)

func init() {
	Cprintf = func(color uint8, format string, v ...interface{}) (int, error) {
		if color > 37 {
			color = color - 8
			fmt.Printf("\033[0;%d;1m", color)
			n, err := fmt.Printf(format, v...)
			fmt.Printf("\033[0m")
			return n, err
		} else {
			fmt.Printf("\033[1;%d;1m", color)
			n, err := fmt.Printf(format, v...)
			fmt.Printf("\033[0m")
			return n, err
		}

	}
	Cprintln = func(color uint8, v ...interface{}) (int, error) {
		if color > 37 {
			color = color - 8
			fmt.Printf("\033[0;%d;1m", color)
			n, err := fmt.Println(v...)
			fmt.Printf("\033[0m")
			return n, err
		} else {
			fmt.Printf("\033[1;%d;1m", color)
			n, err := fmt.Println(v...)
			fmt.Printf("\033[0m")
			return n, err
		}

	}
}
