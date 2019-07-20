package config

const PROJECT_URL = "https://github.com/jessewkun/gosrg"
const PROJECT_NAME = "Gosrg"
const PROJECT_VERSION = "unversioned"

const DEBUG = false

const HELP_CONTENT = `
[Global]
 Tab    : Toggle view
 Ctrl-c : Quit
 h      : Show help modal
 Esc    : Close modal

[Keys]
 ↑↓      : Toggle keys when cursor focus on key view
 MouseLeft : Toggle keys when cursor focus on key view
 Ctrl+f    : Open key filter modal
 Ctrl+r    : Refrsh keys
 Ctrl+b    : Jump to the first key
 Ctrl+e    : Jump to the last key
 Delete    : Delete key

[Detail]
 i      : Insert mode
 Esc    : Normal mode
 Ctrl-s : Save detail when cursor focus on detail view
 Ctrl+b : Jump to the beginning of the detail
 Ctrl+e : Jump to the tail of the detail
 Ctrl+y : Copy the detail to clipboard
 Ctrl+p : Paste the content from clipboard to detail view
 Ctrl+l : Clear the detail view

[Db]
 ↑↓      : Chose database when cursor focus on db modal
 Enter     : Select database
 MouseLeft : Toggle database when cursor focus on db modal

[Delete key]
 Enter : Confirm delete key
 Tab   : Toggle button
`

const REDIS_MAX_DB_NUM = 15

var (
	TabView = []string{"server", "key", "detail", "info", "output"}
	TipsMap = map[string]string{
		"server":    "Tab: Toggle view | Ctrl-c: Quit | h: Help",
		"key":       "↑↓ MouseLeft: Toggle keys | Ctrl+f: Filter | Ctrl+r: Refresh | Deltet: Delete | h: Help",
		"keydel":    "Enter: Confirm delete the key | Tab: Toggle button | Esc: Close Db modal",
		"keyfilter": "Enter: Execute keys pattern | Tab: Toggle button | Esc: Close Db modal",
		"detail":    "Ctrl-s: Save | Ctrl+y: Copy | Ctrl+p: Paste | Ctrl+l: Clear | h: Help",
		"output":    "Tab: Toggle view | Ctrl-c: Quit | h: Help",
		"tip":       "Tab: Toggle view | Ctrl-c: Quit | h: Help",
		"help":      "Esc: Close Help modal",
		"db":        "↑↓ MouseLeft: Toggle database | Enter: Select current database | Esc: Close Db modal",
		"info":      "Tab: Toggle view | Ctrl-c: Quit | h: Help",
	}
)
