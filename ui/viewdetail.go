package ui

import (
	"fmt"
	"gosrg/config"
	"gosrg/redis"
	"gosrg/ui/base"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/jessewkun/gocui"
)

var dView *DetailView

type DetailView struct {
	base.GView
}

func init() {
	dView = new(DetailView)
	dView.Name = "detail"
	dView.Title = " Detail (Normal mode) "
	dView.ShortCuts = []base.ShortCut{
		base.ShortCut{Key: gocui.KeyCtrlS, Level: base.SC_LOCAL_Y, Handler: dView.save},
		base.ShortCut{Key: gocui.KeyCtrlY, Level: base.SC_LOCAL_Y, Handler: dView.copy},
		base.ShortCut{Key: gocui.KeyCtrlP, Level: base.SC_LOCAL_Y, Handler: dView.paste},
		base.ShortCut{Key: gocui.KeyCtrlL, Level: base.SC_LOCAL_Y, Handler: dView.clean},
		base.ShortCut{Key: gocui.KeyCtrlB, Level: base.SC_LOCAL_Y, Handler: dView.begin},
		base.ShortCut{Key: gocui.KeyCtrlE, Level: base.SC_LOCAL_Y, Handler: dView.end},
		base.ShortCut{Key: 'i', Level: base.SC_LOCAL_Y, Handler: dView.insertMode},
		base.ShortCut{Key: gocui.KeyEsc, Level: base.SC_LOCAL_Y, Handler: dView.normalmode},
	}
}

func (d *DetailView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(d.Name, maxX/3-14, 0, maxX-30, maxY-15, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = d.Title
		v.Wrap = true
		v.Editable = false
		d.View = v
	}
	return nil
}

func (d *DetailView) Focus() error {
	Ui.G.Cursor = true
	tView.Output(config.TipsMap[d.Name])
	return nil
}

func (d *DetailView) save(g *gocui.Gui, v *gocui.View) error {
	if !d.View.Editable {
		opView.info("Saving is only worked in insert mode, pressing 'i' to switch insert mode")
		return nil
	}
	redis.R.SetKey(v.ViewBuffer())
	return nil
}

func (d *DetailView) formatOutput(detail interface{}) {
	d.Clear()
	d.CursorBegin()
	switch t := detail.(type) {
	case int64:
		d.Output(strconv.FormatInt(t, 10))
	case string:
		d.Output(t)
	case []string:
		for i, v := range t {
			if i+1 == len(t) {
				d.GView.Output(v)
			} else {
				d.GView.Outputln(v)
			}
		}
	case map[string]string:
		i := 0
		for k, v := range t {
			i++
			if i == len(t) {
				d.GView.Output(k + redis.SEPARATOR + v)
			} else {
				d.GView.Outputln(k + redis.SEPARATOR + v)
			}
		}
	default:
		opView.error(fmt.Sprintf("Unexpected type %T\n", t))
	}
}

func (d *DetailView) copy(g *gocui.Gui, v *gocui.View) error {
	if err := clipboard.WriteAll(v.ViewBuffer()); err != nil {
		opView.error(err.Error())
		return err
	}
	opView.info("Copy detail success")
	return nil
}

// paste Used for copy content from clipboard to detail view
// bug: linebreak is lost after pasted, fixed with the following code but bring a new bug: the text before last line is missing
func (d *DetailView) paste(g *gocui.Gui, v *gocui.View) error {
	if !d.View.Editable {
		opView.info("Pasting is only worked in insert mode, pressing 'i' to switch insert mode")
		return nil
	}
	text, err := clipboard.ReadAll()
	if err != nil {
		opView.error(err.Error())
		return err
	}
	line := d.GetCurrentLine()
	cx, cy := d.View.Cursor()
	newText := line[:cx] + text + line[cx:]
	d.setLine(cy, newText)
	// textArr := strings.Split(text, "\n")
	// for i, v := range textArr {
	// 	if i == 0 {
	// 		d.setLine(cy, line[:cx]+v)
	// 	}
	// }
	newCx := len(line[:cx] + text)
	// i := 0
	// newCx := cx
	// for {
	// 	_, cy := d.View.Cursor()
	// 	if i == 0 {
	// 		d.setLine(cy, line[:cx]+textArr[i])
	// 		utils.Logger.Print(cy)
	// 		utils.Logger.Print(line[:cx] + textArr[i])
	// 	} else if i+1 == len(textArr) {
	// 		d.setLine(cy, textArr[i]+line[cx:])
	// 		utils.Logger.Print(cy)
	// 		utils.Logger.Print(textArr[i] + line[cx:])
	// 		newCx = len(textArr[i])
	// 		break
	// 	} else {
	// 		d.setLine(cy, textArr[i])
	// 		utils.Logger.Print(cy)
	// 		utils.Logger.Print(textArr[i])
	// 	}
	// 	if i < len(textArr) {
	// 		d.View.EditNewLine()
	// 	}
	// 	i++
	// }
	// _, cy = d.View.Cursor()
	return d.SetCursor(newCx, cy)
}

func (d *DetailView) setLine(cy int, text string) error {
	if err := d.View.SetLine(cy, text); err != nil {
		opView.error(err.Error())
		return err
	}
	return nil
}

func (d *DetailView) clean(g *gocui.Gui, v *gocui.View) error {
	if !d.View.Editable {
		opView.info("Clearing is only worked in insert mode, pressing 'i' to switch insert mode")
		return nil
	}
	return d.Clear()
}

func (d *DetailView) begin(g *gocui.Gui, v *gocui.View) error {
	return d.CursorBegin()
}

func (d *DetailView) end(g *gocui.Gui, v *gocui.View) error {
	return d.CursorEnd(true)
}

func (d *DetailView) insertMode(g *gocui.Gui, v *gocui.View) error {
	d.View.Editable = true
	d.View.Title = " Detail (Insert mode) "
	opView.info("Switch to insert mode")
	return nil
}

func (d *DetailView) normalmode(g *gocui.Gui, v *gocui.View) error {
	d.View.Editable = false
	d.View.Title = " Detail (Normal mode) "
	opView.info("Switch to normal mode")
	return nil
}
