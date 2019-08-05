package ui

import (
	"gosrg/ui/base"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var tView *TipView

type TipView struct {
	base.GView
}

func init() {
	tView = new(TipView)
	tView.Name = "tip"
}

func (t *TipView) Layout(g *gocui.Gui) error {
	maxX, maxY := Ui.G.Size()
	if v, err := g.SetView(t.Name, 0, maxY-2, maxX-20, maxY, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Frame = false
		t.View = v
	}
	return nil
}

func (t *TipView) Output(arg interface{}) error {
	t.Clear()
	if str, ok := arg.(string); ok {
		return t.GView.Output(utils.Bule(str))
	}
	return nil
}
