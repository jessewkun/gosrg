package ui

import (
	"gosrg/redis"
	"gosrg/ui/base"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var gView *GLobalView

type GLobalView struct {
	base.GView
}

func init() {
	gView = new(GLobalView)
	gView.ShortCuts = []base.ShortCut{
		base.ShortCut{Key: gocui.KeyCtrlC, Level: base.SC_GLOBAL_N, Handler: gView.quit},
		base.ShortCut{Key: 'h', Level: base.SC_GLOBAL_N, Handler: gView.showHelp},
		base.ShortCut{Key: gocui.KeyTab, Level: base.SC_GLOBAL_Y, Handler: gView.tab},
		base.ShortCut{Key: gocui.KeyCtrlD, Level: base.SC_GLOBAL_Y, Handler: gView.showDb},
		base.ShortCut{Key: gocui.KeyCtrlN, Level: base.SC_GLOBAL_Y, Handler: gView.newConn},
		base.ShortCut{Key: gocui.KeyCtrlT, Level: base.SC_GLOBAL_Y, Handler: gView.showCommand},
	}
}

func (gl *GLobalView) quit(g *gocui.Gui, v *gocui.View) error {
	redis.R.Send(redis.RES_EXIT, 0)
	return gocui.ErrQuit
}

func (gl *GLobalView) tab(g *gocui.Gui, v *gocui.View) error {
	setNextView()
	if err := Ui.NextView.SetCurrent(Ui.NextView); err != nil {
		utils.Error.Println(err)
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

func (gl *GLobalView) newConn(g *gocui.Gui, v *gocui.View) error {
	return connView.Layout(g)
}
