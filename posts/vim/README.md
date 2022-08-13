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
| h j k l | Commands | move left / down / up / right |
| x | Command | delete character under cursor |
| dw | Command | delete until the start of next word (excluding its first character) |
| d$ | Command | delete to end of the line |
| de | Command | delete to the end of the current word (including the last character) |
| d2w | Command | delete 2 words |
| dd | Command | delete line |
| 2dd | Command | delete 2 lines |
| 2w | Command | move 2 words forward |
| 2e | Command | move 2 word forward to the end of the second word |
| 0 | Command | move to the start of line |
| $ | Commnad | move to the end of line |
| p | Command | put previously deleted text after the cursor |
| ce / cw | Commands | delete word place cursor on end of the deleteed word, turn on insert mode |
| cc | Command | delete line and turn on insert mode |
| nG | Command | move to the n line |
| Ctrl+O | Command | go back where you come from (search) |
| Ctrl+I | Command | go forward in search (reverse ctrl+O) |
| % | Command | place cursor to the nearest `(, ), {, }, ], [` |
| :s/old/new/g | Command | replace all `old` word with `new` word in the line where cursor stands |
| :s/old/new | Command | replace first `old` word with `new` word in the line where cursor stands |
| :%s/old/new/g | Command | replace all `old` word with `new` word in whole file |
| :%s/old/new/gc | Command | replace all `old` word with `new` word in whole file but with prompting which to change |
| :#,#s/old/new/g | Command | replace all `old` word with `new` word between #,# lines |
| :! | Command | Execute external shell command |
| v | Command | Turns on visual mode ( by using j k h l) we can select text in this mode, we can save this part by using :w FILENAME
| :r filename | Command | insert content from filename after cursor |
| o | Command | open new line and turn on insert mode |
| a | Command | append text after the cursor |
| A | Command | append text to the end of a line where cursor stands |
| R | Command | replace more than one character (replacing mode) |
| y | Command | copy line |
| p | Command | paste copied text |
| :set ic | Command | sets the ignore case option, (when search for some word, it ignores upper/lowercasing) | 
| set hls is | Command | sets hlsearch and incsearch options (now when search, word which matches is highlighted) |
| set nohls nois | Command | disable previously turned on option |
| /word\c | Command | turn on ignore case in search but only once (per search) | 


*** 
# Relates to: 
* http://www.rwx.gg/tools/editors/vi/how/survive
* https://vimsheet.com/

#VIM #Shortcuts
