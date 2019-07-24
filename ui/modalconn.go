package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"
	"strings"

	"github.com/jessewkun/gocui"
)

var connView *ConnView

type ConnView struct {
	GView
}

func init() {
	connView = new(ConnView)
	connView.Name = "conn"
	connView.Title = " Create new redis conn "
	connView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyEsc, Level: GLOBAL_Y, Handler: connView.hide},
		ShortCut{Key: gocui.KeyTab, Level: GLOBAL_Y, Handler: connView.tab},
	}
}

func (c *ConnView) Layout(g *gocui.Gui) error {
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

func (c *ConnView) initialize() error {
	gView.unbindShortCuts()
	c.btn()
	c.setCurrent(c)
	c.bindShortCuts()
	return nil
}

func (c *ConnView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[c.Name])
	return nil
}

func (c *ConnView) tab(g *gocui.Gui, v *gocui.View) error {
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

func (c *ConnView) hide(g *gocui.Gui, v *gocui.View) error {
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

func (c *ConnView) btn() error {
	maxX, maxY := Ui.G.Size()
	confirmBtn = NewButtonWidget("confirmconn", maxX/3-5, maxY/3-1, "CONFIRM", func(g *gocui.Gui, v *gocui.View) error {
		str := utils.Trim(c.View.ViewBuffer())
		temp := [4]string{}
		temp2 := strings.Split(str, ":")
		if len(temp2) != 2 && len(temp2) != 3 && len(temp2) != 4 {
			opView.info("The parameteris incorrect, please use the colon to splicing the host, port, password and pattern")
			return nil
		}
		if temp2[0] == redis.R.Host && temp2[1] == redis.R.Port {
			opView.info("The new conn is same as the current conn")
			return nil
		}
		for i, v := range temp2 {
			temp[i] = v
		}
		if err := redis.InitRedis(temp[0], temp[1], temp[2], temp[3]); err != nil {
			opView.error(err.Error())
		} else {
			kView.clear()
			iView.clear()
			opView.clear()
			opView.info("connect to " + temp[0] + ":" + temp[1] + " success")
			Ui.NextView = sView
			c.hide(g, v)
			kView.initialize()
		}
		return nil
	})
	cancelBtn = NewButtonWidget("cancelconn", maxX/3+5, maxY/3-1, "CANCEL", func(g *gocui.Gui, v *gocui.View) error {
		c.hide(g, v)
		return nil
	})
	confirmBtn.Layout(Ui.G)
	cancelBtn.Layout(Ui.G)

	Ui.G.SetCurrentView(confirmBtn.Name)

	return nil
}
