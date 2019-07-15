package ui

import (
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

var iView *InfoView

type InfoView struct {
	GView
}

func init() {
	iView = new(InfoView)
	iView.Name = "info"
	iView.Title = " Key Info "
}

func (i *InfoView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(i.Name, maxX-29, 0, maxX-1, maxY-15, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = i.Title
		v.Wrap = true
		i.View = v
		i.initialize()
	}
	return nil
}
