package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

var KeyView = &config.View{
	Name:         "key",
	Title:        " Keys ",
	InitHandler:  KeyInitHandler,
	FocusHandler: KeyFocusHandler,
	BlurHandler:  KeyBlurHandler,
	ShortCuts: []config.ShortCut{
		config.ShortCut{Key: gocui.KeyArrowUp, Mod: gocui.ModNone, Handler: KeyUpHandler},
		config.ShortCut{Key: gocui.KeyArrowDown, Mod: gocui.ModNone, Handler: KeyDownHandler},
		config.ShortCut{Key: gocui.MouseLeft, Mod: gocui.ModNone, Handler: KeyDetailHandler},
	},
}

func KeyInitHandler() error {
	config.Srg.AllView["key"].View.Clear()
	redis.Keys()
	return nil
}

func KeyFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = true
	utils.Toutput(config.TipsMap["key"])
	// 暂时关闭 key view 的 KeyDetail, 因为要看 info 的时候必须经过 key, 如果开启的话就会覆盖掉 detail 了
	// if key := getCurrentLine(config.Srg.AllView["key"].View); key != "" {
	// 	redis.KeyDetail(key)
	// }
	return nil
}

func KeyBlurHandler() error {
	return nil
}

func KeyUpHandler(g *gocui.Gui, v *gocui.View) error {
	if err := up(v); err != nil {
		return err
	}
	if key := getCurrentLine(v); key != "" {
		redis.KeyDetail(key)
	}
	return nil
}

func KeyDownHandler(g *gocui.Gui, v *gocui.View) error {
	if err := down(v); err != nil {
		return err
	}
	if key := getCurrentLine(v); key != "" {
		redis.KeyDetail(key)
	}
	return nil
}

func KeyDetailHandler(g *gocui.Gui, v *gocui.View) error {
	if key := getCurrentLine(v); key != "" {
		redis.KeyDetail(key)
	}

	return nil
}
