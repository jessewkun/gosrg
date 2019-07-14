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

func InitSrg() {
	Srg = Gosrg{
		Host: "127.0.0.1",
		Port: "6379",
	}
}
