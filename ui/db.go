package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"
	"strconv"
	"strings"

	"github.com/jessewkun/gocui"
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
		ShortCut{Key: gocui.KeyEsc, Level: LOCAL_N, Handler: dbView.hide},
		ShortCut{Key: gocui.KeyArrowUp, Level: LOCAL_Y, Handler: dbView.up},
		ShortCut{Key: gocui.KeyArrowDown, Level: LOCAL_Y, Handler: dbView.down},
		ShortCut{Key: gocui.MouseLeft, Level: LOCAL_Y, Handler: dbView.choice},
		ShortCut{Key: gocui.KeyEnter, Level: LOCAL_Y, Handler: dbView.enter},
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
		db.View = v
		db.initialize()
	}
	return nil
}

func (db *DbView) initialize() error {
	gView.unbindShortCuts()
	db.bindShortCuts()
	db.setCurrent(db)
	for i := 0; i <= config.REDIS_MAX_DB_NUM; i++ {
		if i == config.REDIS_MAX_DB_NUM {
			db.output("> database " + strconv.Itoa(i))
		} else {
			db.outputln("> database " + strconv.Itoa(i))
		}
	}
	db.setCursor(0, redis.R.Db)
	return nil
}

func (db *DbView) hide(g *gocui.Gui, v *gocui.View) error {
	if err := Ui.G.DeleteView(db.Name); err != nil {
		return err
	}
	db.unbindShortCuts()
	gView.bindShortCuts()
	Ui.NextView.setCurrent(Ui.NextView)
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
		if output := redis.R.SelectDb(dbNo); len(output) > 0 {
			redis.R.Db = dbNo
			opView.formatOutput(output)
			utils.Logger.Println("select db to " + tmp[2])
		}
		return db.hide(g, v)
	}

	return nil
}
