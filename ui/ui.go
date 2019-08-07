package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/ui/base"
	"gosrg/utils"

	"github.com/jessewkun/gocui"
)

var Ui UI

type UI struct {
	G        *gocui.Gui
	AllView  map[string]base.GInterfacer
	TabNo    int
	NextView base.GInterfacer
}

func RestNextView() {
	Ui.TabNo = 0
	Ui.NextView = sView
}

func setNextView() {
	Ui.TabNo++
	next := Ui.TabNo % len(config.TabView)
	Ui.NextView = Ui.AllView[config.TabView[next]]
}

func InitUI() {
	Ui.AllView = map[string]base.GInterfacer{
		"global":  gView,
		"server":  sView,
		"info":    iView,
		"key":     kView,
		"detail":  dView,
		"output":  opView,
		"tip":     tView,
		"project": pView,
	}
	Ui.NextView = sView
	Ui.G.SetManager(iView, tView, pView, opView, dView, sView, kView)
	for _, item := range Ui.AllView {
		item.SetG(Ui.G)
		item.BindShortCuts()
	}
}

func Render() {
	for {
		select {
		case res := <-redis.R.ResultChan:
			for rtype, item := range res {
				switch rtype {
				case redis.RES_OUTPUT_COMMAND:
					fallthrough
				case redis.RES_OUTPUT_INFO:
					fallthrough
				case redis.RES_OUTPUT_ERROR:
					fallthrough
				case redis.RES_OUTPUT_RES:
					Ui.G.Update(func(*gocui.Gui) error {
						opView.formatOutput(rtype, item)
						return nil
					})
				case redis.RES_KEYS:
					Ui.G.Update(func(*gocui.Gui) error {
						kView.formatOutput(item)
						return nil
					})
				case redis.RES_DETAIL:
					Ui.G.Update(func(*gocui.Gui) error {
						dView.formatOutput(item)
						return nil
					})
				case redis.RES_INFO:
					Ui.G.Update(func(*gocui.Gui) error {
						iView.formatOutput(item)
						return nil
					})
				case redis.RES_EXIT:
					return
				default:
					utils.Error.Println(res)
				}
			}
			redis.Locker.Unlock()
		}
	}
}
