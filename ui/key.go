package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"
	"strings"

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
		ShortCut{Key: gocui.KeyCtrlF, Level: LOCAL_Y, Handler: kView.filter},
		ShortCut{Key: gocui.KeyCtrlR, Level: LOCAL_Y, Handler: kView.refresh},
		ShortCut{Key: gocui.KeyCtrlB, Level: LOCAL_Y, Handler: kView.begin},
		ShortCut{Key: gocui.KeyCtrlE, Level: LOCAL_Y, Handler: kView.end},
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
		k.View = v
		k.initialize()
	}
	return nil
}

func (k *KeyView) initialize() error {
	k.clear()
	k.cursorBegin()
	if output, keys := redis.R.Keys(); len(output) > 0 {
		opView.formatOutput(output)
		for i, key := range keys {
			if i+1 == len(keys) {
				kView.output(key)
			} else {
				kView.outputln(key)
			}
		}
		k.View.Title = " Keys " + redis.R.Pattern + " "
	}
	return nil
}

func (k *KeyView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[k.Name])
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
		if key == redis.R.CurrentKey {
			return nil
		}
		if output, detail, info := redis.R.KeyDetail(key); len(output) > 0 {
			opView.formatOutput(output)
			dView.formatOutput(detail)
			iView.clear()
			for _, v := range info {
				iView.outputln(utils.Yellow(strings.ToLower(v[0])+":") + v[1])
				// iView.outputln("    " + v[1])
			}
		}
	}

	return nil
}

func (k *KeyView) delete(g *gocui.Gui, v *gocui.View) error {
	if key := k.getCurrentLine(); key != "" {
		redis.R.CurrentKey = key
		return kdView.Layout(g)
	}
	return nil
}

func (k *KeyView) filter(g *gocui.Gui, v *gocui.View) error {
	return kfView.Layout(g)
}

func (k *KeyView) refresh(g *gocui.Gui, v *gocui.View) error {
	k.initialize()
	return k.click(g, v)
}

func (k *KeyView) begin(g *gocui.Gui, v *gocui.View) error {
	k.cursorBegin()
	return k.click(g, v)
}

func (k *KeyView) end(g *gocui.Gui, v *gocui.View) error {
	k.cursorEnd(false)
	return k.click(g, v)
}
