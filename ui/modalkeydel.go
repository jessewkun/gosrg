package ui

import (
	"gosrg/redis"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var kdView *KeyDelView

type KeyDelView struct {
	Modal
}

func init() {
	kdView = new(KeyDelView)
	kdView.Name = "keydel"
	kdView.Title = " Delete key "
	kdView.TabSelf = false
	kdView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyEsc, Level: GLOBAL_Y, Handler: kdView.hide},
		ShortCut{Key: gocui.KeyTab, Level: GLOBAL_Y, Handler: kdView.tab},
	}
}

func (kd *KeyDelView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(kd.Name, maxX/3-10, maxY/3-6, maxX/2+40, maxY/2-5, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = kd.Title
		v.Wrap = true
		kd.View = v
		kd.initialize()
	}
	return nil
}

func (kd *KeyDelView) initialize() error {
	gView.unbindShortCuts()
	kd.setCurrent(kd)
	kd.btn(kd)
	kd.bindShortCuts()
	kd.outputln("")
	kd.outputln(utils.Red("     Confirm delete `" + redis.R.CurrentKey + "` ?"))
	return nil
}

func (kd *KeyDelView) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	redis.R.Exec("del", "")
	opView.formatOutput()
	kView.deleteCursorLine()
	kd.hide(g, v)
	return nil
}
