package utils

import (
	"fmt"
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

func Green(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m ", 82, str)
}

func Orange(str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m ", 208, str)
}

func UnderLine(str string) string {
	return fmt.Sprintf("\x1b[3%d;%dm%s\x1b[0m ", 5, 4, str)
}

func Now() string {
	nowTime := time.Now()
	t := nowTime.String()
	timeStr := "[" + t[:19] + "] "
	return timeStr
}
