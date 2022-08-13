# UNIX

## *hard link* vs *symbolic link*

* They both are way to access files, anywhere from the hard drive

* But what is link?

    * Relationship between two things.

    * It relates to the *inode* (metadata of a file (ownership, file size, type, access mode etc.)) in the file system
***
* ***Hard link***

    * It's some kind of physical link, you can add another name to the file via hard link concept. Using hard link to the file, its like a pointer, you change something in the hard link, orig file changes also. If you remove a hard link, file  still exists (cause you just removed a pointer) - even if you remove a orig file, but hard link still exists, file will still live in the memory, via hard link (if you remove all hard links then file will be completely erased)

    * How to create hard link? `ln <target-name> <hardlinkname>`

    * Note: You cannot create a hard link between partitions, it works only on 1 partition.
***
* ***Symbolic link***

    * It's pointing to the *path* of a file instead of *inode* as hard link does

    * If you change directory of a file, remove a file symbolic link brokes (file is erasing also) You can imagine that it's like a reference to a file

    * U can use *symbolic links* between partitions

    * How to create it? `ln -s <target-name> <softlinkname>`
***
## VIM stuff

* rewind

* Commands

    * `vimtutor` - vim tutorial

    * ` h j k l` - move left / down / up / right

    * `x` - delete character under cursor

    * `dw` - delete until the start of next word (excluding its first character)

    * `d$` - delete to end of the line

    * `de` - delete to the end of the current word (including the last character)

    * `d2w` - delete 2 words

    * `dd` - delete line

    * `2dd` - delete 2 lines

    * `2w` - move 2 wrods forward

    * `2e` - move 2 words forward to the end of second word

    * `0` - move to the start of line

    * `$` - move to the end of a line

    * `u` - undo

    * `ctrl+r` - redo

    * `p` - put previously deleted text after the cursor

    * `rx` - replace character under cursor with x

    * `ce / cw` - delete word place cursor on end of the deleted word, turn on insert mode

    * `cc` - delete line and turn on insert mode

    * `G` - move to the bottom of the file

    * `gg` - move to the start of the file

    * `nG` - write number of line to jump to e.g. 470G (jumps to 470 line)

    * `/` - search command

    * `n` - serach for same phrase again

    * `N` - search for the same phrase in opposite direction

    * `?` - search for phrase in backward direction

    * `CTRL+O` - go back where you come from (search)

    * `CTRL+I` - go forward

    * `%` - place cursor to the nearest *), }, ]*

    * `:s/old/new/g` - replace all 'old' word for 'new' in the line where cursor stands

    * `:s/old/new` - replace first 'old' word with 'new'

    * `:%s/old/new/g` - change every occurence of 'old' word with 'new' in the whole file 

    * `:%s/old/new/gc` - find every occurence in the whole file with a prompt to substitute or not.

    * `:#,#s/old/new/g` - change 'old' word with 'new' in the #,# line range

    * `:!` - execute external shell command

    * `v` - turns visual mode ( by using j k h l) we can select text in this mode, we can save this part by using :w FILENAME

    * `:r filename` insert content from filename after cursor

    * `o` - open new line below and turn on insert mode

    * `a` append text after the cursor

    * `A` append text starting from end of line

    * `R` - replace more than one character (after turning mode on start writing (isnert does same))

    * `y` - copy line

    * `p` - paste copied text

    * `:set ic` - sets the ignore case option, (when search for some word, it ignores upper/lowercasing)

    * `:set hls is` - sets the hlsearch and incsearch options (now when search words are highlighted)

    * `to disable type no between option` e.g. `:set nohls nois`

    * If you want to ignore case only for one search use `\c` after word to serach e.g. `/ignore\c`















# Related to

* https://youtu.be/PSSg7NWXe4Y?t=6938




#hardlink #softlink #reference #pointer #inode #vim