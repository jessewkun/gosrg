package ui

import (
	"gosrg/config"

	"github.com/jessewkun/gocui"
)

var cView *CommandView

type CommandView struct {
	Modal
}

func init() {
	cView = new(CommandView)
	cView.Name = "command"
	cView.Title = " Command "
	cView.TabSelf = true
	cView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyEsc, Level: GLOBAL_Y, Handler: cView.hide},
		ShortCut{Key: gocui.KeyTab, Level: GLOBAL_Y, Handler: cView.tab},
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
		c.View = v
		c.initialize()
	}
	return nil
}

func (c *CommandView) initialize() error {
	gView.unbindShortCuts()
	c.btn(c)
	c.setCurrent(c)
	c.bindShortCuts()
	return nil
}

func (c *CommandView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[c.Name])
	return nil
}

func (c *CommandView) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	opView.info("TODO")
	c.hide(g, v)
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
