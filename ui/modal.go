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
	if m.form != nil && !m.form.isTabEnd() {
		m.form.tab()
	} else {
		nextViewName := ""
		l := len(m.Buttons)
		if m.TabSelf {
			l++
		}
		m.FocusIndex++
		nextIndex := m.FocusIndex % l
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
			m.form.initCursor()
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
	m.form = nil
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

func (m *Modal) newBtns() {
	maxX, maxY := Ui.G.Size()
	confirm := NewButtonWidget("confirm", maxX/3-5, maxY/3-1, "CONFIRM", m.ConfirmHandler)
	cancel := NewButtonWidget("cancel", maxX/3+5, maxY/3-1, "CANCEL", m.CancelHandler)
	m.Buttons = []*ButtonWidget{confirm, cancel}
}

func (m *Modal) initBtn(bi ButtonInterfacer) error {
	bi.newBtns()
	for _, b := range m.Buttons {
		b.Layout(Ui.G)
	}
	Ui.G.SetCurrentView(m.Buttons[0].Name)
	return nil
}
