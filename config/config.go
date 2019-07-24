package config

const PROJECT_URL = "https://github.com/jessewkun/gosrg"
const PROJECT_NAME = "Gosrg"
const PROJECT_VERSION = "v0.1"

const DEBUG = false

const HELP_CONTENT = `[Global]
	Ctrl-c : Quit
	h      : Display help modal
	Tab    : Toggle next view
	ctrl-d : Display database modal
	ctrl-n : Display new redis connection modal

[help]
	↑↓  : Move up/down one line
	Esc   : Close current modal

[Database]
	↑↓      : Move up/down one line
	Enter     : Select current database
	MouseLeft : Select current database
	Esc   	: Close current modal

[Create new redis connection]
	Enter : Confirm create new redis connection
	Tab   : Toggle focus
	Esc   : Close current modal

[Keys]
	↑↓      : Move up/down one line
	MouseLeft : Show datail
	Delete    : Delete key
	Ctrl+r    : Refresh keys
	Ctrl+f    : Filter key
	Ctrl+b    : Jump to first key
	Ctrl+e    : Jump to last key
	Ctrl+y    : Copy current key to clipboard

[Delete key]
	Enter : Confirm delete key
	Tab   : Toggle focus
	Esc   : Close current modal

[Filter key]
	Enter : Confirm filter pattern
	Tab   : Toggle focus
	Esc   : Close current modal

[Detail]
	i      : Toggle to insert mode
	Esc    : Toggle to normal mode
	↑↓   : Move up/down one line
	Ctrl-s : Save detail
	Ctrl+b : Jump to the beginning
	Ctrl+e : Jump to the end
	Ctrl+y : Copy detail to clipboard
	Ctrl+p : Paste content
	Ctrl+l : Clear detail

[Info]
	Ctrl+y : Copy the info to clipboard

[Output]
	↑↓    : Move up/down one line
	Ctrl+b  : Jump to the begining
	Ctrl+e  : Jump to the end

`

const REDIS_MAX_DB_NUM = 15

var (
	TabView = []string{"server", "key", "detail", "info", "output"}
	TipsMap = map[string]string{
		"server":    "Ctrl-n: Create new redis connection | Tab: Toggle view | Ctrl-c: Quit | h: Help",
		"key":       "↑↓ MouseLeft: Toggle keys | Ctrl+f: Filter | Ctrl+r: Refresh | Deltet: Delete | h: Help",
		"keydel":    "Enter: Confirm delete | Tab: Toggle focus | Esc: Close current modal",
		"keyfilter": "Enter: Execute keys pattern | Tab: Toggle focus | Esc: Close current modal",
		"detail":    "Ctrl-s: Save | Ctrl+y: Copy | Ctrl+p: Paste | Ctrl+l: Clear | h: Help",
		"output":    "Ctrl-b: Jump to the begining | Ctrl-e: Jump to the end | Tab: Toggle view | Ctrl-c: Quit | h: Help",
		"tip":       "Ctrl-n: Create new redis connection | Tab: Toggle view | Ctrl-c: Quit | h: Help",
		"help":      "Esc: Close current modal",
		"db":        "↑↓ MouseLeft: Toggle database | Enter: Select current database | Esc: Close current modal",
		"info":      "Ctrl-y: Copy | Tab: Toggle view | Ctrl-c: Quit | h: Help",
		"conn":      "Enter: Confirm create new redis connection | Tab: Toggle button | Esc: Close connection modal",
	}
)
