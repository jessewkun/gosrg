package ui

import (
	"gosrg/config"

	"github.com/jessewkun/gocui"
)

var hView *HelpView

type HelpView struct {
	GView
}

func init() {
	hView = new(HelpView)
	hView.Name = "help"
	hView.Title = " Help "
	hView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyEsc, Level: LOCAL_N, Handler: hView.hide},
		ShortCut{Key: gocui.KeyArrowUp, Level: LOCAL_Y, Handler: hView.up},
		ShortCut{Key: gocui.KeyArrowDown, Level: LOCAL_Y, Handler: hView.down},
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
		h.initialize()
	}
	return nil
}

func (h *HelpView) initialize() error {
	gView.unbindShortCuts()
	h.setCurrent(h)
	h.bindShortCuts()
	h.output(config.HELP_CONTENT)
	return nil
}

func (h *HelpView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[h.Name])
	return nil
}

func (h *HelpView) hide(g *gocui.Gui, v *gocui.View) error {
	if err := Ui.G.DeleteView(h.Name); err != nil {
		return err
	}
	h.unbindShortCuts()
	gView.bindShortCuts()
	Ui.NextView.setCurrent(Ui.NextView)
	return nil
}

func (h *HelpView) up(g *gocui.Gui, v *gocui.View) error {
	if err := h.cursorUp(); err != nil {
		return err
	}
	return nil
}

func (h *HelpView) down(g *gocui.Gui, v *gocui.View) error {
	if err := h.cursorDown(); err != nil {
		return err
	}
	return nil
}
