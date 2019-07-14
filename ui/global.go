package ui

import (
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

var gView *GLobalView

type GLobalView struct {
	GView
}

func init() {
	gView = new(GLobalView)
	gView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyCtrlC, Mod: gocui.ModNone, Handler: gView.quit},
		ShortCut{Key: gocui.KeyTab, Mod: gocui.ModNone, Handler: gView.tab},
		ShortCut{Key: gocui.KeySpace, Mod: gocui.ModNone, Handler: gView.showHelp},
		ShortCut{Key: gocui.KeyCtrlD, Mod: gocui.ModNone, Handler: gView.showDb},
	}
}

func (gl *GLobalView) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (gl *GLobalView) tab(g *gocui.Gui, v *gocui.View) error {
	setNextView()
	if err := Ui.NextView.setCurrent(); err != nil {
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
