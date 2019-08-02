package ui

import (
	"fmt"
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var Ui UI
var ResultChan chan map[int]interface{}

type GHandler interface {
	Layout(g *gocui.Gui) error
	initialize() error
	bindShortCuts() error
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

const (
	GLOBAL_N = iota // global and cannnot unbind
	GLOBAL_Y        // global and can unbind
	LOCAL_N         // local and cannot unbind
	LOCAL_Y         // local and can unbind
)

type ShortCut struct {
	Key     interface{}
	Level   int
	Handler func(*gocui.Gui, *gocui.View) error
}

func (gv *GView) Layout(g *gocui.Gui) error {
	return nil
}

func (gv *GView) bindShortCuts() error {
	for _, sc := range gv.ShortCuts {
		vName := gv.Name
		if sc.Level == GLOBAL_Y || sc.Level == GLOBAL_N {
			vName = ""
		}
		if err := Ui.G.SetKeybinding(vName, sc.Key, gocui.ModNone, sc.Handler); err != nil {
			utils.Error.Println(err)
			return err
		}
	}
	return nil
}

func (gv *GView) unbindShortCuts() error {
	for _, sc := range gv.ShortCuts {
		if sc.Level == GLOBAL_N || sc.Level == LOCAL_N {
			continue
		}
		vName := gv.Name
		if sc.Level == GLOBAL_Y {
			vName = ""
		}
		if err := Ui.G.DeleteKeybinding(vName, sc.Key, gocui.ModNone); err != nil {
			utils.Error.Println(err)
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
	if tip, ok := config.TipsMap[gv.Name]; ok {
		tView.output(tip)
	} else {
		tView.clear()
	}
	return nil
}

func (gv *GView) clear() error {
	gv.View.Clear()
	return nil
}

func (gv *GView) output(arg interface{}) error {
	if _, err := fmt.Fprint(gv.View, arg); err != nil {
		utils.Error.Println(err)
		return err
	}

	return nil
}

func (gv *GView) outputln(arg interface{}) error {
	if _, err := fmt.Fprintln(gv.View, arg); err != nil {
		utils.Error.Println(err)
		return err
	}

	return nil
}

func (gv *GView) setCurrent(v GHandler, arg ...interface{}) error {
	if _, err := Ui.G.SetCurrentView(gv.Name); err != nil {
		utils.Error.Println(err)
		return err
	}
	v.focus(arg...)
	return nil
}

func RestNextView() {
	Ui.TabNo = 0
	Ui.NextView = sView
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
		utils.Error.Println(err)
		return ""
	}
	return line
}

func (gv *GView) deleteCursorLine() error {
	cx, cy := gv.View.Cursor()
	ox, oy := gv.View.Origin()
	if err := gv.View.DeleteLine(cy); err != nil {
		utils.Error.Println(err)
		return err
	}
	if cy > 0 {
		return gv.setCursor(cx, cy-1)
	} else if oy > 0 {
		return gv.View.SetOrigin(ox, oy-1)
	}
	return nil
}

func InitUI() {
	Ui.AllView = map[string]GHandler{
		"global":  gView,
		"server":  sView,
		"info":    iView,
		"key":     kView,
		"detail":  dView,
		"output":  opView,
		"tip":     tView,
		"project": pView,
	}
	Ui.NextView = sView
	Ui.G.SetManager(iView, tView, pView, opView, dView, sView, kView)
	for _, item := range Ui.AllView {
		item.bindShortCuts()
	}
}

func Render() {
	for {
		res := <-ResultChan
		for rtype, item := range res {
			switch rtype {
			case redis.RES_OUTPUT_COMMAND:
				fallthrough
			case redis.RES_OUTPUT_INFO:
				fallthrough
			case redis.RES_OUTPUT_ERROR:
				fallthrough
			case redis.RES_OUTPUT_RES:
				opView.formatOutput(rtype, item)
			case redis.RES_KEYS:
				kView.formatOutput(item)
			case redis.RES_DETAIL:
				dView.formatOutput(item)
			case redis.RES_INFO:
				iView.formatOutput(item)
			case redis.RES_EXIT:
				return
			default:
				utils.Error.Println(rtype)
			}
		}
	}
}
