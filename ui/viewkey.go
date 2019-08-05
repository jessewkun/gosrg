package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/ui/base"

	"github.com/atotto/clipboard"
	"github.com/jessewkun/gocui"
)

var kView *KeyView

type KeyView struct {
	base.GView
}

func init() {
	kView = new(KeyView)
	kView.Name = "key"
	kView.Title = " Keys "
	kView.ShortCuts = []base.ShortCut{
		base.ShortCut{Key: gocui.KeyArrowUp, Level: base.SC_LOCAL_Y, Handler: kView.up},
		base.ShortCut{Key: gocui.KeyArrowDown, Level: base.SC_LOCAL_Y, Handler: kView.down},
		base.ShortCut{Key: gocui.MouseLeft, Level: base.SC_LOCAL_Y, Handler: kView.click},
		base.ShortCut{Key: gocui.KeyBackspace2, Level: base.SC_LOCAL_Y, Handler: kView.delete},
		base.ShortCut{Key: gocui.KeyCtrlF, Level: base.SC_LOCAL_Y, Handler: kView.filter},
		base.ShortCut{Key: gocui.KeyCtrlR, Level: base.SC_LOCAL_Y, Handler: kView.refresh},
		base.ShortCut{Key: gocui.KeyCtrlB, Level: base.SC_LOCAL_Y, Handler: kView.begin},
		base.ShortCut{Key: gocui.KeyCtrlE, Level: base.SC_LOCAL_Y, Handler: kView.end},
		base.ShortCut{Key: gocui.KeyCtrlY, Level: base.SC_LOCAL_Y, Handler: kView.copy},
	}
}

func (k *KeyView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(k.Name, 0, maxY/10+1, maxX/3-15, maxY-2, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = k.Title
		v.Wrap = true
		k.View = v
		k.Initialize()
	}
	return nil
}

func (k *KeyView) Initialize() error {
	redis.R.Exec("keys", "")
	k.View.Title = " Keys " + redis.R.Pattern + " "
	return nil
}

func (k *KeyView) formatOutput(argv interface{}) {
	if keys, ok := argv.([]string); ok {
		k.Clear()
		k.CursorBegin()
		l := len(keys)
		for i, key := range keys {
			if i+1 == l {
				kView.Output(key)
			} else {
				kView.Outputln(key)
			}
		}
	} else {
		opView.error("argv does not contain a variable of type []string")
	}
}

func (k *KeyView) Focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.Output(config.TipsMap[k.Name])
	// return k.click(Ui.G, k.View)
	return nil
}

func (k *KeyView) up(g *gocui.Gui, v *gocui.View) error {
	if err := k.CursorUp(); err != nil {
		return err
	}
	return k.click(g, v)
}

func (k *KeyView) down(g *gocui.Gui, v *gocui.View) error {
	if err := k.CursorDown(); err != nil {
		return err
	}
	return k.click(g, v)
}

func (k *KeyView) click(g *gocui.Gui, v *gocui.View) error {
	if key := k.GetCurrentLine(); key != "" {
		if key == redis.R.CurrentKey {
			return nil
		}
		iView.Clear()
		redis.R.GetKey(key)
	}

	return nil
}

func (k *KeyView) delete(g *gocui.Gui, v *gocui.View) error {
	if key := k.GetCurrentLine(); key != "" {
		redis.R.CurrentKey = key
		return kdView.Layout(g)
	}
	return nil
}

func (k *KeyView) filter(g *gocui.Gui, v *gocui.View) error {
	return kfView.Layout(g)
}

func (k *KeyView) refresh(g *gocui.Gui, v *gocui.View) error {
	k.Initialize()
	return k.click(g, v)
}

func (k *KeyView) begin(g *gocui.Gui, v *gocui.View) error {
	k.CursorBegin()
	return k.click(g, v)
}

func (k *KeyView) end(g *gocui.Gui, v *gocui.View) error {
	k.CursorEnd(false)
	return k.click(g, v)
}

func (k *KeyView) copy(g *gocui.Gui, v *gocui.View) error {
	if err := clipboard.WriteAll(k.GetCurrentLine()); err != nil {
		opView.error(err.Error())
		return err
	}
	opView.info("Copy key success")
	return nil
}
