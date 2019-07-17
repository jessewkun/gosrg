package ui

import (
	"gosrg/config"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var pView *ProjectView

type ProjectView struct {
	GView
}

func init() {
	pView = new(ProjectView)
	pView.Name = "project"
	pView.ShortCuts = []ShortCut{
		ShortCut{Key: gocui.MouseLeft, Level: LOCAL_N, Handler: pView.openGit},
	}
}

func (p *ProjectView) Layout(g *gocui.Gui) error {
	maxX, maxY := Ui.G.Size()
	if v, err := g.SetView(p.Name, maxX-19, maxY-2, maxX-1, maxY, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			utils.Logger.Fatalln(err)
			return err
		}
		v.Frame = false
		p.View = v
		p.initialize()
	}
	return nil
}

func (p *ProjectView) initialize() error {
	p.output(utils.Pink(utils.UnderLine(config.PROJECT_NAME + " " + config.PROJECT_VERSION)))
	return nil
}

func (p *ProjectView) openGit(g *gocui.Gui, v *gocui.View) error {
	utils.OpenLink(config.PROJECT_URL)
	return nil
}
