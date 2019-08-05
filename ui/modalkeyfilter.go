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
		f := new(Form)
		v.Editor = gocui.EditorFunc(f.Edit)
		kf.form = f
		kf.View = v
		f.modal = &kf.Modal
		kf.initialize()
	}
	return nil
}

func (kf *KeyFilterView) initialize() error {
	gView.unbindShortCuts()
	kf.initBtn(kf)
	kf.setCurrent(kf)
	kf.setForm()
	kf.bindShortCuts()
	return nil
}

func (kf *KeyFilterView) setForm() {
	kf.form.marginTop = 1
	kf.form.marginLeft = 2
	kf.form.labelAlign = ALIGN_RIGHT
	kf.form.labelColor = utils.C_GREEN
	kf.form.setInput("PATTERN", "pattern", "")
	kf.form.initForm()
}

func (kf *KeyFilterView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	kf.output(redis.R.Pattern)
	kf.cursorEnd(true)
	tView.output(config.TipsMap[kf.Name])
	return nil
}

func (kf *KeyFilterView) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	res := kf.form.submit()
	if len(res["pattern"]) == 0 {
		res["pattern"] = "*"
	}
	if res["pattern"] == redis.R.Pattern {
		opView.info("pattern has no change")
		return kf.hide(g, v)
	}
	redis.R.Pattern = res["pattern"]
	kView.initialize()
	kView.click(g, v)
	kf.hide(g, v)
	return nil
}
