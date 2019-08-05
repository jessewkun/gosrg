package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/ui/base"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var kfView *KeyFilterView

type KeyFilterView struct {
	base.Modal
}

func init() {
	kfView = new(KeyFilterView)
	kfView.Name = "keyfilter"
	kfView.Title = " key filter "
	kfView.TabSelf = true
	kfView.ShortCuts = []base.ShortCut{
		base.ShortCut{Key: gocui.KeyEsc, Level: base.SC_GLOBAL_Y, Handler: kfView.CancelHandler},
		base.ShortCut{Key: gocui.KeyTab, Level: base.SC_GLOBAL_Y, Handler: kfView.Tab},
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
		f := new(base.Form)
		v.Editor = gocui.EditorFunc(f.Edit)
		kf.Form = f
		kf.View = v
		f.Modal = &kf.Modal
		kf.SetG(g)
		kf.Initialize()
	}
	return nil
}

func (kf *KeyFilterView) Initialize() error {
	gView.UnbindShortCuts()
	kf.InitBtn(kf)
	kf.SetCurrent(kf)
	kf.setForm()
	kf.BindShortCuts()
	return nil
}

func (kf *KeyFilterView) setForm() {
	kf.Form.MarginTop = 1
	kf.Form.MarginLeft = 2
	kf.Form.LabelAlign = base.ALIGN_RIGHT
	kf.Form.LabelColor = utils.C_GREEN
	kf.Form.SetInput("PATTERN", "pattern", redis.R.Pattern)
	kf.Form.InitForm()
}

func (kf *KeyFilterView) Focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	kf.Output(redis.R.Pattern)
	kf.CursorEnd(true)
	tView.Output(config.TipsMap[kf.Name])
	return nil
}

func (kf *KeyFilterView) CancelHandler(g *gocui.Gui, v *gocui.View) error {
	kf.HideModal(g, v)
	gView.BindShortCuts()
	utils.Info.Println(Ui.NextView)
	return Ui.NextView.SetCurrent(Ui.NextView)
}

func (kf *KeyFilterView) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	res := kf.Form.Submit()
	if len(res["pattern"]) == 0 {
		res["pattern"] = "*"
	}
	if res["pattern"] == redis.R.Pattern {
		opView.info("pattern has no change")
		return kf.CancelHandler(g, v)
	}
	redis.R.Pattern = res["pattern"]
	kView.Initialize()
	kView.click(g, v)
	kf.CancelHandler(g, v)
	return nil
}
