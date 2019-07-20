package ui

import (
	"gosrg/utils"
	"strings"

	"github.com/jessewkun/gocui"
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
			return err
		}
		v.Title = i.Title
		v.Wrap = true
		i.View = v
		i.initialize()
	}
	return nil
}

func (i *InfoView) formatOuput(info [][]string) error {
	i.clear()
	for _, v := range info {
		i.outputln(utils.Yellow(strings.ToLower(v[0])+":") + v[1])
		// iView.outputln("    " + v[1])
	}
	return nil
}
