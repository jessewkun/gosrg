package ui

import (
	"fmt"
	"gosrg/config"
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

var Ui UI

type GHandler interface {
	Layout(g *gocui.Gui) error
	registerShortCuts() error
	initialize() error
	focus(arg ...interface{}) error
	clear() error
	setCurrent(v GHandler, arg ...interface{}) error
	getCurrentLine() string
	output(arg interface{}) error
	outputln(arg interface{}) error
}

type UI struct {
	G        *gocui.Gui
	AllView  map[string]GHandler
	TabNo    int
	NextView GHandler
}

type GView struct {
	Name      string
	Title     string
	View      *gocui.View
	ShortCuts []ShortCut
}

type ShortCut struct {
	Key     interface{}
	Mod     gocui.Modifier
	Handler func(*gocui.Gui, *gocui.View) error
}

func (gv *GView) Layout(g *gocui.Gui) error {
	return nil
}

func (gv *GView) registerShortCuts() error {
	for _, sc := range gv.ShortCuts {
		if err := Ui.G.SetKeybinding(gv.Name, sc.Key, sc.Mod, sc.Handler); err != nil {
			utils.Logger.Fatalln(err)
			return err
		}
	}
	return nil
}

func (gv *GView) initialize() error {
	return nil
}

func (gv *GView) focus(arg ...interface{}) error {
	Ui.G.Cursor = false
	return nil
}

func (gv *GView) clear() error {
	gv.View.Clear()
	return nil
}

func (gv *GView) output(arg interface{}) error {
	if _, err := fmt.Fprint(gv.View, arg); err != nil {
		utils.Logger.Fatalln(err)
		return err
	}

	return nil
}

func (gv *GView) outputln(arg interface{}) error {
	if _, err := fmt.Fprintln(gv.View, arg); err != nil {
		utils.Logger.Fatalln(err)
		return err
	}

	return nil
}

func (gv *GView) setCurrent(v GHandler, arg ...interface{}) error {
	if _, err := Ui.G.SetCurrentView(gv.Name); err != nil {
		utils.Logger.Fatalln(err)
		return err
	}
	utils.Debug(fmt.Sprintf("current view: %s", gv.Name))
	v.focus(arg...)
	return nil
}

func setNextView() {
	Ui.TabNo++
	next := Ui.TabNo % len(config.TabView)
	Ui.NextView = Ui.AllView[config.TabView[next]]
}

func (gv *GView) getCurrentLine() string {
	var line string
	var err error

	_, cy := gv.View.Cursor()
	if line, err = gv.View.Line(cy); err != nil {
		utils.Logger.Println(err)
		return ""
	}
	return line
}

func (gv *GView) cursorUp() error {
	ox, oy := gv.View.Origin()
	cx, cy := gv.View.Cursor()
	if err := gv.View.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := gv.View.SetOrigin(ox, oy-1); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (gv *GView) cursorDown() error {
	cx, cy := gv.View.Cursor()
	if err := gv.View.SetCursor(cx, cy+1); err != nil {
		ox, oy := gv.View.Origin()
		if err := gv.View.SetOrigin(ox, oy+1); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func InitUI() {
	Ui.AllView = map[string]GHandler{
		"global":  gView,
		"server":  sView,
		"key":     kView,
		"detail":  dView,
		"output":  opView,
		"tip":     tView,
		"project": pView,
		"help":    hView,
		"db":      dbView,
	}
	Ui.NextView = sView
	Ui.G.SetManager(tView, pView, opView, dView, sView, kView)
	for _, item := range Ui.AllView {
		item.registerShortCuts()
	}
}
