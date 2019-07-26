package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"
	"strings"

	"github.com/jessewkun/gocui"
)

var cView *CommandView

type CommandView struct {
	GView
}

func init() {
	cView = new(CommandView)
	cView.Name = "command"
	cView.Title = " Command "
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
	c.btn()
	c.setCurrent(c)
	c.bindShortCuts()
	return nil
}

func (c *CommandView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[c.Name])
	return nil
}

func (c *CommandView) tab(g *gocui.Gui, v *gocui.View) error {
	nextViewName := ""
	currentView := Ui.G.CurrentView().Name()
	if currentView == c.Name {
		nextViewName = confirmBtn.Name
	} else if currentView == confirmBtn.Name {
		nextViewName = cancelBtn.Name
	} else {
		nextViewName = c.Name
	}
	if _, err := Ui.G.SetCurrentView(nextViewName); err != nil {
		utils.Error.Println(err)
		return err
	}
	return nil
}

func (c *CommandView) hide(g *gocui.Gui, v *gocui.View) error {
	c.unbindShortCuts()
	gView.bindShortCuts()
	if err := Ui.G.DeleteView(confirmBtn.Name); err != nil {
		utils.Error.Println(err)
		return err
	}
	if err := Ui.G.DeleteView(cancelBtn.Name); err != nil {
		utils.Error.Println(err)
		return err
	}
	if err := Ui.G.DeleteView(c.Name); err != nil {
		utils.Error.Println(err)
		return err
	}
	Ui.NextView.setCurrent(Ui.NextView)
	return nil
}

func (c *CommandView) btn() error {
	maxX, maxY := Ui.G.Size()
	confirmBtn = NewButtonWidget("confirfilter", maxX/3-5, maxY/3-1, "CONFIRM", func(g *gocui.Gui, v *gocui.View) error {
		str := utils.Trim(c.View.ViewBuffer())
		if str == "" {
			opView.error("The command is incorrect")
			return nil
		}
		argv := strings.Split(str, " ")
		if _, err := redis.R.CommandIsExisted(argv[0]); err != nil {
			opView.error(err.Error())
			return nil
		}
		content := ""
		if len(argv) > 1 {
			content = strings.Join(argv[1:], " ")
		}
		redis.R.Exec(argv[0], content)
		opView.formatOutput()
		dView.formatOutput()
		c.hide(g, v)
		return nil
	})
	cancelBtn = NewButtonWidget("cancelfilter", maxX/3+5, maxY/3-1, "CANCEL", func(g *gocui.Gui, v *gocui.View) error {
		c.hide(g, v)
		return nil
	})
	confirmBtn.Layout(Ui.G)
	cancelBtn.Layout(Ui.G)

	return nil
}
