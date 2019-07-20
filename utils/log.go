package utils

import (
	"gosrg/config"
	"log"
	"os"
)

var (
	Command *log.Logger
	Result  *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
)

func InitLog(logPath string) {
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	Command = log.New(file, "[COMMAND] ", log.LstdFlags|log.Lshortfile)
	Result = log.New(file, "[RESULT] ", log.LstdFlags|log.Lshortfile)
	Info = log.New(file, "[INFO] ", log.LstdFlags|log.Lshortfile)
	Error = log.New(file, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	Debug = log.New(file, "[DEBUG] ", log.LstdFlags|log.Lshortfile)
}

func Exit(err interface{}) {
	Error.Println(err)
	if config.DEBUG {
		Error.Fatalln(err)
	} else {
		log.Println(err)
		os.Exit(0)
	}
}
