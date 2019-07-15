package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

var dView *DetailView

type DetailView struct {
	GView
}

func init() {
	dView = new(DetailView)
	dView.Name = "detail"
	dView.Title = " Detail "
	dView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyCtrlS, Mod: gocui.ModNone, Handler: dView.save},
	}
}

func (d *DetailView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(d.Name, maxX/3+1, 0, maxX-1, maxY-15, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = d.Title
		v.Wrap = true
		v.Editable = true
		d.View = v
		d.initialize()
	}
	return nil
}

func (d *DetailView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[d.Name])
	return nil
}

func (d *DetailView) save(g *gocui.Gui, v *gocui.View) error {
	redis.SetKeyDetail(v.ViewBuffer())
	return nil
}

func (d *DetailView) output(arg interface{}) error {
	d.clear()
	return d.GView.output(arg)
}