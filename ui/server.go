package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"
	"strconv"
)

var ServerView = &config.View{
	Name:         "server",
	Title:        " Server Info ",
	InitHandler:  ServerInitHandler,
	FocusHandler: ServerFocusHandler,
	BlurHandler:  ServerBlurHandler,
}

func ServerInitHandler() error {
	utils.Soutput("Current Host: " + config.Srg.Host)
	utils.Soutput("Current Port: " + config.Srg.Port)
	utils.Soutput("Current Db  : " + strconv.Itoa(config.Srg.Db))
	return nil
}

func ServerFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = false
	config.Srg.AllView["server"].View.Clear()
	ServerInitHandler()
	utils.Toutput(config.TipsMap["server"])
	redis.Info()
	return nil
}

func ServerBlurHandler() error {
	return nil
}
