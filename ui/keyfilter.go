package ui

import (
	"gosrg/redis"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var kfView *KeyFilterView

type KeyFilterView struct {
	GView
}

func init() {
	kfView = new(KeyFilterView)
	kfView.Name = "keyfilter"
	kfView.Title = " key filter "
	kfView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.KeyEsc, Level: GLOBAL_Y, Handler: kfView.hide},
		ShortCut{Key: gocui.KeyTab, Level: GLOBAL_Y, Handler: kfView.tab},
	}
}

func (kf *KeyFilterView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(kf.Name, maxX/3-10, maxY/3-6, maxX/2+40, maxY/2-5, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = kf.Title
		v.Wrap = true
		v.Editable = true
		kf.View = v
		kf.initialize()
	}
	return nil
}

func (kf *KeyFilterView) initialize() error {
	gView.unbindShortCuts()
	kf.setCurrent(kf)
	kf.btn()
	kf.bindShortCuts()
	kf.outputln(redis.R.Pattern)
	return nil
}

func (kf *KeyFilterView) tab(g *gocui.Gui, v *gocui.View) error {
	nextViewName := ""
	currentView := Ui.G.CurrentView().Name()
	if currentView == confirmBtn.Name {
		nextViewName = cancelBtn.Name
	} else {
		nextViewName = confirmBtn.Name
	}
	if _, err := Ui.G.SetCurrentView(nextViewName); err != nil {
		utils.Logger.Fatalln(err)
		return err
	}
	return nil
}

func (kf *KeyFilterView) hide(g *gocui.Gui, v *gocui.View) error {
	kf.unbindShortCuts()
	gView.bindShortCuts()
	if err := Ui.G.DeleteView(confirmBtn.Name); err != nil {
		utils.Logger.Println(err)
		return err
	}
	if err := Ui.G.DeleteView(cancelBtn.Name); err != nil {
		utils.Logger.Println(err)
		return err
	}
	if err := Ui.G.DeleteView(kf.Name); err != nil {
		utils.Logger.Println(err)
		return err
	}
	Ui.NextView.setCurrent(Ui.NextView)
	return nil
}

func (kf *KeyFilterView) btn() error {
	maxX, maxY := Ui.G.Size()
	confirmBtn = NewButtonWidget("confirfilter", maxX/3-5, maxY/3-1, "CONFIRM", func(g *gocui.Gui, v *gocui.View) error {
		pattern := utils.Trim(kf.View.ViewBuffer())
		if len(pattern) == 0 {
			pattern = "*"
		}
		redis.R.Pattern = pattern
		output, keys := redis.R.Keys()
		opView.formatOutput(output)
		kView.clear()
		for _, key := range keys {
			kView.outputln(key)
		}
		kView.View.Title = " Keys " + redis.R.Pattern
		kf.hide(g, v)
		return nil
	})
	cancelBtn = NewButtonWidget("cancelfilter", maxX/3+5, maxY/3-1, "CANCEL", func(g *gocui.Gui, v *gocui.View) error {
		kf.hide(g, v)
		return nil
	})
	confirmBtn.Layout(Ui.G)
	cancelBtn.Layout(Ui.G)

	return nil
}
