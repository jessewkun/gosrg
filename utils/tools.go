package utils

import (
	"fmt"
	"strings"
	"time"
)

const (
	C_YELLOW = iota
	C_BULE
	C_TIANQING
	C_RED
	C_PINK
	C_GREEN
	C_ORANGE
)

var ColorFunMap map[int]func(str string) string

func init() {
	ColorFunMap = map[int]func(str string) string{
		C_YELLOW:   Yellow,
		C_BULE:     Bule,
		C_TIANQING: Tianqing,
		C_RED:      Red,
		C_PINK:     Pink,
		C_GREEN:    Green,
		C_ORANGE:   Orange,
	}
}

func Yellow(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m", 226, str)
}

func Bule(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m", 69, str)
}

func Tianqing(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m", 45, str)
}

func Red(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m", 160, str)
}

func Pink(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m", 9, str)
}

func Green(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m", 82, str)
}

func Orange(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m", 208, str)
}

func UnderLine(str string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 4, str)
}

func Bold(str string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[0m", str)
}

func Now() string {
	nowTime := time.Now()
	t := nowTime.String()
	timeStr := "[" + t[:19] + "] "
	return timeStr
}

func Trim(str string) string {
	str = strings.Trim(str, " ")
	str = strings.Trim(str, "\n")
	return str
}
