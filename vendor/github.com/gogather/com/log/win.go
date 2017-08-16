// +build windows
// +build !nocolor

package log

/*
#if defined _WIN32

#include <stdio.h>
#include <stdlib.h>
#include <windows.h>

CONSOLE_SCREEN_BUFFER_INFO csbi;

void set_console_color(WORD wAttributes)
{
    HANDLE hConsole = GetStdHandle(STD_OUTPUT_HANDLE);
    if (hConsole == INVALID_HANDLE_VALUE)
        return ;

	GetConsoleScreenBufferInfo(hConsole, &csbi );
    SetConsoleTextAttribute(hConsole, wAttributes);

}

void reset_color()
{
    HANDLE hConsole = GetStdHandle(STD_OUTPUT_HANDLE);
    SetConsoleTextAttribute(hConsole, csbi.wAttributes);
}

#else

void set_console_color(WORD wAttributes){}
void void reset_color(){}

#endif

*/
import "C"
import (
	"fmt"
	"sync"
)

const (
	BLACK       = 0x0
	DARK_BLUE   = 0x1
	DEEP_GREEN  = 0x2
	DARK_INDIGO = 0x3
	DARK_RED    = 0x4
	DARK_PINK   = 0x5
	DARK_YELLOW = 0x6
	LIGHT_WHITE = 0x7

	GRAY   = 0x8
	BLUE   = 0x9
	GREEN  = 0xA
	INDIGO = 0xB
	RED    = 0xC
	PINK   = 0xD
	YELLOW = 0xE
	WHITE  = 0xF
)

var (
	mutex = new(sync.Mutex)
)

func init() {
	Cprintf = func(color uint8, format string, v ...interface{}) (int, error) {
		mutex.Lock()
		defer mutex.Unlock()
		C.set_console_color(C.WORD(color))
		n, err := fmt.Printf(format, v...)
		C.reset_color()
		return n, err
	}
	Cprintln = func(color uint8, v ...interface{}) (int, error) {
		mutex.Lock()
		defer mutex.Unlock()
		C.set_console_color(C.WORD(color))
		n, err := fmt.Println(v...)
		C.reset_color()
		return n, err
	}
}
