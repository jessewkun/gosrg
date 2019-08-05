package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/ui/base"
	"strconv"
	"strings"

	"github.com/jessewkun/gocui"
)

var dbView *DbView

type DbView struct {
	base.GView
}

func init() {
	dbView = new(DbView)
	dbView.Name = "db"
	dbView.Title = " Redis Database "
	dbView.ShortCuts = []base.ShortCut{
		base.ShortCut{Key: gocui.KeyEsc, Level: base.SC_LOCAL_N, Handler: dbView.hide},
		base.ShortCut{Key: gocui.KeyArrowUp, Level: base.SC_LOCAL_Y, Handler: dbView.up},
		base.ShortCut{Key: gocui.KeyArrowDown, Level: base.SC_LOCAL_Y, Handler: dbView.down},
		base.ShortCut{Key: gocui.MouseLeft, Level: base.SC_LOCAL_Y, Handler: dbView.choice},
		base.ShortCut{Key: gocui.KeyEnter, Level: base.SC_LOCAL_Y, Handler: dbView.enter},
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
		db.View = v
		db.SetG(g)
		db.Initialize()
	}
	return nil
}

func (db *DbView) Initialize() error {
	gView.UnbindShortCuts()
	db.BindShortCuts()
	db.SetCurrent(db)
	for i := 0; i <= config.REDIS_MAX_DB_NUM; i++ {
		if i == config.REDIS_MAX_DB_NUM {
			db.Output("> database " + strconv.Itoa(i))
		} else {
			db.Outputln("> database " + strconv.Itoa(i))
		}
	}
	db.SetCursor(0, redis.R.Db)
	return nil
}

func (db *DbView) Focus(arg ...interface{}) error {
	db.G.Cursor = false
	if tip, ok := config.TipsMap[db.Name]; ok {
		tView.Output(tip)
	} else {
		tView.Clear()
	}
	return nil
}

func (db *DbView) hide(g *gocui.Gui, v *gocui.View) error {
	if err := Ui.G.DeleteView(db.Name); err != nil {
		return err
	}
	db.UnbindShortCuts()
	gView.BindShortCuts()
	Ui.NextView.SetCurrent(Ui.NextView)
	sView.Initialize()
	kView.Initialize()
	return nil
}

func (db *DbView) up(g *gocui.Gui, v *gocui.View) error {
	return db.CursorUp()
}

func (db *DbView) down(g *gocui.Gui, v *gocui.View) error {
	return db.CursorDown()
}

func (db *DbView) enter(g *gocui.Gui, v *gocui.View) error {
	return db.choice(g, v)
}

func (db *DbView) choice(g *gocui.Gui, v *gocui.View) error {
	if str := db.GetCurrentLine(); str != "" {
		tmp := strings.Split(str, " ")
		redis.R.Exec("select", tmp[2])
		return db.hide(g, v)
	}

	return nil
}
