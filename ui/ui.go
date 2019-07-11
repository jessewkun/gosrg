package ui

import (
	"fmt"
	"gosrg/config"
	"gosrg/utils"

	"github.com/awesome-gocui/gocui"
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView(TipView.Name, 0, maxY-2, maxX-20, maxY, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Frame = false
		TipView.View = v
		TipView.InitHandler()
	}

	if v, err := g.SetView(ProjectView.Name, maxX-19, maxY-2, maxX-1, maxY, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Frame = false
		ProjectView.View = v
		ProjectView.InitHandler()
	}

	if v, err := g.SetView(OutputView.Name, maxX/3+1, maxY-14, maxX-1, maxY-2, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = OutputView.Title
		v.Wrap = true
		v.Autoscroll = true
		OutputView.View = v
		OutputView.InitHandler()
	}

	if v, err := g.SetView(DetailView.Name, maxX/3+1, 0, maxX-1, maxY-15, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = DetailView.Title
		v.Wrap = true
		v.Editable = true
		DetailView.View = v
		DetailView.InitHandler()
	}

	if v, err := g.SetView(ServerView.Name, 0, 0, maxX/3, maxY/10, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = ServerView.Title
		v.Wrap = true
		ServerView.View = v
		ServerView.InitHandler()
		setCurrent(&ServerView)
	}

	if v, err := g.SetView(KeyView.Name, 0, maxY/10+1, maxX/3, maxY-2, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Title = KeyView.Title
		v.Wrap = true
		v.Autoscroll = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		KeyView.View = v
		KeyView.InitHandler()
	}

	return nil
}

func setCurrent(v *config.View, arg ...interface{}) error {
	if _, err := config.Srg.G.SetCurrentView(v.Name); err != nil {
		utils.Logger.Fatalln(err)
		return err
	}
	str := fmt.Sprintf("current view: %s", config.Srg.NextView.Name)
	utils.Debug(str)
	v.FocusHandler(arg)
	return nil
}

func InitShortCuts() {
	for _, sc := range GlobalShortCuts {
		if err := config.Srg.G.SetKeybinding("", sc.Key, sc.Mod, sc.Handler); err != nil {
			utils.Logger.Fatalln(err)
		}
	}

	for _, v := range config.Srg.AllView {
		for _, sc := range v.ShortCuts {
			if err := config.Srg.G.SetKeybinding(v.Name, sc.Key, sc.Mod, sc.Handler); err != nil {
				utils.Logger.Fatalln(err)
			}
		}
	}
}
