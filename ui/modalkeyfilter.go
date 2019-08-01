package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var kfView *KeyFilterView

type KeyFilterView struct {
	Modal
}

func init() {
	kfView = new(KeyFilterView)
	kfView.Name = "keyfilter"
	kfView.Title = " key filter "
	kfView.TabSelf = true
	kfView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyEsc, Level: GLOBAL_Y, Handler: kfView.hide},
		ShortCut{Key: gocui.KeyTab, Level: GLOBAL_Y, Handler: kfView.tab},
	}
}

func (kf *KeyFilterView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(kf.Name, maxX/3-10, maxY/3-6, maxX/2+40, maxY/2-5, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = kf.Title
		v.Wrap = true
		v.Editable = true
		kf.View = v
		kf.initialize()
	}
	return nil
}

func (kf *KeyFilterView) initialize() error {
	gView.unbindShortCuts()
	kf.btn(kf)
	kf.setCurrent(kf)
	kf.bindShortCuts()
	return nil
}

func (kf *KeyFilterView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	kf.output(redis.R.Pattern)
	kf.cursorEnd(true)
	tView.output(config.TipsMap[kf.Name])
	return nil
}

func (kf *KeyFilterView) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	pattern := utils.Trim(kf.View.ViewBuffer())
	if len(pattern) == 0 {
		pattern = "*"
	}
	if pattern == redis.R.Pattern {
		opView.info("pattern has no change")
		return kf.hide(g, v)
	}
	redis.R.Pattern = pattern
	kView.initialize()
	kView.click(g, v)
	kf.hide(g, v)
	return nil
}
