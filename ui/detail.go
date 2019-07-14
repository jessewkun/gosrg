package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

var DetailView *config.View

func init() {
	DetailView = &config.View{
		Name:         "detail",
		Title:        " Detail ",
		InitHandler:  DetailInitHandler,
		FocusHandler: DetailFocusHandler,
		BlurHandler:  DetailBlurHandler,
		ShortCuts: []config.ShortCut{
			config.ShortCut{Key: gocui.KeyCtrlS, Mod: gocui.ModNone, Handler: DetailSaveHandler},
		},
	}
}

func DetailInitHandler() error {
	return nil
}
func DetailFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = true
	utils.Toutput(config.TipsMap["detail"])
	return nil
}
func DetailBlurHandler() error {
	return nil
}

func DetailSaveHandler(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		redis.SetKeyDetail(v.ViewBuffer())
	}
	return nil
}
