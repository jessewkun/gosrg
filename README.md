# Gosrg

Terminal GUI management tool for Redis



## Installation and usage

```
$ go get github.com/jessewkun/gosrg
$ "$GOPATH/bin/gosrg" --help
```
[Binary releases](https://github.com/jessewkun/gosrg/releases) are also available.


### Commands

Keybinding               | Scope                | Description
-------------------------|----------------------|--------------------------------------------------
<kbd>Ctrl+C</kbd>        | Global               | Quit
<kbd>h</kbd>             | Global               | Display help modal
<kbd>Tab</kbd>           | Global               | Toggle next view
<kbd>Ctrl+d</kbd>        | Global               | Display db modal
<kbd>Down</kbd>          | Keys                 | Move down one view line and show detail
<kbd>Up</kbd>            | Keys                 | Move up one view line and show detail
<kbd>MouseLeft</kbd>     | Keys                 | Show detail
<kbd>Backspace</kbd>     | Keys                 | Show delete key modal
<kbd>Ctrl+r</kbd>        | keys                 | Refresh keys
<kbd>Ctrl+f</kbd>        | keys                 | Display key filter modal
<kbd>Ctrl+b</kbd>        | keys                 | Jump to first key
<kbd>Ctrl+e</kbd>        | keys                 | Jump to last key
<kbd>i</kbd>             | Detail               | Toggle to insert mode
<kbd>Esc</kbd>           | Detail               | Toggle to normal mode
<kbd>Ctrl+y</kbd>        | Detail               | Copy detail to clipbpard
<kbd>Ctrl+p</kbd>        | Detail               | Paste detail to detail view
<kbd>Ctrl+l</kbd>        | Detail               | Clear detail
<kbd>Ctrl+s</kbd>        | Detail               | Save detail
<kbd>Ctrl+b</kbd>        | Detail               | Jump to detail begining
<kbd>Ctrl+e</kbd>        | Detail               | Jump to end of detail
<kbd>Esc</kbd>           | Help modal           | Close help modal
<kbd>Down</kbd>          | Help modal           | Move down one view line
<kbd>Up</kbd>            | Help modal           | Move up one view line
<kbd>Enter</kbd>         | Delete key modal     | Confirm delete key when focus on CONFIRM button
<kbd>Esc</kbd>           | Delete key modal     | Close key delete modal
<kbd>Tab</kbd>           | Delete key modal     | Toggle focus button
<kbd>Esc</kbd>           | Filter key modal     | Close key filter modal
<kbd>Enter</kbd>         | Filter key modal     | Confirm filter when focus on CONFIRM button
<kbd>Tab</kbd>           | Delete key modal     | Toggle focus
<kbd>Down</kbd>          | Db modal             | Move down one view line
<kbd>Up</kbd>            | Db modal             | Move up one view
<kbd>MouseLeft</kbd>     | Db modal             | Select current db
<kbd>Enter</kbd>         | Db modal             | Select current db
<kbd>Esc</kbd>           | Db modal             | Close db modal


## Bugs

Bugs or suggestions? Visit the [issue tracker](https://github.com/jessewkun/gosrg/issues)
