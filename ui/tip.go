package ui

import (
	"gosrg/config"
	"gosrg/utils"
)

var TipView = &config.View{
	Name:         "tip",
	InitHandler:  TipInitHandler,
	FocusHandler: TipFocusHandler,
	BlurHandler:  TipBlurHandler,
}

func TipInitHandler() error {
	utils.Toutput(config.TipsMap["tip"])
	return nil
}
func TipFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = false
	return nil
}
func TipBlurHandler() error {
	return nil
}
