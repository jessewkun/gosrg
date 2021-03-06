package base

import (
	"fmt"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

type ButtonWidget struct {
	G       *gocui.Gui
	Name    string
	x, y    int
	w       int
	label   string
	handler func(g *gocui.Gui, v *gocui.View) error
}

type ButtonInterfacer interface {
	ConfirmHandler(g *gocui.Gui, v *gocui.View) error
	CancelHandler(g *gocui.Gui, v *gocui.View) error
	NewBtns(bi ButtonInterfacer)
}

func NewButtonWidget(g *gocui.Gui, name string, x, y int, label string, handler func(g *gocui.Gui, v *gocui.View) error) *ButtonWidget {
	return &ButtonWidget{G: g, Name: name, x: x, y: y, w: len(label) + 1, label: label, handler: handler}
}

func (w *ButtonWidget) Layout(g *gocui.Gui) error {
	if v, err := w.G.SetView(w.Name, w.x, w.y, w.x+w.w, w.y+2, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		w.G.Cursor = false
		if err := w.G.SetKeybinding(w.Name, gocui.KeyEnter, gocui.ModNone, w.handler); err != nil {
			utils.Error.Println(err)
			return err
		}
		fmt.Fprint(v, w.label)
	}
	return nil
}

func (w *ButtonWidget) UnbindShortCuts() {
	w.G.DeleteKeybindings(w.Name)
}
