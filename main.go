package main

import (
	"flag"
	"fmt"
	"os"

	"gosrg/config"
	"gosrg/redis"
	"gosrg/ui"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var (
	help    bool
	version bool
	host    string
	port    string
	pwd     string
	pattern string
	logPath string
)

func initFlag() {
	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&version, "v", false, "show version")
	flag.StringVar(&host, "h", "127.0.0.1", "redis host")
	flag.StringVar(&port, "p", "6379", "redis port")
	flag.StringVar(&pwd, "P", "", "redis password")
	flag.StringVar(&pattern, "f", "*", "default key filter pattern")
	flag.StringVar(&logPath, "l", "/var/log/gosrg.log", "default log path")

	flag.Usage = Usage

	flag.Parse()
	if help {
		Usage()
		os.Exit(0)
	}
	if version {
		fmt.Println(config.PROJECT_NAME + "/" + config.PROJECT_VERSION)
		os.Exit(0)
	}
}

func Usage() {
	fmt.Fprintf(os.Stderr, `%s/%s
Usage:
  gosrg -h -p -P -f

Options:
`, config.PROJECT_NAME, config.PROJECT_VERSION)
	flag.PrintDefaults()
}

func main() {
	initFlag()
	utils.InitLog(logPath)
	redis.InitRedis(host, port, pwd, pattern)

	var err error
	ui.Ui.G, err = gocui.NewGui(gocui.Output256, true)
	if err != nil {
		utils.Exit(err)
	}
	defer ui.Ui.G.Close()

	ui.Ui.G.Highlight = true
	ui.Ui.G.Cursor = true
	ui.Ui.G.Mouse = true
	ui.Ui.G.SelFrameColor = gocui.ColorGreen
	ui.Ui.G.SelFgColor = gocui.ColorGreen

	ui.InitUI()

	if err := ui.Ui.G.MainLoop(); err != nil && err != gocui.ErrQuit {
		utils.Exit(err)
	}
}
