package ui

import (
	"gosrg/config"
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

func setNextView() {
	config.Srg.TabNo++
	next := config.Srg.TabNo % len(config.TabView)
	config.Srg.NextView = config.Srg.AllView[config.TabView[next]]
}

func getCurrentLine(v *gocui.View) string {
	var line string
	var err error

	_, cy := v.Cursor()
	if line, err = v.Line(cy); err != nil {
		utils.Logger.Println(err)
		return ""
	}
	return line
}

func up(v *gocui.View) error {
	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := v.SetOrigin(ox, oy-1); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func down(v *gocui.View) error {
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy+1); err != nil {
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func InitConfigAllView() {
	config.Srg.AllView = map[string]*config.View{
		"server":  ServerView,
		"key":     KeyView,
		"detail":  DetailView,
		"output":  OutputView,
		"tip":     TipView,
		"project": ProjectView,
		"help":    HelpView,
		"db":      DbView,
	}
	config.Srg.NextView = ServerView
}
