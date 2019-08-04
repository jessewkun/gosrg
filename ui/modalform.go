package ui

import (
	"gosrg/config"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var fView *FormView

type FormView struct {
	Modal
}

func init() {
	fView = new(FormView)
	fView.Name = "form"
	fView.Title = " Form "
	fView.TabSelf = true
	fView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyEsc, Level: GLOBAL_Y, Handler: fView.hide},
		ShortCut{Key: gocui.KeyTab, Level: GLOBAL_Y, Handler: fView.tab},
	}
}

func (c *FormView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(c.Name, maxX/3-10, maxY/3-6, maxX/2+40, maxY/2-3, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = c.Title
		v.Wrap = true
		v.Editable = true
		f := new(Form)
		v.Editor = gocui.EditorFunc(f.Edit)
		c.form = f
		c.View = v
		f.modal = &c.Modal
		c.initialize()
	}
	return nil
}

func (c *FormView) initialize() error {
	gView.unbindShortCuts()
	c.btn(c)
	c.setCurrent(c)
	c.setForm()
	c.bindShortCuts()
	return nil
}

func (c *FormView) setForm() {
	c.form.marginLeft = 2
	c.form.labelAlign = ALIGN_RIGHT
	c.form.labelColor = utils.C_GREEN
	c.form.SetInput("host", TYPE_TEXT)
	c.form.SetInput("port", TYPE_TEXT)
	c.form.SetInput("password", TYPE_PASSWORD)
	c.form.SetInput("pattern", TYPE_TEXT)
	c.form.initForm(c.Modal.GView)
}

func (c *FormView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[c.Name])
	return nil
}

func (c *FormView) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	opView.info("TODO")
	c.hide(g, v)
	return nil
}
