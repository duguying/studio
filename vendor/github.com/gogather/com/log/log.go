package log

import (
	"fmt"
	"os"
)

const (
	_VERSION = "0.3.0807"
)

func Version() string {
	return _VERSION
}

var (
	Cprintf func(uint8, string, ...interface{}) (int, error) = func(color uint8, format string, v ...interface{}) (int, error) {
		return fmt.Printf(format, v...)
	}
	Cprintln func(uint8, ...interface{}) (int, error) = func(color uint8, v ...interface{}) (int, error) {
		return fmt.Println(v...)
	}
)

var (
	Debug bool = true
)

func Warnf(format string, v ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}
	n, err = Cprintf(YELLOW, format, v...)
	return n, err
}

func Yellowf(format string, v ...interface{}) (n int, err error) {
	return Warnf(format, v...)
}

func Dangerf(format string, v ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}
	n, err = Cprintf(RED, format, v...)
	return n, err
}

func Redf(format string, v ...interface{}) (n int, err error) {
	return Dangerf(format, v...)
}

func Finef(format string, v ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}
	n, err = Cprintf(GREEN, format, v...)
	return n, err
}

func Greenf(format string, v ...interface{}) (n int, err error) {
	return Finef(format, v...)
}

func Bluef(format string, v ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}
	n, err = Cprintf(BLUE, format, v...)
	return n, err
}

func Infof(format string, v ...interface{}) (n int, err error) {
	return Bluef(format, v...)
}

func Pinkf(format string, v ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}
	n, err = Cprintf(PINK, format, v...)
	return n, err
}

// line

func Warnln(a ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}
	n, err = Cprintln(YELLOW, a...)
	return n, err
}

func Yellowln(a ...interface{}) (n int, err error) {
	return Warnln(a...)
}

func Dangerln(a ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}

	n, err = Cprintln(RED, a...)
	return n, err
}

func Redln(a ...interface{}) (n int, err error) {
	return Dangerln(a...)
}

func Fineln(a ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}

	n, err = Cprintln(GREEN, a...)
	return n, err
}

func Greenln(a ...interface{}) (n int, err error) {
	return Fineln(a...)
}

func Blueln(a ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}

	n, err = Cprintln(BLUE, a...)
	return n, err
}

func Infoln(a ...interface{}) (n int, err error) {
	return Blueln(a...)
}

func Pinkln(a ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}

	n, err = Cprintln(PINK, a...)
	return n, err
}

// fatal

func Fatalf(format string, v ...interface{}) {
	Dangerf(format, v...)
	os.Exit(1)
}

func Fatal(v ...interface{}) {
	Dangerln(v...)
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	Dangerln(v...)
	os.Exit(1)
}

func Println(a ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}

	return fmt.Println(a...)
}

func Printf(format string, v ...interface{}) (n int, err error) {
	if !Debug {
		return 0, nil
	}

	return fmt.Printf(format, v...)
}
