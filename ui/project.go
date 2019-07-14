package ui

import (
	"gosrg/config"
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

var ProjectView *config.View

func init() {
	ProjectView = &config.View{
		Name:         "project",
		InitHandler:  ProjectInitHandler,
		FocusHandler: ProjectFocusHandler,
		BlurHandler:  ProjectBlurHandler,
		ShortCuts: []config.ShortCut{
			config.ShortCut{Key: gocui.MouseLeft, Mod: gocui.ModNone, Handler: ProjectOpenHandler},
		},
	}
}

func ProjectInitHandler() error {
	utils.Poutput(config.PROJECT_NAME + " " + config.PROJECT_VERSION)
	return nil
}
func ProjectFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = false
	return nil
}
func ProjectBlurHandler() error {
	return nil
}

func ProjectOpenHandler(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		utils.OpenLink(config.PROJECT_URL)
	}
	return nil
}
