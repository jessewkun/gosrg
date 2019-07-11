package ui

import (
	"gosrg/config"
	"gosrg/redis"
	"gosrg/utils"
	"strconv"

	"github.com/awesome-gocui/gocui"
)

func setNextView() {
	config.Srg.TabNo++
	next := config.Srg.TabNo % len(config.TabView)
	config.Srg.NextView = config.Srg.AllView[config.TabView[next]]
}

func getCurrentLine(v *gocui.View) string {
	var line string
	var err error

	_, cy := v.Cursor()
	if line, err = v.Line(cy); err != nil {
		utils.Logger.Println(err)
		return ""
	}
	return line
}

func up(v *gocui.View) error {
	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := v.SetOrigin(ox, oy-1); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func down(v *gocui.View) error {
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy+1); err != nil {
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func ServerInitHandler() error {
	utils.Soutput("Current Host: " + config.Srg.Host)
	utils.Soutput("Current Port: " + config.Srg.Port)
	utils.Soutput("Current Db  : " + strconv.Itoa(config.Srg.Db))
	return nil
}

func ServerFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = false
	utils.Clear(config.Srg.AllView["server"].View)
	utils.Soutput("Current Host: " + config.Srg.Host)
	utils.Soutput("Current Port: " + config.Srg.Port)
	utils.Soutput("Current Db  : " + strconv.Itoa(config.Srg.Db))
	utils.Toutput(config.TipsMap["server"])
	redis.Info()
	return nil
}

func ServerBlurHandler() error {
	return nil
}

func KeyInitHandler() error {
	redis.Keys()
	return nil
}

func KeyFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = true
	utils.Toutput(config.TipsMap["key"])
	// 暂时关闭 key view 的 KeyDetail, 因为要看 info 的时候必须经过 key, 如果开启的话就会覆盖掉 detail 了
	// if key := getCurrentLine(config.Srg.AllView["key"].View); key != "" {
	// 	redis.KeyDetail(key)
	// }
	return nil
}

func KeyBlurHandler() error {
	return nil
}

func KeyUpHandler(g *gocui.Gui, v *gocui.View) error {
	if err := up(v); err != nil {
		return err
	}
	if key := getCurrentLine(v); key != "" {
		redis.KeyDetail(key)
	}
	return nil
}

func KeyDownHandler(g *gocui.Gui, v *gocui.View) error {
	if err := down(v); err != nil {
		return err
	}
	if key := getCurrentLine(v); key != "" {
		redis.KeyDetail(key)
	}
	return nil
}

func KeyDetailHandler(g *gocui.Gui, v *gocui.View) error {
	if key := getCurrentLine(v); key != "" {
		redis.KeyDetail(key)
	}

	return nil
}

func DetailInitHandler() error {
	return nil
}
func DetailFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = true
	utils.Toutput(config.TipsMap["detail"])
	return nil
}
func DetailBlurHandler() error {
	return nil
}

func DetailSaveHandler(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		redis.SetKeyDetail(v.ViewBuffer())
	}
	return nil
}

func OutputInitHandler() error {
	return nil
}
func OutputFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = false
	utils.Toutput(config.TipsMap["output"])
	return nil
}
func OutputBlurHandler() error {
	return nil
}

func TipInitHandler() error {
	utils.Toutput(config.TipsMap["tip"])
	return nil
}
func TipFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = false
	return nil
}
func TipBlurHandler() error {
	return nil
}

func ProjectInitHandler() error {
	utils.Poutput(config.PROJECT_NAME + " " + config.PROJECT_VERSION)
	return nil
}
func ProjectFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = false
	return nil
}
func ProjectBlurHandler() error {
	return nil
}

func ProjectOpenHandler(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		utils.OpenLink(config.PROJECT_URL)
	}
	return nil
}
func HelpInitHandler() error {
	return nil
}
func HelpFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = false
	utils.Toutput(config.TipsMap["help"])
	utils.Houtput(config.HELP_CONTENT)
	return nil
}
func HelpBlurHandler() error {
	return nil
}

func HelpHideHandler(g *gocui.Gui, v *gocui.View) error {
	name := config.Srg.AllView["help"].Name
	if err := config.Srg.G.DeleteView(name); err != nil {
		return err
	}
	setCurrent(config.Srg.NextView)
	return nil
}

func DbInitHandler() error {
	return nil
}
func DbFocusHandler(arg ...interface{}) error {
	config.Srg.G.Cursor = true
	utils.Toutput(config.TipsMap["db"])
	for index := 0; index <= config.REDIS_MAX_DB_NUM; index++ {
		utils.DBoutput("> database " + strconv.Itoa(index))
	}
	return nil
}
func DbBlurHandler() error {
	return nil
}

func DbHideHandler(g *gocui.Gui, v *gocui.View) error {
	name := config.Srg.AllView["db"].Name
	if err := config.Srg.G.DeleteView(name); err != nil {
		return err
	}
	redis.Db()
	setCurrent(config.Srg.NextView)
	return nil
}

func DbUpHandler(g *gocui.Gui, v *gocui.View) error {
	return up(v)
}

func DbDownHandler(g *gocui.Gui, v *gocui.View) error {
	return down(v)
}

func DbSelectHandler(g *gocui.Gui, v *gocui.View) error {
	if str := getCurrentLine(v); str != "" {
		// db, _ := strconv.Atoi(str)
		db := 1
		config.Srg.Db = db
		DbHideHandler(g, v)
	}

	return nil
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
