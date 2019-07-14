package ui

import (
	"gosrg/config"
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

var opView *OutputView

type OutputView struct {
	GView
}

func init() {
	opView = new(OutputView)
	opView.Name = "output"
	opView.Title = " Output "
}

func (op *OutputView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(op.Name, maxX/3+1, maxY-14, maxX-1, maxY-2, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = op.Title
		v.Wrap = true
		v.Autoscroll = true
		op.View = v
		op.initialize()
	}
	return nil
}

func (op *OutputView) focus(arg ...interface{}) error {
	Ui.G.Cursor = false
	tView.output(config.TipsMap[op.Name])
	return nil
}
