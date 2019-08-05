package base

import (
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

type Modal struct {
	GView
	TabSelf    bool
	FocusIndex int
	Buttons    []*ButtonWidget
	Form       *Form
}

func (m *Modal) Tab(g *gocui.Gui, v *gocui.View) error {
	if m.Form != nil && !m.Form.IsTabEnd() {
		m.Form.Tab()
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
		_, err := m.G.SetCurrentView(nextViewName)
		if err != nil {
			utils.Error.Println(err)
			return err
		} else if nextViewName == m.Name && m.Form != nil {
			m.Form.Tab()
		}
	}
	return nil
}

func (m *Modal) HideModal(g *gocui.Gui, v *gocui.View) error {
	m.UnbindShortCuts()
	for _, b := range m.Buttons {
		b.UnbindShortCuts()
		if err := m.G.DeleteView(b.Name); err != nil {
			utils.Error.Println(err)
			return err
		}
	}
	if err := m.G.DeleteView(m.Name); err != nil {
		utils.Error.Println(err)
		return err
	}
	m.Buttons = []*ButtonWidget{}
	m.FocusIndex = 0
	m.Form = nil
	return nil
}

func (m *Modal) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	m.HideModal(g, v)
	utils.Info.Println(2)
	utils.Info.Println(m)
	return nil
}

func (m *Modal) CancelHandler(g *gocui.Gui, v *gocui.View) error {
	m.HideModal(g, v)
	utils.Info.Println(1)
	utils.Info.Println(m)
	return nil
}

func (m *Modal) NewBtns(bi ButtonInterfacer) {
	maxX, maxY := m.G.Size()
	confirm := NewButtonWidget(m.G, "confirm", maxX/3-5, maxY/3-1, "CONFIRM", bi.ConfirmHandler)
	cancel := NewButtonWidget(m.G, "cancel", maxX/3+5, maxY/3-1, "CANCEL", bi.CancelHandler)
	m.Buttons = []*ButtonWidget{confirm, cancel}
}

func (m *Modal) InitBtn(bi ButtonInterfacer) error {
	bi.NewBtns(bi)
	for _, b := range m.Buttons {
		b.Layout(m.G)
	}
	m.G.SetCurrentView(m.Buttons[0].Name)
	return nil
}
