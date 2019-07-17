package ui

import (
	"gosrg/redis"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var kdView *KeyDelView
var confirmBtn *ButtonWidget
var cancelBtn *ButtonWidget

type KeyDelView struct {
	GView
}

func init() {
	kdView = new(KeyDelView)
	kdView.Name = "keydel"
	kdView.Title = " WARNING "
	kdView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyEsc, Level: LOCAL_Y, Handler: kdView.hide},
		ShortCut{Key: gocui.KeyTab, Level: GLOBAL_Y, Handler: kdView.tab},
	}
}

func (kd *KeyDelView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(kd.Name, maxX/3-10, maxY/3-6, maxX/2+40, maxY/2-5, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = kd.Title
		v.Wrap = true
		kd.View = v
		kd.initialize()
	}
	return nil
}

func (kd *KeyDelView) initialize() error {
	kd.btn()
	gView.unbindShortCuts()
	kd.bindShortCuts()
	kd.setCurrent(kd)
	kd.outputln("")
	kd.outputln(utils.Red("     Confirm delete this key?"))
	return nil
}

func (kd *KeyDelView) tab(g *gocui.Gui, v *gocui.View) error {
	nextViewName := ""
	currentView := Ui.G.CurrentView().Name()
	if currentView == kd.Name {
		nextViewName = confirmBtn.Name
	} else if currentView == confirmBtn.Name {
		nextViewName = cancelBtn.Name
	} else {
		nextViewName = kd.Name
	}
	if _, err := Ui.G.SetCurrentView(nextViewName); err != nil {
		utils.Logger.Fatalln(err)
		return err
	}
	return nil
}

func (kd *KeyDelView) hide(g *gocui.Gui, v *gocui.View) error {
	kd.unbindShortCuts()
	gView.bindShortCuts()
	if err := Ui.G.DeleteView(confirmBtn.Name); err != nil {
		utils.Logger.Println(err)
		return err
	}
	// for _, item := range Ui.G.Views() {
	// 	utils.Debug(item.Name())
	// }
	if err := Ui.G.DeleteView(cancelBtn.Name); err != nil {
		utils.Logger.Println(err)
		return err
	}
	if err := Ui.G.DeleteView(kd.Name); err != nil {
		utils.Logger.Println(err)
		return err
	}
	Ui.NextView.setCurrent(Ui.NextView)
	return nil
}

func (kd *KeyDelView) btn() error {
	maxX, maxY := Ui.G.Size()
	confirmBtn = NewButtonWidget("confirmdel", maxX/3-5, maxY/3-1, "CONFIRM", func(g *gocui.Gui, v *gocui.View) error {
		if output := redis.R.Del(); len(output) > 0 {
			opView.formatOutput(output)
			kView.deleteCursorLine()
		}
		kd.hide(g, v)
		return nil
	})
	cancelBtn = NewButtonWidget("canceldel", maxX/3+5, maxY/3-1, "CANCEL", func(g *gocui.Gui, v *gocui.View) error {
		kd.hide(g, v)
		return nil
	})
	Ui.G.AppendManager(confirmBtn, cancelBtn)

	return nil
}
