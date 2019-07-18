package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"

	"github.com/atotto/clipboard"
	"github.com/jessewkun/gocui"
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
		ShortCut{Key: gocui.KeyCtrlS, Level: LOCAL_Y, Handler: dView.save},
		ShortCut{Key: gocui.KeyCtrlY, Level: LOCAL_Y, Handler: dView.copy},
		ShortCut{Key: gocui.KeyCtrlP, Level: LOCAL_Y, Handler: dView.paste},
		ShortCut{Key: gocui.KeyCtrlL, Level: LOCAL_Y, Handler: dView.clean},
		ShortCut{Key: gocui.KeyCtrlB, Level: LOCAL_Y, Handler: dView.begin},
		ShortCut{Key: gocui.KeyCtrlE, Level: LOCAL_Y, Handler: dView.end},
	}
}

func (d *DetailView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(d.Name, maxX/3-14, 0, maxX-30, maxY-15, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = d.Title
		v.Wrap = true
		v.Editable = true
		d.View = v
	}
	return nil
}

func (d *DetailView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[d.Name])
	return nil
}

func (d *DetailView) save(g *gocui.Gui, v *gocui.View) error {
	if output := redis.R.SetKeyDetail(v.ViewBuffer()); len(output) > 0 {
		opView.formatOutput(output)
	}
	return nil
}

func (d *DetailView) output(arg ...interface{}) error {
	d.clear()
	return d.GView.output(arg...)
}

func (d *DetailView) copy(g *gocui.Gui, v *gocui.View) error {
	if err := clipboard.WriteAll(v.ViewBuffer()); err != nil {
		opView.error(err.Error())
		utils.Logger.Println(err)
		return err
	}
	opView.info("copy success")
	return nil
}

func (d *DetailView) paste(g *gocui.Gui, v *gocui.View) error {
	text, err := clipboard.ReadAll()
	if err != nil {
		opView.error(err.Error())
		utils.Logger.Println(err)
		return err
	}
	return dView.output(text)
}

func (d *DetailView) clean(g *gocui.Gui, v *gocui.View) error {
	return d.clear()
}

func (d *DetailView) begin(g *gocui.Gui, v *gocui.View) error {
	return d.cursorBegin()
}

func (d *DetailView) end(g *gocui.Gui, v *gocui.View) error {
	return d.cursorEnd(true)
}
