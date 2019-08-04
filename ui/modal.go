package ui

import (
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

type Modal struct {
	GView
	TabSelf    bool
	FocusIndex int
	Buttons    []*ButtonWidget
	form       *Form
}

func (m *Modal) tab(g *gocui.Gui, v *gocui.View) error {
	nextViewName := ""
	l := len(m.Buttons)
	if m.TabSelf {
		l++
	}
	nextIndex := 0
	if m.form != nil && !m.form.isCursorEnd() {
		m.form.tabInput()
	} else {
		m.FocusIndex++
		nextIndex = m.FocusIndex % l
		if m.TabSelf {
			if nextIndex == 0 {
				nextViewName = m.Name
			} else {
				nextIndex--
				nextViewName = m.Buttons[nextIndex].Name
			}
		} else {
			nextViewName = m.Buttons[nextIndex].Name
		}
		_, err := Ui.G.SetCurrentView(nextViewName)
		if err != nil {
			utils.Error.Println(err)
			return err
		} else if nextViewName == m.Name && m.form != nil {
			m.form.initCursor(m.GView)
		}
	}
	return nil
}

func (m *Modal) hide(g *gocui.Gui, v *gocui.View) error {
	m.unbindShortCuts()
	gView.bindShortCuts()
	for _, b := range m.Buttons {
		b.unbindShortCuts()
		if err := Ui.G.DeleteView(b.Name); err != nil {
			utils.Error.Println(err)
			return err
		}
	}
	if err := Ui.G.DeleteView(m.Name); err != nil {
		utils.Error.Println(err)
		return err
	}
	Ui.NextView.setCurrent(Ui.NextView)
	m.Buttons = []*ButtonWidget{}
	m.FocusIndex = 0
	return nil
}

func (m *Modal) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	m.hide(g, v)
	return nil
}

func (m *Modal) CancelHandler(g *gocui.Gui, v *gocui.View) error {
	m.hide(g, v)
	return nil
}

func (m *Modal) btn(bi ButtonInterfacer) error {
	maxX, maxY := Ui.G.Size()
	m.Buttons = append(m.Buttons, NewButtonWidget("confirm", maxX/3-5, maxY/3-1, "CONFIRM", bi.ConfirmHandler))
	m.Buttons = append(m.Buttons, NewButtonWidget("cancel", maxX/3+5, maxY/3-1, "CANCEL", bi.CancelHandler))
	for _, b := range m.Buttons {
		b.Layout(Ui.G)
	}
	Ui.G.SetCurrentView(m.Buttons[0].Name)
	return nil
}
