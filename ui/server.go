package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"
	"strconv"

	"github.com/awesome-gocui/gocui"
)

var sView *ServerView

type ServerView struct {
	GView
}

func init() {
	sView = new(ServerView)
	sView.Name = "server"
	sView.Title = " Server Info "
}

func (s *ServerView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(s.Name, 0, 0, maxX/3-15, maxY/10, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = s.Title
		v.Wrap = true
		s.View = v
		s.initialize()
		s.setCurrent(s)
	}
	return nil
}

func (s *ServerView) initialize() error {
	s.clear()
	s.outputln("Current Host: " + redis.R.Host)
	s.outputln("Current Port: " + redis.R.Port)
	s.outputln("Current Db  : " + strconv.Itoa(redis.R.Db))
	return nil
}

func (s *ServerView) focus(arg ...interface{}) error {
	Ui.G.Cursor = false
	s.initialize()
	iView.clear()
	tView.output(config.TipsMap[s.Name])
	if output, res := redis.R.Info(); len(output) > 0 {
		opView.formatOutput(output)
		dView.output(res)
	}
	return nil
}
