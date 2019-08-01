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
}

func (m *Modal) tab(g *gocui.Gui, v *gocui.View) error {
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
	if _, err := Ui.G.SetCurrentView(nextViewName); err != nil {
		utils.Error.Println(err)
		return err
	}
	return nil
}

func (m *Modal) hide(g *gocui.Gui, v *gocui.View) error {
	utils.Info.Println(m)
	m.unbindShortCuts()
	gView.bindShortCuts()
	for _, b := range m.Buttons {
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
	utils.Info.Println(m)
	m.Buttons = []*ButtonWidget{}
	m.FocusIndex = 0
	return nil
}

func (m *Modal) ConfirmHandler(g *gocui.Gui, v *gocui.View) error {
	m.hide(g, v)
	return nil
}

func (m *Modal) CancelmHandler(g *gocui.Gui, v *gocui.View) error {
	m.hide(g, v)
	return nil
}

func (m *Modal) btn(ci ButtonInterfacer) error {
	utils.Info.Println(m)
	maxX, maxY := Ui.G.Size()
	m.Buttons = append(m.Buttons, NewButtonWidget("confirmconn", maxX/3-5, maxY/3-1, "CONFIRM", ci.ConfirmHandler))
	m.Buttons = append(m.Buttons, NewButtonWidget("cancelconn", maxX/3+5, maxY/3-1, "CANCEL", ci.CancelmHandler))
	for _, b := range m.Buttons {
		b.Layout(Ui.G)
	}
	Ui.G.SetCurrentView(m.Buttons[0].Name)
	utils.Info.Println(m)
	return nil
}
