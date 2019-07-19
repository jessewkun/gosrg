package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"

	"github.com/atotto/clipboard"
	"github.com/jessewkun/gocui"
)

var dView *DetailView

type DetailView struct {
	GView
}

func init() {
	dView = new(DetailView)
	dView.Name = "detail"
	dView.Title = " Detail (Normal mode) "
	dView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyCtrlS, Level: LOCAL_Y, Handler: dView.save},
		ShortCut{Key: gocui.KeyCtrlY, Level: LOCAL_Y, Handler: dView.copy},
		ShortCut{Key: gocui.KeyCtrlP, Level: LOCAL_Y, Handler: dView.paste},
		ShortCut{Key: gocui.KeyCtrlL, Level: LOCAL_Y, Handler: dView.clean},
		ShortCut{Key: gocui.KeyCtrlB, Level: LOCAL_Y, Handler: dView.begin},
		ShortCut{Key: 'i', Level: LOCAL_Y, Handler: dView.insertMode},
		ShortCut{Key: gocui.KeyEsc, Level: LOCAL_Y, Handler: dView.normalmode},
	}
}

func (d *DetailView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(d.Name, maxX/3-14, 0, maxX-30, maxY-15, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = d.Title
		v.Wrap = true
		v.Editable = false
		d.View = v
	}
	return nil
}

func (d *DetailView) focus(arg ...interface{}) error {
	Ui.G.Cursor = true
	tView.output(config.TipsMap[d.Name])
	return nil
}

func (d *DetailView) save(g *gocui.Gui, v *gocui.View) error {
	if !d.View.Editable {
		opView.info("Saving is only worked in insert mode, pressing 'i' to switch insert mode")
		return nil
	}
	if output := redis.R.SetKeyDetail(v.ViewBuffer()); len(output) > 0 {
		opView.formatOutput(output)
	}
	return nil
}

func (d *DetailView) output(arg interface{}) error {
	d.clear()
	return d.GView.output(arg)
}

func (d *DetailView) formatOutput(arg interface{}) error {
	d.clear()
	d.cursorBegin()
	switch t := arg.(type) {
	case string:
		d.output(t)
	case []string:
		for i, v := range t {
			if i+1 == len(t) {
				d.GView.output(v)
			} else {
				d.GView.outputln(v)
			}
		}
	case map[string]string:
		i := 0
		for k, v := range t {
			i++
			if i == len(t) {
				d.GView.output(k + ": " + v)
			} else {
				d.GView.outputln(k + ": " + v)
			}
		}
	}
	return nil
}

func (d *DetailView) copy(g *gocui.Gui, v *gocui.View) error {
	if err := clipboard.WriteAll(v.ViewBuffer()); err != nil {
		opView.error(err.Error())
		utils.Logger.Println(err)
		return err
	}
	opView.info("copy success")
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
		utils.Logger.Println(err)
		return err
	}
	line := d.getCurrentLine()
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
	return d.setCursor(newCx, cy)
}

func (d *DetailView) setLine(cy int, text string) error {
	if err := d.View.SetLine(cy, text); err != nil {
		opView.error(err.Error())
		utils.Logger.Println(err)
		return err
	}
	return nil
}

func (d *DetailView) clean(g *gocui.Gui, v *gocui.View) error {
	if !d.View.Editable {
		opView.info("Clearing is only worked in insert mode, pressing 'i' to switch insert mode")
		return nil
	}
	return d.clear()
}

func (d *DetailView) begin(g *gocui.Gui, v *gocui.View) error {
	return d.cursorBegin()
}

func (d *DetailView) end(g *gocui.Gui, v *gocui.View) error {
	return d.cursorEnd(true)
}

func (d *DetailView) insertMode(g *gocui.Gui, v *gocui.View) error {
	d.View.Editable = true
	d.View.Title = " Detail (Insert mode) "
	return nil
}

func (d *DetailView) normalmode(g *gocui.Gui, v *gocui.View) error {
	d.View.Editable = false
	d.View.Title = " Detail (Normal mode) "
	return nil
}
