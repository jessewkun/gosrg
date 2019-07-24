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
<kbd>Ctrl+d</kbd>        | Global               | Display database modal
<kbd>Ctrl+n</kbd>        | Global               | Display new redis connection modal
<kbd>Down</kbd>          | Help modal           | Move down one line
<kbd>Up</kbd>            | Help modal           | Move up one line
<kbd>Esc</kbd>           | Help modal           | Close current modal
<kbd>Down</kbd>          | Database modal       | Move down one line
<kbd>Up</kbd>            | Database modal       | Move up one view
<kbd>MouseLeft</kbd>     | Database modal       | Select current database
<kbd>Enter</kbd>         | Database modal       | Select current database
<kbd>Esc</kbd>           | Database modal       | Close current modal
<kbd>Enter</kbd>         | New connection modal | Confirm create new redis connection
<kbd>Tab</kbd>           | New connection modal | Toggle focus
<kbd>Esc</kbd>           | New connection modal | Close current modal
<kbd>Down</kbd>          | Keys                 | Move down one line
<kbd>Up</kbd>            | Keys                 | Move up one line
<kbd>MouseLeft</kbd>     | Keys                 | Show detail
<kbd>Backspace</kbd>     | Keys                 | Delete key
<kbd>Ctrl+r</kbd>        | keys                 | Refresh keys
<kbd>Ctrl+f</kbd>        | keys                 | Filter key
<kbd>Ctrl+b</kbd>        | keys                 | Jump to first key
<kbd>Ctrl+e</kbd>        | keys                 | Jump to last key
<kbd>Ctrl+y</kbd>        | keys                 | Copy current key to clipbpard
<kbd>Enter</kbd>         | Delete key modal     | Confirm delete key
<kbd>Tab</kbd>           | Delete key modal     | Toggle focus
<kbd>Esc</kbd>           | Delete key modal     | Close current modal
<kbd>Enter</kbd>         | Filter key modal     | Confirm filter pattern
<kbd>Tab</kbd>           | Filter key modal     | Toggle focus
<kbd>Esc</kbd>           | Filter key modal     | Close current modal
<kbd>i</kbd>             | Detail               | Toggle to insert mode
<kbd>Esc</kbd>           | Detail               | Toggle to normal mode
<kbd>Down</kbd>          | Detail               | Move down one line
<kbd>Up</kbd>            | Detail               | Move up one line
<kbd>Ctrl+s</kbd>        | Detail               | Save detail
<kbd>Ctrl+b</kbd>        | Detail               | Jump to the begining
<kbd>Ctrl+e</kbd>        | Detail               | Jump to the end
<kbd>Ctrl+y</kbd>        | Detail               | Copy detail to clipbpard
<kbd>Ctrl+p</kbd>        | Detail               | Paste content
<kbd>Ctrl+l</kbd>        | Detail               | Clear detail
<kbd>Ctrl+y</kbd>        | Info                 | Copy Info to clipbpard
<kbd>Down</kbd>          | Output               | Move down one line
<kbd>Up</kbd>            | Output               | Move up one line
<kbd>Ctrl+b</kbd>        | Output               | Jump to the begining
<kbd>Ctrl+e</kbd>        | Output               | Jump to the end



## Bugs

Bugs or suggestions? Visit the [issue tracker](https://github.com/jessewkun/gosrg/issues)
