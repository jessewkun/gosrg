package main

import (
	"gosrg/redis"
	"gosrg/ui"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

func main() {
	utils.InitLog()
	redis.InitRedis()

	var err error
	ui.Ui.G, err = gocui.NewGui(gocui.Output256, true)
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	defer ui.Ui.G.Close()

	ui.Ui.G.Highlight = true
	ui.Ui.G.Cursor = true
	ui.Ui.G.Mouse = true
	ui.Ui.G.SelFrameColor = gocui.ColorGreen
	ui.Ui.G.SelFgColor = gocui.ColorGreen

	ui.InitUI()

	if err := ui.Ui.G.MainLoop(); err != nil && err != gocui.ErrQuit {
		utils.Logger.Fatalln(err)
	}

}
