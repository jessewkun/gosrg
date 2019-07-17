package config

const PROJECT_URL = "https://github.com/jessewkun/gosrg"
const PROJECT_NAME = "Gosrg"
const PROJECT_VERSION = "unversioned"

const DEBUG = true

const HELP_CONTENT = `
[Global]
 Tab    : Toggle view
 Ctrl-c : Quit
 Space  : Show help modal
 Esc    : Close modal

[Keys]
 ↑↓      : Toggle keys when cursor focus on key view
 MouseLeft : Toggle keys when cursor focus on key view
 Ctrl+f    : Open key filter modal
 Ctrl+r    : Refrsh keys
 Delete    : Delete key

[Detail]
 Ctrl-s : Save detail when cursor focus on detail view

[Db]
 ↑↓      : Chose database when cursor focus on db modal
 Enter     : Select database
 MouseLeft : Toggle database when cursor focus on db modal

[Delete key]
 Enter : Confirm delete key
 Tab   : Toggle button
`

const LOG_FILE = "./gosrg.log"

const REDIS_MAX_DB_NUM = 15

var (
	TabView = []string{"server", "key", "detail", "info", "output"}
	TipsMap = map[string]string{
		"server":    "Tab: Toggle view | Ctrl-c: Quit | Ctrl-space: Help",
		"key":       "↑↓ MouseLeft: Toggle keys | Ctrl+f: Filter keys | Ctrl+r: Refresh keys | Deltet: Delete key",
		"keydel":    "Enter: Confirm delete the key | Tab: Toggle button | Esc: Close Db modal",
		"keyfilter": "Enter: Execute keys pattern | Tab: Toggle button | Esc: Close Db modal",
		"detail":    "Ctrl-s: Save detail",
		"output":    "Tab: Toggle view | Ctrl-c: Quit | Ctrl-space: Help",
		"tip":       "Tab: Toggle view | Ctrl-c: Quit | Ctrl-space: Help",
		"help":      "Esc: Close Help modal",
		"db":        "↑↓ MouseLeft: Toggle database | Enter: Select current database | Esc: Close Db modal",
		"info":      "Tab: Toggle view | Ctrl-c: Quit | Ctrl-space: Help",
	}
)
