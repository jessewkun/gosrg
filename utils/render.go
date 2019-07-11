package utils

import (
	"fmt"
	"gosrg/config"
	"strconv"
	"time"

	"github.com/awesome-gocui/gocui"
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

func Clear(v *gocui.View) {
	v.Clear()
}

func Soutput(str string) {
	v := config.Srg.AllView["server"].View
	if _, err := fmt.Fprintln(v, str); err != nil {
		Logger.Fatalln(err)
	}
}

// Douput used for Detail view output without newline
func Douput(str string) {
	v := config.Srg.AllView["detail"].View
	Clear(v)
	if _, err := fmt.Fprint(v, str); err != nil {
		Logger.Fatalln(err)
	}
}

func OCommandOuput(str string) {
	v := config.Srg.AllView["output"].View
	if _, err := fmt.Fprintln(v, Now()+Bule("[COMMAND]")+str); err != nil {
		Logger.Fatalln(err)
	}
}

func OInfoOuput(str string) {
	v := config.Srg.AllView["output"].View
	if _, err := fmt.Fprintln(v, Now()+Green("[RESULT]")+str); err != nil {
		Logger.Fatalln(err)
	}
}

func OErrorOuput(str string) {
	v := config.Srg.AllView["output"].View
	if _, err := fmt.Fprintln(v, Now()+Red("[ERROR]")+str); err != nil {
		Logger.Fatalln(err)
	}
}

func Kouput(str string) {
	v := config.Srg.AllView["key"].View
	if _, err := fmt.Fprintln(v, str); err != nil {
		Logger.Fatalln(err)
	}
}

func Toutput(str string) {
	v := config.Srg.AllView["tip"].View
	v.Clear()
	if _, err := fmt.Fprint(v, str); err != nil {
		Logger.Fatalln(err)
	}
}

func Poutput(str string) {
	v := config.Srg.AllView["project"].View
	n, err := fmt.Fprint(v, UnderLine(str))
	if err != nil {
		Logger.Fatalln(err)
	} else {
		Debug("Poutput: " + strconv.Itoa(n))
	}
}

func Houtput(str string) {
	v := config.Srg.AllView["help"].View
	n, err := fmt.Fprint(v, str)
	if err != nil {
		Logger.Fatalln(err)
	} else {
		Debug("Houtput: " + strconv.Itoa(n))
	}
}

func DBoutput(str string) {
	v := config.Srg.AllView["db"].View
	n, err := fmt.Fprintln(v, str)
	if err != nil {
		Logger.Fatalln(err)
	} else {
		Debug("DBoutput: " + strconv.Itoa(n))
	}
}

func Now() string {
	nowTime := time.Now()
	t := nowTime.String()
	timeStr := "[" + t[:19] + "] "
	return timeStr
}
