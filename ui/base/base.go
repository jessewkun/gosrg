package base

import (
	"fmt"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

type GInterfacer interface {
	Layout(g *gocui.Gui) error
	SetG(g *gocui.Gui)
	Initialize() error
	BindShortCuts() error
	Focus() error
	SetCurrent(v GInterfacer) error
	Output(arg interface{}) error
	Outputln(arg interface{}) error
}

type GView struct {
	G         *gocui.Gui
	Name      string
	Title     string
	View      *gocui.View
	ShortCuts []ShortCut
}

const (
	SC_GLOBAL_N = iota // global and cannnot unbind
	SC_GLOBAL_Y        // global and can unbind
	SC_LOCAL_N         // local and cannot unbind
	SC_LOCAL_Y         // local and can unbind
)

type ShortCut struct {
	Key     interface{}
	Level   int
	Handler func(*gocui.Gui, *gocui.View) error
}

func (gv *GView) SetG(g *gocui.Gui) {
	gv.G = g
}

func (gv *GView) Layout(g *gocui.Gui) error {
	return nil
}

func (gv *GView) BindShortCuts() error {
	for _, sc := range gv.ShortCuts {
		vName := gv.Name
		if sc.Level == SC_GLOBAL_Y || sc.Level == SC_GLOBAL_N {
			vName = ""
		}
		if err := gv.G.SetKeybinding(vName, sc.Key, gocui.ModNone, sc.Handler); err != nil {
			utils.Error.Println(err)
			return err
		}
	}
	return nil
}

func (gv *GView) UnbindShortCuts() error {
	for _, sc := range gv.ShortCuts {
		if sc.Level == SC_GLOBAL_N || sc.Level == SC_LOCAL_N {
			continue
		}
		vName := gv.Name
		if sc.Level == SC_GLOBAL_Y {
			vName = ""
		}
		if err := gv.G.DeleteKeybinding(vName, sc.Key, gocui.ModNone); err != nil {
			utils.Error.Println(err)
			return err
		}
	}
	return nil
}

func (gv *GView) Initialize() error {
	return nil
}

func (gv *GView) Clear() error {
	gv.View.Clear()
	return nil
}

func (gv *GView) Output(arg interface{}) error {
	if _, err := fmt.Fprint(gv.View, arg); err != nil {
		utils.Error.Println(err)
		return err
	}

	return nil
}

func (gv *GView) Outputln(arg interface{}) error {
	if _, err := fmt.Fprintln(gv.View, arg); err != nil {
		utils.Error.Println(err)
		return err
	}

	return nil
}

func (gv *GView) Focus() error {
	gv.G.Cursor = false
	return nil
}

func (gv *GView) SetCurrent(v GInterfacer) error {
	if _, err := gv.G.SetCurrentView(gv.Name); err != nil {
		utils.Error.Println(err)
		return err
	}
	v.Focus()
	return nil
}

func (gv *GView) GetCurrentLine() string {
	var line string
	var err error

	_, cy := gv.View.Cursor()
	if line, err = gv.View.Line(cy); err != nil {
		utils.Error.Println(err)
		return ""
	}
	return line
}

func (gv *GView) DeleteCursorLine() error {
	cx, cy := gv.View.Cursor()
	ox, oy := gv.View.Origin()
	if err := gv.View.DeleteLine(cy); err != nil {
		utils.Error.Println(err)
		return err
	}
	if cy > 0 {
		return gv.SetCursor(cx, cy-1)
	} else if oy > 0 {
		return gv.View.SetOrigin(ox, oy-1)
	}
	return nil
}
