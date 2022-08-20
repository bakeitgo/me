# Shell scripts

* How to execute script in every directory?

	* We need to add script location to path, but better practice is to put it in *$HOME/bin* directory, or if we want to let execute it by other users, store it in */usr/local/bin*

	* If the path doesn't contain directory , move to *.bashrc* file and include it in file: `export PATH=~/bin:"$PATH"


* VIM options for better scripting

	* Via `:syntax on` option configures VIM for syntax highlighting

	* Via `:set tabstop=4` allows long lines to fit more easily on the screen

	* Via `:set autoindent` indents line the same amount as the line just typed




#shell #unix #scripts
