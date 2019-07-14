package ui

import (
	"gosrg/config"

	"github.com/awesome-gocui/gocui"
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
		ShortCut{Key: gocui.KeyEsc, Mod: gocui.ModNone, Handler: hView.hide},
	}
}

func (h *HelpView) Layout(g *gocui.Gui) error {
	maxX, maxY := Ui.G.Size()
	if v, err := Ui.G.SetView(h.Name, maxX/3-10, maxY/3-6, maxX/2+40, maxY/2+6, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = h.Title
		v.Wrap = true
		h.View = v
		h.setCurrent()
	}
	return nil
}

func (h *HelpView) focus(arg ...interface{}) error {
	Ui.G.Cursor = false
	tView.output(config.TipsMap[h.Name])
	h.output(config.HELP_CONTENT)
	return nil
}

func (h *HelpView) hide(g *gocui.Gui, v *gocui.View) error {
	if err := Ui.G.DeleteView(h.Name); err != nil {
		return err
	}
	Ui.NextView.setCurrent()
	return nil
}
