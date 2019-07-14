package utils

import (
	"gosrg/config"
	"log"
	"os"
)

var Logger *log.Logger

func InitLog() {
	file, err := os.OpenFile(config.LOG_FILE, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	Logger = log.New(file, "", log.LstdFlags)
	Logger.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Debug(v ...interface{}) {
	if config.DEBUG {
		Logger.Println(v)
	}
}
