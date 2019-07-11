package config

import (
	"github.com/awesome-gocui/gocui"
	"github.com/gomodule/redigo/redis"
)

type Gosrg struct {
	Host           string
	Port           string
	Pwd            string
	G              *gocui.Gui
	Redis          redis.Conn
	Db             int
	CurrentKey     string
	CurrentKeyType string
	AllView        map[string]*View
	TabNo          int
	NextView       *View
}

type View struct {
	Name         string
	Title        string
	View         *gocui.View
	InitHandler  func() error
	FocusHandler func(arg ...interface{}) error
	BlurHandler  func() error
	ShortCuts    []ShortCut
}

type ShortCut struct {
	Key     interface{}
	Mod     gocui.Modifier
	Handler func(*gocui.Gui, *gocui.View) error
}

var (
	OS      string
	Srg     Gosrg
	TabView = []string{"server", "key", "detail", "output"}
	TipsMap = map[string]string{
		"server": "Tab: Toggle view | Ctrl-c: Quit | Ctrl-space: Help",
		"key":    "↑↓ MouseLeft: Toggle keys",
		"detail": "Ctrl-s: Save detail",
		"output": "Tab: Toggle view | Ctrl-c: Quit | Ctrl-space: Help",
		"tip":    "Tab: Toggle view | Ctrl-c: Quit | Ctrl-space: Help",
		"help":   "Esc: Close Help view",
	}
)

func InitSrg() {
	Srg = Gosrg{
		Host:  "127.0.0.1",
		Port:  "6379",
		Db:    1,
		TabNo: 0,
	}
}
