package ui

import (
	"gosrg/redis"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var opView *OutputView

type OutputView struct {
	GView
}

func init() {
	opView = new(OutputView)
	opView.Name = "output"
	opView.Title = " Output "
}

func (op *OutputView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(op.Name, maxX/3-14, maxY-14, maxX-1, maxY-2, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = op.Title
		v.Wrap = true
		v.Autoscroll = true
		op.View = v
		op.initialize()
	}
	return nil
}

func (op *OutputView) formatOutput(str [][]string) {
	for _, item := range str {
		if len(item) != 2 {
			continue
		}
		switch item[1] {
		case redis.OUTPUT_COMMAND:
			op.command(item[0])
		case redis.OUTPUT_INFO:
			op.info(item[0])
		case redis.OUTPUT_ERROR:
			op.error(item[0])
		case redis.OUTPUT_RES:
			op.res(item[0])
		case redis.OUTPUT_DEBUG:
			op.debug(item[0])
		}
	}
}
