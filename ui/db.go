package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"
	"strconv"
	"strings"

	"github.com/awesome-gocui/gocui"
)

var DbView = &config.View{
	Name:         "db",
	Title:        " Select Database ",
	InitHandler:  DbInitHandler,
	FocusHandler: DbFocusHandler,
	BlurHandler:  DbBlurHandler,
	ShortCuts: []config.ShortCut{
		config.ShortCut{Key: gocui.KeyEsc, Mod: gocui.ModNone, Handler: DbHideHandler},
		config.ShortCut{Key: gocui.KeyArrowUp, Mod: gocui.ModNone, Handler: DbUpHandler},
		config.ShortCut{Key: gocui.KeyArrowDown, Mod: gocui.ModNone, Handler: DbDownHandler},
		config.ShortCut{Key: gocui.MouseLeft, Mod: gocui.ModNone, Handler: DbSelectHandler},
		config.ShortCut{Key: gocui.KeyEnter, Mod: gocui.ModNone, Handler: DbEnterHandler},
	},
}

func DbInitHandler() error {
	return nil
}
func DbFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = true
	utils.Toutput(config.TipsMap["db"])
	for index := 0; index <= config.REDIS_MAX_DB_NUM; index++ {
		utils.DBoutput("> database " + strconv.Itoa(index))
	}
	return nil
}
func DbBlurHandler() error {
	return nil
}

func DbHideHandler(g *gocui.Gui, v *gocui.View) error {
	name := config.Srg.AllView["db"].Name
	if err := config.Srg.G.DeleteView(name); err != nil {
		return err
	}
	setCurrent(config.Srg.NextView)
	KeyInitHandler()
	return nil
}

func DbUpHandler(g *gocui.Gui, v *gocui.View) error {
	return up(v)
}

func DbDownHandler(g *gocui.Gui, v *gocui.View) error {
	return down(v)
}

func DbEnterHandler(g *gocui.Gui, v *gocui.View) error {
	return DbSelectHandler(g, v)
}

func DbSelectHandler(g *gocui.Gui, v *gocui.View) error {
	if str := getCurrentLine(v); str != "" {
		tmp := strings.Split(str, " ")
		db, _ := strconv.Atoi(tmp[2])
		if err := redis.Db(db); err == nil {
			config.Srg.Db = db
			utils.Logger.Println("select db to " + tmp[2])
		}
		return DbHideHandler(g, v)
	}

	return nil
}
