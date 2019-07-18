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
	}
}

func (h *HelpView) Layout(g *gocui.Gui) error {
	maxX, maxY := Ui.G.Size()
	if v, err := Ui.G.SetView(h.Name, maxX/3-20, maxY/3-10, maxX/2+50, maxY/2+20, 0); err != nil {
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

func (h *HelpView) hide(g *gocui.Gui, v *gocui.View) error {
	if err := Ui.G.DeleteView(h.Name); err != nil {
		return err
	}
	h.unbindShortCuts()
	gView.bindShortCuts()
	Ui.NextView.setCurrent(Ui.NextView)
	return nil
}
