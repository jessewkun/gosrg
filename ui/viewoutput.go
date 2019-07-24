package ui

import (
	"gosrg/config"
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
	opView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyArrowUp, Level: LOCAL_Y, Handler: opView.up},
		ShortCut{Key: gocui.KeyArrowDown, Level: LOCAL_Y, Handler: opView.down},
		ShortCut{Key: gocui.KeyCtrlB, Level: LOCAL_Y, Handler: opView.begin},
		ShortCut{Key: gocui.KeyCtrlE, Level: LOCAL_Y, Handler: opView.end},
	}
}

func (op *OutputView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(op.Name, maxX/3-14, maxY-14, maxX-1, maxY-2, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = op.Title
		v.Wrap = true
		op.View = v
		op.initialize()
	}
	return nil
}

func (op *OutputView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[op.Name])
	return nil
}

func (op *OutputView) command(str string) {
	op.outputln(utils.Now() + utils.Bule("[COMMAND]") + str)
	utils.Command.Println(str)
	op.cursorEnd(false)
}

func (op *OutputView) info(str string) {
	op.outputln(utils.Now() + utils.Tianqing("[INFO]") + str)
	utils.Info.Println(str)
	op.cursorEnd(false)
}

func (op *OutputView) res(str string) {
	op.outputln(utils.Now() + utils.Green("[RESULT]") + str)
	utils.Result.Println(str)
	op.cursorEnd(false)
}

func (op *OutputView) error(str string) {
	op.outputln(utils.Now() + utils.Red("[ERROR]") + str)
	utils.Error.Println(str)
	op.cursorEnd(false)
}

func (op *OutputView) formatOutput() {
	for _, item := range redis.R.Output {
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

func (op *OutputView) up(g *gocui.Gui, v *gocui.View) error {
	return op.cursorUp()
}

func (op *OutputView) down(g *gocui.Gui, v *gocui.View) error {
	return op.cursorDown()
}

func (op *OutputView) begin(g *gocui.Gui, v *gocui.View) error {
	return op.cursorBegin()
}

func (op *OutputView) end(g *gocui.Gui, v *gocui.View) error {
	return op.cursorEnd(false)
}
