package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/ui/base"
	"strconv"

	"github.com/jessewkun/gocui"
)

var sView *ServerView

type ServerView struct {
	base.GView
}

func init() {
	sView = new(ServerView)
	sView.Name = "server"
	sView.Title = " Server "
	sView.ShortCuts = []base.ShortCut{
		base.ShortCut{Key: gocui.KeyCtrlR, Level: base.SC_LOCAL_Y, Handler: sView.refresh},
	}
}

func (s *ServerView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(s.Name, 0, 0, maxX/3-15, maxY/10, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = s.Title
		v.Wrap = true
		s.View = v
		s.SetCurrent(s)
		s.refresh(g, v)
	}
	return nil
}

func (s *ServerView) Initialize() error {
	s.Clear()
	s.Outputln("Current Host: " + redis.R.Host)
	s.Outputln("Current Port: " + redis.R.Port)
	s.Outputln("Current Db  : " + strconv.Itoa(redis.R.Db))
	return nil
}

func (s *ServerView) Focus(arg ...interface{}) error {
	Ui.G.Cursor = false
	s.Initialize()
	tView.Output(config.TipsMap[s.Name])
	return nil
}

func (s *ServerView) refresh(g *gocui.Gui, v *gocui.View) error {
	iView.Clear()
	redis.R.Exec("info", "")
	return nil
}
