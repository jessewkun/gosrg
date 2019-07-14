package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"
	"strconv"
	"strings"

	"github.com/awesome-gocui/gocui"
)

var dbView *DbView

type DbView struct {
	GView
}

func init() {
	dbView = new(DbView)
	dbView.Name = "db"
	dbView.Title = " Redis Database "
	dbView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyEsc, Mod: gocui.ModNone, Handler: dbView.hide},
		ShortCut{Key: gocui.KeyArrowUp, Mod: gocui.ModNone, Handler: dbView.up},
		ShortCut{Key: gocui.KeyArrowDown, Mod: gocui.ModNone, Handler: dbView.down},
		ShortCut{Key: gocui.MouseLeft, Mod: gocui.ModNone, Handler: dbView.choice},
		ShortCut{Key: gocui.KeyEnter, Mod: gocui.ModNone, Handler: dbView.enter},
	}

}

func (db *DbView) Layout(g *gocui.Gui) error {
	maxX, maxY := Ui.G.Size()
	if v, err := Ui.G.SetView(db.Name, maxX/3-10, maxY/3-6, maxX/2+40, maxY/2+6, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = db.Title
		v.Wrap = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		db.View = v
		db.setCurrent()
	}
	return nil
}

func (db *DbView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	utils.Debug(1)
	tView.output(config.TipsMap[db.Name])
	for index := 0; index <= config.REDIS_MAX_DB_NUM; index++ {
		db.outputln("> database " + strconv.Itoa(index))
	}
	return nil
}

func (db *DbView) hide(g *gocui.Gui, v *gocui.View) error {
	if err := Ui.G.DeleteView(db.Name); err != nil {
		return err
	}
	Ui.NextView.setCurrent()
	sView.initialize()
	kView.initialize()
	return nil
}

func (db *DbView) up(g *gocui.Gui, v *gocui.View) error {
	return db.cursorUp()
}

func (db *DbView) down(g *gocui.Gui, v *gocui.View) error {
	return db.cursorDown()
}

func (db *DbView) enter(g *gocui.Gui, v *gocui.View) error {
	return db.choice(g, v)
}

func (db *DbView) choice(g *gocui.Gui, v *gocui.View) error {
	if str := db.getCurrentLine(); str != "" {
		tmp := strings.Split(str, " ")
		dbNo, _ := strconv.Atoi(tmp[2])
		if err := redis.Db(dbNo); err == nil {
			redis.R.Db = dbNo
			utils.Logger.Println("select db to " + tmp[2])
		}
		return db.hide(g, v)
	}

	return nil
}
