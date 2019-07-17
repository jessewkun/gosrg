package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var kView *KeyView

type KeyView struct {
	GView
}

func init() {
	kView = new(KeyView)
	kView.Name = "key"
	kView.Title = " Keys "
	kView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyArrowUp, Level: LOCAL_Y, Handler: kView.up},
		ShortCut{Key: gocui.KeyArrowDown, Level: LOCAL_Y, Handler: kView.down},
		ShortCut{Key: gocui.MouseLeft, Level: LOCAL_Y, Handler: kView.click},
		ShortCut{Key: gocui.KeyBackspace2, Level: LOCAL_Y, Handler: kView.delete},
	}
}

func (k *KeyView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(k.Name, 0, maxY/10+1, maxX/3, maxY-2, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = k.Title
		v.Wrap = true
		v.Autoscroll = true
		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		k.View = v
		k.initialize()
	}
	return nil
}

func (k *KeyView) initialize() error {
	k.clear()
	if output, keys := redis.R.Keys(); len(output) > 0 {
		opView.formatOutput(output)
		for _, key := range keys {
			kView.outputln(key)
		}
	}
	return nil
}

func (k *KeyView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[k.Name])
	if key := k.getCurrentLine(); key != "" {
		redis.R.CurrentKey = key
	}
	// 暂时关闭 key view 的 KeyDetail, 因为要看 info 的时候必须经过 key, 如果开启的话就会覆盖掉 detail 了
	// return k.click(Ui.G, k.View)
	return nil
}

func (k *KeyView) up(g *gocui.Gui, v *gocui.View) error {
	if err := k.cursorUp(); err != nil {
		return err
	}
	return k.click(g, v)
}

func (k *KeyView) down(g *gocui.Gui, v *gocui.View) error {
	if err := k.cursorDown(); err != nil {
		return err
	}
	return k.click(g, v)
}

func (k *KeyView) click(g *gocui.Gui, v *gocui.View) error {
	if key := k.getCurrentLine(); key != "" {
		if output, detail := redis.R.KeyDetail(key); len(output) > 0 {
			opView.formatOutput(output)
			dView.output(detail)
		}
	}

	return nil
}

func (k *KeyView) delete(g *gocui.Gui, v *gocui.View) error {
	if key := k.getCurrentLine(); key == "" {
		return nil
	}
	return kdView.Layout(g)
}
