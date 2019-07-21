package ui

import (
	"gosrg/utils"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/jessewkun/gocui"
)

var iView *InfoView

type InfoView struct {
	GView
}

func init() {
	iView = new(InfoView)
	iView.Name = "info"
	iView.Title = " Info "
	iView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyCtrlY, Level: LOCAL_Y, Handler: iView.copy},
	}
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
	}
	return nil
}

func (i *InfoView) copy(g *gocui.Gui, v *gocui.View) error {
	if err := clipboard.WriteAll(v.ViewBuffer()); err != nil {
		opView.error(err.Error())
		return err
	}
	opView.info("Copy key info success")
	return nil
}
