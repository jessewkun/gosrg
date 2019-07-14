package ui

import (
	"gosrg/config"
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

var GlobalShortCuts = []config.ShortCut{
	config.ShortCut{Key: gocui.KeyCtrlC, Mod: gocui.ModNone, Handler: GlobalQuitHandler},
	config.ShortCut{Key: gocui.KeyTab, Mod: gocui.ModNone, Handler: GlobalTabHandler},
	config.ShortCut{Key: gocui.KeySpace, Mod: gocui.ModNone, Handler: GlobalShowHelpViewHandler},
	config.ShortCut{Key: gocui.KeyCtrlD, Mod: gocui.ModNone, Handler: GlobalShowDbViewHandler},
}

func GlobalQuitHandler(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func GlobalTabHandler(g *gocui.Gui, v *gocui.View) error {
	setNextView()
	if err := setCurrent(config.Srg.NextView); err != nil {
		utils.Logger.Fatalln(err)
		return err
	}
	return nil
}

func GlobalShowHelpViewHandler(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := config.Srg.G.Size()
	name := config.Srg.AllView["help"].Name
	if v, err := config.Srg.G.SetView(name, maxX/3-10, maxY/3-6, maxX/2+40, maxY/2+6, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = HelpView.Title
		v.Wrap = true
		config.Srg.AllView["help"].View = v
		setCurrent(config.Srg.AllView["help"])
	}
	return nil
}

func GlobalShowDbViewHandler(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := config.Srg.G.Size()
	name := config.Srg.AllView["db"].Name
	if v, err := config.Srg.G.SetView(name, maxX/3-10, maxY/3-6, maxX/2+40, maxY/2+6, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = DbView.Title
		v.Wrap = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		config.Srg.AllView["db"].View = v
		setCurrent(config.Srg.AllView["db"])
	}
	return nil
}
