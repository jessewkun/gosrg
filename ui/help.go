package ui

import (
	"gosrg/config"
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

var HelpView *config.View

func init() {
	HelpView = &config.View{
		Name:         "help",
		Title:        " Help ",
		InitHandler:  HelpInitHandler,
		FocusHandler: HelpFocusHandler,
		BlurHandler:  HelpBlurHandler,
		ShortCuts: []config.ShortCut{
			config.ShortCut{Key: gocui.KeyEsc, Mod: gocui.ModNone, Handler: HelpHideHandler},
		},
	}
}

func HelpInitHandler() error {
	return nil
}
func HelpFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = false
	utils.Toutput(config.TipsMap["help"])
	utils.Houtput(config.HELP_CONTENT)
	return nil
}
func HelpBlurHandler() error {
	return nil
}

func HelpHideHandler(g *gocui.Gui, v *gocui.View) error {
	if err := config.Srg.G.DeleteView(HelpView.Name); err != nil {
		return err
	}
	setCurrent(config.Srg.NextView)
	return nil
}
