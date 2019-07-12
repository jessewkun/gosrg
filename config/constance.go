package config

const REDIS_NETWORK = "tcp"

const PROJECT_URL = "https://github.com/jessewkun/gosrg"
const PROJECT_NAME = "Gosrg"
const PROJECT_VERSION = "unversioned"

const DEBUG = true

const HELP_CONTENT = `
[Global]
 tab    : Toggle view
 ctrl-c : Quit
 Space  : Show help view

[Keys]
 ↑↓      : Toggle keys when cursor focus on key view
 MouseLeft : Toggle keys when cursor focus on key view

[Detail]
 ctrl-s : Save detail when cursor focus on detail view

[Db]
 ↑↓      : Chose database when cursor focus on key view
 Enter     : Select current database
 MouseLeft : Toggle database when cursor focus on key view
`

const LOG_FILE = "./gosrg.log"

const REDIS_MAX_DB_NUM = 15
