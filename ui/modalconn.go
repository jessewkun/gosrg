package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var connView *ConnView

type ConnView struct {
	Modal
}

func init() {
	connView = new(ConnView)
	connView.Name = "conn"
	connView.Title = " Create new redis connection "
	connView.TabSelf = true
	connView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyEsc, Level: GLOBAL_Y, Handler: connView.CancelHandler},
		ShortCut{Key: gocui.KeyTab, Level: GLOBAL_Y, Handler: connView.tab},
		ShortCut{Key: gocui.KeyArrowUp, Level: LOCAL_Y, Handler: connView.up},
		ShortCut{Key: gocui.KeyArrowDown, Level: LOCAL_Y, Handler: connView.down},
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
		f := new(Form)
		v.Editor = gocui.EditorFunc(f.Edit)
		c.form = f
		c.View = v
		f.modal = &c.Modal
		c.initialize()
	}
	return nil
}

func (c *ConnView) initialize() error {
	gView.unbindShortCuts()
	c.btn(c)
	c.setCurrent(c)
	c.setForm()
	c.bindShortCuts()
	c.up(Ui.G, c.View)
	return nil
}

func (c *ConnView) setForm() {
	c.form.marginLeft = 2
	c.form.labelAlign = ALIGN_RIGHT
	c.form.labelColor = utils.C_GREEN
	c.form.setInput("HOST", "host", "")
	c.form.setInput("PORT", "port", "")
	c.form.setInput("PASSWORD", "pwd", "")
	c.form.setInput("PATTERN", "pattern", "")
}

func (c *ConnView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[c.Name])
	return nil
}

func (c *ConnView) up(g *gocui.Gui, v *gocui.View) error {
	l := len(redis.R.History)
	if redis.R.Current > 0 && redis.R.Current <= l {
		redis.R.Current--
		c.form.setInputValue(redis.R.History[redis.R.Current])
		c.form.initForm()
	}
	return nil
}

func (c *ConnView) down(g *gocui.Gui, v *gocui.View) error {
	l := len(redis.R.History)
	if redis.R.Current+1 < l {
		redis.R.Current++
		c.form.setInputValue(redis.R.History[redis.R.Current])
		c.form.initForm()
	}
	return nil
}

func (c *ConnView) CancelHandler(g *gocui.Gui, v *gocui.View) error {
	redis.R.ResetCurrent()
	c.hide(g, v)
	return nil
}

func (c *ConnView) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	res := c.form.submit()
	if res["host"] == redis.R.Host && res["port"] == redis.R.Port {
		opView.info("The new conn is same as the current conn")
		return nil
	}
	if err := redis.InitRedis(res["host"], res["port"], res["password"], res["pattern"]); err != nil {
		opView.error(err.Error())
	} else {
		kView.clear()
		opView.clear()
		RestNextView()
		c.hide(g, v)
		sView.refresh(g, sView.View)
		kView.initialize()
	}
	return nil
}
