package ui

import (
	"gosrg/config"
	"gosrg/utils"
)

var OutputView *config.View

func init() {
	OutputView = &config.View{
		Name:         "output",
		Title:        " Output ",
		InitHandler:  OutputInitHandler,
		FocusHandler: OutputFocusHandler,
		BlurHandler:  OutputBlurHandler,
	}
}

func OutputInitHandler() error {
	return nil
}
func OutputFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = false
	utils.Toutput(config.TipsMap["output"])
	return nil
}
func OutputBlurHandler() error {
	return nil
}
