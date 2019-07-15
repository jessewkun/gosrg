package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"
	"strconv"

	"github.com/awesome-gocui/gocui"
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
		ShortCut{Key: gocui.KeyArrowUp, Mod: gocui.ModNone, Handler: kView.up},
		ShortCut{Key: gocui.KeyArrowDown, Mod: gocui.ModNone, Handler: kView.down},
		ShortCut{Key: gocui.MouseLeft, Mod: gocui.ModNone, Handler: kView.click},
	}
}

func (k *KeyView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(k.Name, 0, maxY/10+1, maxX/3-15, maxY-2, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = k.Title
		v.Wrap = true
		v.Autoscroll = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
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
	// 暂时关闭 key view 的 KeyDetail, 因为要看 server info 的时候必须经过 key, 如果开启的话就会覆盖掉 detail 了
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
		if output, detail, info := redis.R.KeyDetail(key); len(output) > 0 {
			opView.formatOutput(output)
			dView.output(detail)
			iView.clear()
			for k, v := range info {
				iView.outputln(utils.Yellow(k + " :"))
				iView.outputln("    " + strconv.FormatInt(v, 10))
			}
		}
	}

	return nil
}
