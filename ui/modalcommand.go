package ui

import (
	"gosrg/config"
	"gosrg/ui/base"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var cView *CommandView

type CommandView struct {
	base.Modal
}

func init() {
	cView = new(CommandView)
	cView.Name = "command"
	cView.Title = " Command "
	cView.TabSelf = true
	cView.ShortCuts = []base.ShortCut{
		base.ShortCut{Key: gocui.KeyEsc, Level: base.SC_GLOBAL_Y, Handler: cView.Hide},
		base.ShortCut{Key: gocui.KeyTab, Level: base.SC_GLOBAL_Y, Handler: cView.Tab},
	}
}

func (c *CommandView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(c.Name, maxX/3-10, maxY/3-6, maxX/2+40, maxY/2-5, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = c.Title
		v.Wrap = true
		v.Editable = true
		f := new(base.Form)
		v.Editor = gocui.EditorFunc(f.Edit)
		c.Form = f
		c.View = v
		f.Modal = &c.Modal
		c.SetG(g)
		c.Initialize()
	}
	return nil
}

func (c *CommandView) Initialize() error {
	gView.UnbindShortCuts()
	c.InitBtn(c)
	c.SetCurrent(c)
	c.setForm()
	c.BindShortCuts()
	return nil
}

func (c *CommandView) setForm() {
	c.Form.MarginTop = 1
	c.Form.MarginLeft = 2
	c.Form.LabelAlign = base.ALIGN_RIGHT
	c.Form.LabelColor = utils.C_GREEN
	c.Form.SetInput("COMMAND", "cmd", "")
	c.Form.InitForm()
}

func (c *CommandView) Focus() error {
	Ui.G.Cursor = true
	tView.Output(config.TipsMap[c.Name])
	return nil
}

func (c *CommandView) Hide(g *gocui.Gui, v *gocui.View) error {
	c.Modal.HideModal(g, v)
	gView.BindShortCuts()
	return Ui.NextView.SetCurrent(Ui.NextView)
}

func (c *CommandView) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	res := c.Form.Submit()
	opView.info("TODO " + res["cmd"])
	c.Hide(g, v)
	// str := utils.Trim(c.View.ViewBuffer())
	// if str == "" {
	// 	opView.error("The command is incorrect")
	// 	return nil
	// }
	// argv := strings.Split(str, " ")
	// if _, err := redis.R.CommandIsExisted(argv[0]); err != nil {
	// 	opView.error(err.Error())
	// 	return nil
	// }
	// content := ""
	// if len(argv) > 1 {
	// 	content = strings.Join(argv[1:], " ")
	// }
	// redis.R.Exec(argv[0], content)
	// opView.formatOutput()
	// dView.formatOutput()
	// iView.formatOutput()
	// kView.formatOutput()
	// c.hide(g, v)
	return nil
}
