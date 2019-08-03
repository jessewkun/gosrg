# Gosrg

Terminal GUI management tool for Redis

![Gosrg](https://raw.githubusercontent.com/jessewkun/gosrg/master/docs/images/Screenshots.png)

## Installation and usage

```
$ go get github.com/jessewkun/gosrg
$ "$GOPATH/bin/gosrg" --help
Terminal GUI management tool for Redis

Version: v0.1.3
Build Time: 2019-08-01 18:25:49
Commit SHA1: ed505c85a92acd30194f8033d71b438a7e645d6a

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
<kbd>Ctrl+t</kbd>        | Global               | Display command modal
<kbd>Ctrl+r</kbd>        | Server               | Refresh redis info
<kbd>Down/Up</kbd>       | Help modal           | Move down/up one line
<kbd>Esc</kbd>           | Help modal           | Close current modal
<kbd>Down/Up</kbd>       | Database modal       | Move down/up one line
<kbd>MouseLeft</kbd>     | Database modal       | Select current database
<kbd>Enter</kbd>         | Database modal       | Select current database
<kbd>Esc</kbd>           | Database modal       | Close current modal
<kbd>Down/Up</kbd>       | New connection modal | Choose historical connection
<kbd>Enter</kbd>         | New connection modal | Confirm create new redis connection
<kbd>Tab</kbd>           | New connection modal | Toggle focus
<kbd>Esc</kbd>           | New connection modal | Close current modal
<kbd>Down/Up</kbd>       | Keys                 | Move down/up one line
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
<kbd>Down/Up</kbd>       | Detail               | Move down/up one line
<kbd>Ctrl+s</kbd>        | Detail               | Save detail
<kbd>Ctrl+b</kbd>        | Detail               | Jump to the begining
<kbd>Ctrl+e</kbd>        | Detail               | Jump to the end
<kbd>Ctrl+y</kbd>        | Detail               | Copy detail to clipbpard
<kbd>Ctrl+p</kbd>        | Detail               | Paste content
<kbd>Ctrl+l</kbd>        | Detail               | Clear detail
<kbd>Ctrl+y</kbd>        | Info                 | Copy Info to clipbpard
<kbd>Down/Up</kbd>       | Output               | Move down/up one line
<kbd>Ctrl+b</kbd>        | Output               | Jump to the begining
<kbd>Ctrl+e</kbd>        | Output               | Jump to the end



## Bugs

Bugs or suggestions? Visit the [issue tracker](https://github.com/jessewkun/gosrg/issues)
