package ui

import (
	"fmt"
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
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

func (op *OutputView) focus(arg ...interface{}) error {
	Ui.G.Cursor = false
	tView.output(config.TipsMap[op.Name])
	return nil
}

func (op *OutputView) formatOutput(str [][]string) {
	for _, item := range str {
		if len(item) != 2 {
			continue
		}
		switch item[1] {
		case redis.OUTPUT_COMMAND:
			op.commandOuput(item[0])
		case redis.OUTPUT_INFO:
			op.infoOuput(item[0])
		case redis.OUTPUT_ERROR:
			op.errorOuput(item[0])
		}
	}
}

func (op *OutputView) commandOuput(str string) {
	if _, err := fmt.Fprintln(op.View, utils.Now()+utils.Bule("[COMMAND]")+str); err != nil {
		utils.Logger.Fatalln(err)
	}
}

func (op *OutputView) infoOuput(str string) {
	if _, err := fmt.Fprintln(op.View, utils.Now()+utils.Green("[RESULT]")+str); err != nil {
		utils.Logger.Fatalln(err)
	}
}

func (op *OutputView) errorOuput(str string) {
	if _, err := fmt.Fprintln(op.View, utils.Now()+utils.Red("[ERROR]")+str); err != nil {
		utils.Logger.Fatalln(err)
	}
}
