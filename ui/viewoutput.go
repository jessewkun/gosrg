package ui

import (
	"fmt"
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"
	"strconv"

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

func (op *OutputView) formatOutput(rtype int, argv interface{}) {
	switch rtype {
	case redis.RES_OUTPUT_COMMAND:
		op.command(argv.(string))
	case redis.RES_OUTPUT_INFO:
		op.info(argv.(string))
	case redis.RES_OUTPUT_ERROR:
		op.error(argv.(string))
	case redis.RES_OUTPUT_RES:
		switch t := argv.(type) {
		case int64:
			op.res(strconv.FormatInt(t, 10))
		case string:
			op.res(t)
		default:
			opView.error(fmt.Sprintf("Unexpected type %T\n", t))
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
