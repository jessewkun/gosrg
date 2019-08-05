package ui

import (
	"gosrg/config"
	"gosrg/ui/base"

	"github.com/jessewkun/gocui"
)

var hView *HelpView

type HelpView struct {
	base.GView
}

func init() {
	hView = new(HelpView)
	hView.Name = "help"
	hView.Title = " Help "
	hView.ShortCuts = []base.ShortCut{
		base.ShortCut{Key: gocui.KeyEsc, Level: base.SC_LOCAL_N, Handler: hView.hide},
		base.ShortCut{Key: gocui.KeyArrowUp, Level: base.SC_LOCAL_Y, Handler: hView.up},
		base.ShortCut{Key: gocui.KeyArrowDown, Level: base.SC_LOCAL_Y, Handler: hView.down},
	}
}

func (h *HelpView) Layout(g *gocui.Gui) error {
	maxX, maxY := Ui.G.Size()
	if v, err := Ui.G.SetView(h.Name, maxX/3-20, maxY/3-10, maxX/2+50, maxY/2+15, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = h.Title
		v.Wrap = true
		h.View = v
		h.SetG(g)
		h.Initialize()
	}
	return nil
}

func (h *HelpView) Initialize() error {
	gView.UnbindShortCuts()
	h.SetCurrent(h)
	h.BindShortCuts()
	h.Output(config.HELP_CONTENT)
	return nil
}

func (h *HelpView) Focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.Output(config.TipsMap[h.Name])
	return nil
}

func (h *HelpView) hide(g *gocui.Gui, v *gocui.View) error {
	if err := Ui.G.DeleteView(h.Name); err != nil {
		return err
	}
	h.UnbindShortCuts()
	gView.BindShortCuts()
	Ui.NextView.SetCurrent(Ui.NextView)
	return nil
}

func (h *HelpView) up(g *gocui.Gui, v *gocui.View) error {
	return h.CursorUp()
}

func (h *HelpView) down(g *gocui.Gui, v *gocui.View) error {
	return h.CursorDown()
}
