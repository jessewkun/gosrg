package ui

import (
	"gosrg/config"

	"github.com/awesome-gocui/gocui"
)

func InitConfigAllView() {
	config.Srg.AllView = map[string]*config.View{
		"server":  &ServerView,
		"key":     &KeyView,
		"detail":  &DetailView,
		"output":  &OutputView,
		"tip":     &TipView,
		"project": &ProjectView,
		"help":    &HelpView,
	}
	config.Srg.NextView = &ServerView
}

var GlobalShortCuts = []config.ShortCut{
	config.ShortCut{Key: gocui.KeyCtrlC, Mod: gocui.ModNone, Handler: GlobalQuitHandler},
	config.ShortCut{Key: gocui.KeyTab, Mod: gocui.ModNone, Handler: GlobalTabHandler},
	config.ShortCut{Key: gocui.KeyCtrlSpace, Mod: gocui.ModNone, Handler: GlobalShowHelpViewHandler},
}

var ServerView = config.View{
	Name:         "server",
	Title:        " Server Info ",
	InitHandler:  ServerInitHandler,
	FocusHandler: ServerFocusHandler,
	BlurHandler:  ServerBlurHandler,
}

var KeyView = config.View{
	Name:         "key",
	Title:        " Keys ",
	InitHandler:  KeyInitHandler,
	FocusHandler: KeyFocusHandler,
	BlurHandler:  KeyBlurHandler,
	ShortCuts: []config.ShortCut{
		config.ShortCut{Key: gocui.KeyArrowUp, Mod: gocui.ModNone, Handler: KeyUpHandler},
		config.ShortCut{Key: gocui.KeyArrowDown, Mod: gocui.ModNone, Handler: KeyDownHandler},
		config.ShortCut{Key: gocui.MouseLeft, Mod: gocui.ModNone, Handler: KeyDetailHandler},
	},
}

var DetailView = config.View{
	Name:         "detail",
	Title:        " Detail ",
	InitHandler:  DetailInitHandler,
	FocusHandler: DetailFocusHandler,
	BlurHandler:  DetailBlurHandler,
	ShortCuts: []config.ShortCut{
		config.ShortCut{Key: gocui.KeyCtrlS, Mod: gocui.ModNone, Handler: DetailSaveHandler},
	},
}

var OutputView = config.View{
	Name:         "output",
	Title:        " Output ",
	InitHandler:  OutputInitHandler,
	FocusHandler: OutputFocusHandler,
	BlurHandler:  OutputBlurHandler,
}

var TipView = config.View{
	Name:         "tip",
	InitHandler:  TipInitHandler,
	FocusHandler: TipFocusHandler,
	BlurHandler:  TipBlurHandler,
}

var ProjectView = config.View{
	Name:         "project",
	InitHandler:  ProjectInitHandler,
	FocusHandler: ProjectFocusHandler,
	BlurHandler:  ProjectBlurHandler,
	ShortCuts: []config.ShortCut{
		config.ShortCut{Key: gocui.MouseLeft, Mod: gocui.ModNone, Handler: ProjectOpenHandler},
	},
}

var HelpView = config.View{
	Name:         "help",
	Title:        " Help ",
	InitHandler:  HelpInitHandler,
	FocusHandler: HelpFocusHandler,
	BlurHandler:  HelpBlurHandler,
	ShortCuts: []config.ShortCut{
		config.ShortCut{Key: gocui.KeyEsc, Mod: gocui.ModNone, Handler: HelpHideHandler},
	},
}
