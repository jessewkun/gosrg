package ui

import (
	"gosrg/utils"
	"math"
	"strings"

	"github.com/jessewkun/gocui"
)

const (
	ALIGN_LEFT = iota
	ALIGN_MIDDLE
	ALIGN_RIGHT
)

const (
	TYPE_TEXT             = "text"
	TYPE_PASSWORD         = "password"
	PAD_STR               = " "
	LABEL_COLON           = " : "
	DEFAULT_CURSOR_MARGIN = 2
)

type Form struct {
	marginLeft int
	marginTop  int
	labelAlign int
	labelColor int
	MaxLabel   int
	modal      *Modal
	Input      []*Input
}

type Input struct {
	Name string
	Type string
}

func (f *Form) tab(g *gocui.Gui, v *gocui.View) error {

	return nil
}

func (f *Form) SetInput(name string, itype string) {
	f.Input = append(f.Input, &Input{Name: name, Type: itype})
	l := len(name)
	if l > f.MaxLabel {
		f.MaxLabel = l
	}
}

func (i *Input) padLabel(f *Form) {
	padLen := f.MaxLabel - len(i.Name)
	if padLen > 0 {
		switch f.labelAlign {
		case ALIGN_LEFT:
			i.Name += strings.Repeat(PAD_STR, padLen)
		case ALIGN_MIDDLE:
			if padLen%2 != 0 {
				t := int(math.Floor(float64(padLen / 2)))
				left := strings.Repeat(PAD_STR, t)
				right := strings.Repeat(PAD_STR, t+1)
				i.Name = left + i.Name + right
			} else {
				t := strings.Repeat(PAD_STR, padLen/2)
				i.Name = t + i.Name + t
			}
		case ALIGN_RIGHT:
			i.Name = strings.Repeat(PAD_STR, padLen) + i.Name
		}
	}
	if f.marginLeft > 0 {
		i.Name = strings.Repeat(PAD_STR, f.marginLeft) + i.Name
	}
}

func (f *Form) initTop(v GView) {
	if f.marginTop > 0 {
		for i := 0; i < f.marginTop; i++ {
			v.outputln("")
		}
	}
}

func (f *Form) initInput(v GView) {
	l := len(f.Input)
	for k, item := range f.Input {
		item.padLabel(f)
		t := utils.Bold(item.Name + LABEL_COLON)
		if f, ok := utils.ColorFunMap[f.labelColor]; ok {
			t = f(t)
		}
		if k+1 == l {
			v.output(t)
		} else {
			v.outputln(t)
		}
	}
}

func (f *Form) initCursor(v GView) {
	v.setCursor(f.MaxLabel+len(LABEL_COLON)+DEFAULT_CURSOR_MARGIN, f.marginTop)
}

func (f *Form) initForm(v GView) error {
	f.initTop(v)
	f.initInput(v)
	f.initCursor(v)
	return nil
}

func (f *Form) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	l := len(f.Input) - 1
	switch {
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeyTab:
		_, cy := v.Cursor()
		utils.Info.Println(cy)
		utils.Info.Println()
		if cy < l {
			NextLine := cy + 1
			NextLineStr, _ := v.Line(NextLine)
			fView.setCursor(len(NextLineStr), NextLine)
		} else {
			f.modal.tab(Ui.G, v)
		}
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		cx, _ := v.Cursor()
		utils.Info.Println(cx)
		if cx > f.MaxLabel+DEFAULT_CURSOR_MARGIN+len(LABEL_COLON) {
			v.EditDelete(true)
		}
	}
}
