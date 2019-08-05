package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/ui/base"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var kdView *KeyDelView

type KeyDelView struct {
	base.Modal
}

func init() {
	kdView = new(KeyDelView)
	kdView.Name = "keydel"
	kdView.Title = " Delete key "
	kdView.TabSelf = false
	kdView.ShortCuts = []base.ShortCut{
		base.ShortCut{Key: gocui.KeyEsc, Level: base.SC_GLOBAL_Y, Handler: kdView.Hide},
		base.ShortCut{Key: gocui.KeyTab, Level: base.SC_GLOBAL_Y, Handler: kdView.Tab},
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
		kd.SetG(g)
		kd.Initialize()
	}
	return nil
}

func (kd *KeyDelView) Focus(arg ...interface{}) error {
	kd.G.Cursor = false
	if tip, ok := config.TipsMap[kd.Name]; ok {
		tView.Output(tip)
	} else {
		tView.Clear()
	}
	return nil
}

func (kd *KeyDelView) Hide(g *gocui.Gui, v *gocui.View) error {
	kd.Modal.HideModal(g, v)
	gView.BindShortCuts()
	return Ui.NextView.SetCurrent(Ui.NextView)
}

func (kd *KeyDelView) Initialize() error {
	gView.UnbindShortCuts()
	kd.SetCurrent(kd)
	kd.InitBtn(kd)
	kd.BindShortCuts()
	kd.Outputln("")
	kd.Outputln(utils.Red("     Confirm delete `" + redis.R.CurrentKey + "` ?"))
	return nil
}

func (kd *KeyDelView) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	redis.R.Exec("del", "")
	kView.DeleteCursorLine()
	kView.click(g, v)
	kd.Hide(g, v)
	return nil
}
