package ui

import (
	"gosrg/config"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var tView *TipView

type TipView struct {
	GView
}

func init() {
	tView = new(TipView)
	tView.Name = "tip"
}

func (t *TipView) Layout(g *gocui.Gui) error {
	maxX, maxY := Ui.G.Size()
	if v, err := g.SetView(t.Name, 0, maxY-2, maxX-20, maxY, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Frame = false
		t.View = v
		t.initialize()
	}
	return nil
}

func (t *TipView) initialize() error {
	t.output(config.TipsMap[t.Name])
	return nil
}

func (t *TipView) output(arg ...interface{}) error {
	t.clear()
	if str, ok := arg[0].(string); ok {
		return t.GView.output(utils.Bule(str))
	}
	return nil
}
