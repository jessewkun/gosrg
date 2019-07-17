package utils

import (
	"fmt"
	"strings"
	"time"
)

func Yellow(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m ", 226, str)
}

func Bule(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m ", 69, str)
}

func Red(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m ", 160, str)
}

func Pink(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m ", 9, str)
}

func Green(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m ", 82, str)
}

func Orange(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m ", 208, str)
}

func UnderLine(str string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m ", 4, str)
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
