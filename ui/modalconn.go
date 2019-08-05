package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/ui/base"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var connView *ConnView

type ConnView struct {
	base.Modal
}

func init() {
	connView = new(ConnView)
	connView.Name = "conn"
	connView.Title = " Create new redis connection "
	connView.TabSelf = true
	connView.ShortCuts = []base.ShortCut{
		base.ShortCut{Key: gocui.KeyEsc, Level: base.SC_GLOBAL_Y, Handler: connView.CancelHandler},
		base.ShortCut{Key: gocui.KeyTab, Level: base.SC_GLOBAL_Y, Handler: connView.Tab},
		base.ShortCut{Key: gocui.KeyArrowUp, Level: base.SC_LOCAL_Y, Handler: connView.up},
		base.ShortCut{Key: gocui.KeyArrowDown, Level: base.SC_LOCAL_Y, Handler: connView.down},
	}
}

func (c *ConnView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(c.Name, maxX/3-10, maxY/3-6, maxX/2+40, maxY/2-3, 0); err != nil {
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

func (c *ConnView) Initialize() error {
	gView.UnbindShortCuts()
	c.InitBtn(c)
	c.SetCurrent(c)
	c.setForm()
	c.BindShortCuts()
	c.up(Ui.G, c.View)
	return nil
}

func (c *ConnView) NewBtns(bi base.ButtonInterfacer) {
	maxX, maxY := Ui.G.Size()
	confirm := base.NewButtonWidget(c.G, "confirm", maxX/3-5, maxY/3+1, "CONFIRM", c.ConfirmHandler)
	cancel := base.NewButtonWidget(c.G, "cancel", maxX/3+5, maxY/3+1, "CANCEL", c.CancelHandler)
	c.Buttons = []*base.ButtonWidget{confirm, cancel}
}

func (c *ConnView) setForm() {
	c.Form.MarginTop = 1
	c.Form.MarginLeft = 2
	c.Form.LabelAlign = base.ALIGN_RIGHT
	c.Form.LabelColor = utils.C_GREEN
	c.Form.SetInput("HOST", "host", "")
	c.Form.SetInput("PORT", "port", "")
	c.Form.SetInput("PASSWORD", "pwd", "")
	c.Form.SetInput("PATTERN", "pattern", "")
}

func (c *ConnView) Focus() error {
	Ui.G.Cursor = true
	tView.Output(config.TipsMap[c.Name])
	return nil
}

func (c *ConnView) up(g *gocui.Gui, v *gocui.View) error {
	l := len(redis.R.History)
	if redis.R.Current > 0 && redis.R.Current <= l {
		redis.R.Current--
		c.Form.SetInputValue(redis.R.History[redis.R.Current])
		c.Form.InitForm()
	}
	return nil
}

func (c *ConnView) down(g *gocui.Gui, v *gocui.View) error {
	l := len(redis.R.History)
	if redis.R.Current+1 < l {
		redis.R.Current++
		c.Form.SetInputValue(redis.R.History[redis.R.Current])
		c.Form.InitForm()
	}
	return nil
}

func (c *ConnView) CancelHandler(g *gocui.Gui, v *gocui.View) error {
	redis.R.ResetCurrent()
	c.HideModal(g, v)
	gView.BindShortCuts()
	return Ui.NextView.SetCurrent(Ui.NextView)
}

func (c *ConnView) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	res := c.Form.Submit()
	if res["host"] == redis.R.Host && res["port"] == redis.R.Port {
		opView.info("The new conn is same as the current conn")
		return nil
	}
	if err := redis.InitRedis(res["host"], res["port"], res["password"], res["pattern"]); err != nil {
		opView.error(err.Error())
	} else {
		kView.Clear()
		opView.Clear()
		RestNextView()
		c.CancelHandler(g, v)
		sView.refresh(g, sView.View)
		kView.Initialize()
	}
	return nil
}
