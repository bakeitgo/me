# VIM

| Action | Mode | Description |
|--------|------|-------------|
| i | Command | Change into -- *INSERT* -- mode |
| ESC | INSERT | Escape -- *INSERT* -- to Command. |
| &larr; | Both | Move left |
| &darr; | Both | Move down |
| &uarr; | Both | Move up |
| &rarr; | Both | Move right |
| :wq | Command | Save and quit | 
| :q! | Command | Quit without saving | 
| :q | Command | Quit (fails if anything changed |
| :w | Command | Save without exiting |
| w / b | Command | Next / previous word |
| e / ge | Command | Next / previous end of word |
| 0 / $ | Command | Start / end of line |
| d / dd | Command | delete / delete line |
| c / cc | Command | delete / delete line then start insert mode |
| > / < | Command | indent / unindent one level |
| v / V | Command | Start visual mode |
| ctrl + v | Command | Start visual block mode | 
| Esc | Command | Exit visual / insert mode |
| /pattern | Command | Search for pattern
| ?pattern | Command | Search backward for pattern |
| n / N | Command | Repeat search in same / opposite direction |
| u / ctrl + r | Command | Undo / Redo |
| ctrl + d / ctrl + u | Command | Move down / up half a page |
| { / } | Command | Go forward / backward by paragraph |
| gg / G | Command | Go to the top / bottom of page |
| : [num] [enter] | Command | Go to that line in the document |
| ctrl + e / ctrl + y | Command | Scroll down/up one line |
| m{a-z} | Command | Set mark {a-z} at cursor position |
| '{a-z} | Command | Move the cursor to the start of the line where mark was set |
| '' | Command | Go back to the previous jump location |
| di | Command | deletes everything between parenthesis | 
| . | Command | Repeat last command |
| gv | Command | Select last selected block of text |
| Ctrl + r + 0 | Command | In insert mode inserts the last yanked text (or in command mode) |




# Relates to: 
* http://www.rwx.gg/tools/editors/vi/how/survive
* https://vimsheet.com/

#VIM #Shortcuts
