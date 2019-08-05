package base

import (
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

const (
	ALIGN_LEFT = iota
	ALIGN_MIDDLE
	ALIGN_RIGHT
)

const (
	PAD_STR               = " "
	LABEL_COLON           = " : "
	DEFAULT_CURSOR_MARGIN = 2
)

type Form struct {
	Done       bool
	MarginLeft int
	MarginTop  int
	LabelAlign int
	LabelColor int
	MaxLabel   int
	Modal      *Modal
	Cursor     int
	Input      []*Input
}

func (f *Form) SetInput(label string, name string, value string) {
	f.Input = append(f.Input, &Input{Label: label, Name: name, Value: value})
	l := len(label)
	if l > f.MaxLabel {
		f.MaxLabel = l
	}
}

func (f *Form) initTop() {
	if f.MarginTop > 0 {
		for i := 0; i < f.MarginTop; i++ {
			f.Modal.Outputln("")
		}
	}
}

func (f *Form) initInput() {
	l := len(f.Input)
	for k, item := range f.Input {
		item.padLabel(f)
		t := utils.Bold(item.Label + LABEL_COLON)
		if f, ok := utils.ColorFunMap[f.LabelColor]; ok {
			t = f(t)
		}
		t += item.Value
		if k+1 == l {
			f.Modal.Output(t)
		} else {
			f.Modal.Outputln(t)
		}
	}
}

// The input values are only updated when submit, and the value can be modified but not submitted
func (f *Form) initCursor() {
	f.Cursor = 0
	firstLine, _ := f.Modal.View.Line(f.MarginTop)
	f.Modal.SetCursor(len(firstLine), f.MarginTop)
}

func (f *Form) InitForm() error {
	f.Modal.Clear()
	f.initTop()
	f.initInput()
	f.initCursor()
	f.Done = true
	return nil
}

func (f *Form) Tab() {
	_, cy := f.Modal.View.Cursor()
	if cy < len(f.Input) {
		NextLine := cy + 1
		NextLineStr, _ := f.Modal.View.Line(NextLine)
		f.Modal.SetCursor(len(NextLineStr), NextLine)
		f.Cursor++
	} else {
		f.initCursor()
	}
}

func (f *Form) IsTabEnd() bool {
	return f.Cursor+1 == len(f.Input)
}

func (f *Form) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeyTab:
		f.Tab()
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		cx, _ := v.Cursor()
		if cx > f.MaxLabel+DEFAULT_CURSOR_MARGIN+len(LABEL_COLON) {
			v.EditDelete(true)
		}
	case key == gocui.KeyArrowLeft:
		cx, _ := v.Cursor()
		if cx > f.MaxLabel+DEFAULT_CURSOR_MARGIN+len(LABEL_COLON) {
			v.MoveCursor(-1, 0, false)
		}
	case key == gocui.KeyArrowRight:
		cx, cy := v.Cursor()
		if cx < f.MaxLabel+DEFAULT_CURSOR_MARGIN+len(LABEL_COLON)+len(f.Input[cy].Value) {
			v.MoveCursor(1, 0, false)
		}
	}
}

func (f *Form) Reset() {
	for _, i := range f.Input {
		i.Value = ""
	}
}

func (f *Form) SetInputValue(data map[string]string) {
	if len(data) < 1 {
		f.Reset()
		return
	}
	for key, item := range data {
		for _, i := range f.Input {
			if key == i.Name {
				i.Value = item
			}
		}
	}
}

func (f *Form) Submit() map[string]string {
	l := f.MaxLabel + len(LABEL_COLON) + DEFAULT_CURSOR_MARGIN
	res := make(map[string]string)
	buf := f.Modal.View.ViewBufferLines()
	buf = buf[f.MarginTop:]
	for key, item := range f.Input {
		res[item.Name] = utils.Trim(buf[key][l:])
	}
	return res
}
