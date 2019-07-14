package main

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/ui"
	"gosrg/utils"
	"runtime"

	"github.com/awesome-gocui/gocui"
)

func main() {
	utils.InitLog()

	if runtime.GOOS == "windows" {
		utils.Logger.Fatalln(config.PROJECT_NAME + " is not support Windows")
	}
	config.InitSrg()
	redis.InitRedis()

	var err error
	config.Srg.G, err = gocui.NewGui(gocui.Output256, true)
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	defer config.Srg.G.Close()

	config.Srg.G.Highlight = true
	config.Srg.G.Cursor = true
	config.Srg.G.Mouse = true
	config.Srg.G.SelFrameColor = gocui.ColorGreen
	config.Srg.G.SelFgColor = gocui.ColorGreen

	ui.InitConfigAllView()
	config.Srg.G.SetManagerFunc(ui.Layout)
	ui.InitShortCuts()

	if err := config.Srg.G.MainLoop(); err != nil && err != gocui.ErrQuit {
		utils.Logger.Fatalln(err)
	}

}
