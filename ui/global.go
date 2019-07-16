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
		ShortCut{Key: gocui.KeyCtrlC, Level: GLOABL_N, Handler: gView.quit},
		ShortCut{Key: gocui.KeySpace, Level: GLOABL_N, Handler: gView.showHelp},
		ShortCut{Key: gocui.KeyTab, Level: GLOABL_Y, Handler: gView.tab},
		ShortCut{Key: gocui.KeyCtrlD, Level: GLOABL_Y, Handler: gView.showDb},
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
