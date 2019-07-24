# Gosrg

Terminal GUI management tool for Redis

![Gosrg](https://raw.githubusercontent.com/jessewkun/gosrg/master/docs/images/Screenshots.png)

## Installation and usage

```
$ go get github.com/jessewkun/gosrg
$ "$GOPATH/bin/gosrg" --help
Usage:
  gosrg -h -p -P -f

Options:
  -help show help
  -P	redis password
  -f	default key filter pattern (default "*")
  -h	redis host (default "127.0.0.1")
  -l	default log path (default "/var/log/gosrg.log")
  -p	redis port (default "6379")
  -v	show version
```
> Warning:
- The option `-f` used to specify default pattern for command `keys`, if this option is not specified, Gosrg will execute `keys *`
- Gosrg will execute `hgetall` when key type is hash
>>
[Binary releases](https://github.com/jessewkun/gosrg/releases) are also available.


## Example
```
# Connect redis with default configuration
#   host: 127.0.0.1
#   port: 6379
#   password: ""
#   pattrn: "*"
$ gosrg

# Connect redis with custom configuration
$ gosrg -h=192.168.33.10 -p=6380 -P=123456 -f=abc
```

### Shortcuts

Keybinding               | Scope                | Description
-------------------------|----------------------|--------------------------------------------------
<kbd>Ctrl+C</kbd>        | Global               | Quit
<kbd>h</kbd>             | Global               | Display help modal
<kbd>Tab</kbd>           | Global               | Toggle next view
<kbd>Ctrl+d</kbd>        | Global               | Display db modal
<kbd>Enter</kbd>         | Server               | Confirm new redis connection
<kbd>Esc</kbd>           | Server               | Close connection modal
<kbd>Tab</kbd>           | Server               | Toggle focus
<kbd>Down</kbd>          | Keys                 | Move down one view line and show detail
<kbd>Up</kbd>            | Keys                 | Move up one view line and show detail
<kbd>MouseLeft</kbd>     | Keys                 | Show detail
<kbd>Backspace</kbd>     | Keys                 | Show delete key modal
<kbd>Ctrl+r</kbd>        | keys                 | Refresh keys
<kbd>Ctrl+f</kbd>        | keys                 | Display key filter modal
<kbd>Ctrl+b</kbd>        | keys                 | Jump to first key
<kbd>Ctrl+e</kbd>        | keys                 | Jump to last key
<kbd>Ctrl+y</kbd>        | keys                 | Copy current key to clipbpard
<kbd>i</kbd>             | Detail               | Toggle to insert mode
<kbd>Esc</kbd>           | Detail               | Toggle to normal mode
<kbd>Down</kbd>          | Detail               | Move down one line
<kbd>Up</kbd>            | Detail               | Move up one line
<kbd>Ctrl+b</kbd>        | Detail               | Jump to detail begining
<kbd>Ctrl+e</kbd>        | Detail               | Jump to end of detail
<kbd>Ctrl+y</kbd>        | Detail               | Copy detail to clipbpard
<kbd>Ctrl+p</kbd>        | Detail               | Paste detail to detail view
<kbd>Ctrl+l</kbd>        | Detail               | Clear detail
<kbd>Ctrl+s</kbd>        | Detail               | Save detail
<kbd>Ctrl+y</kbd>        | Info                 | Copy Info to clipbpard
<kbd>Down</kbd>          | Output               | Move down one line
<kbd>Up</kbd>            | Output               | Move up one line
<kbd>Ctrl+b</kbd>        | Output               | Jump to output begining
<kbd>Ctrl+e</kbd>        | Output               | Jump to end of output
<kbd>Esc</kbd>           | Help modal           | Close help modal
<kbd>Down</kbd>          | Help modal           | Move down one view line
<kbd>Up</kbd>            | Help modal           | Move up one view line
<kbd>Enter</kbd>         | Delete key modal     | Confirm delete key
<kbd>Esc</kbd>           | Delete key modal     | Close key delete modal
<kbd>Tab</kbd>           | Delete key modal     | Toggle focus button
<kbd>Enter</kbd>         | Filter key modal     | Confirm filter pattern
<kbd>Esc</kbd>           | Filter key modal     | Close key filter modal
<kbd>Tab</kbd>           | Filter key modal     | Toggle focus
<kbd>Down</kbd>          | Db modal             | Move down one view line
<kbd>Up</kbd>            | Db modal             | Move up one view
<kbd>MouseLeft</kbd>     | Db modal             | Select current db
<kbd>Enter</kbd>         | Db modal             | Select current db
<kbd>Esc</kbd>           | Db modal             | Close db modal


## Bugs

Bugs or suggestions? Visit the [issue tracker](https://github.com/jessewkun/gosrg/issues)
