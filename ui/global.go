package ui

import (
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var gView *GLobalView

type GLobalView struct {
	GView
}

func init() {
	gView = new(GLobalView)
	gView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyCtrlC, Level: GLOBAL_N, Handler: gView.quit},
		ShortCut{Key: 'h', Level: GLOBAL_N, Handler: gView.showHelp},
		ShortCut{Key: gocui.KeyTab, Level: GLOBAL_Y, Handler: gView.tab},
		ShortCut{Key: gocui.KeyCtrlD, Level: GLOBAL_Y, Handler: gView.showDb},
		// ShortCut{Key: gocui.KeyCtrlE, Level: GLOBAL_Y, Handler: gView.showCommand},
	}
}

func (gl *GLobalView) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (gl *GLobalView) tab(g *gocui.Gui, v *gocui.View) error {
	setNextView()
	if err := Ui.NextView.setCurrent(Ui.NextView); err != nil {
		utils.Logger.Fatalln(err)
		return err
	}
	return nil
}

func (gl *GLobalView) showHelp(g *gocui.Gui, v *gocui.View) error {
	return hView.Layout(g)
}

func (gl *GLobalView) showDb(g *gocui.Gui, v *gocui.View) error {
	return dbView.Layout(g)
}

func (gl *GLobalView) showCommand(g *gocui.Gui, v *gocui.View) error {
	return cView.Layout(g)
}
