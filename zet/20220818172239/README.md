# Files


* `locate` finds files by name, if we want to search directory we use `find`, with `find` we can specify which file type we want to find.

* Using find we can use Tests (options) to make finding more clear, and operators to make it even more precise. `-and, -or, -not, ( )`. -and can be shortened to -a, -or to -o, -not to !, parentheses are used to grouptests / operators together to form larger expressions. e.g. `find ~ \( -type f -not -perm 0600 \) -or \( -type d -not -perm 0700 \)`

	* with `find` we can make predefined actions like `-print, -delete, -ls, -quit` - `-print` is default

	* always test by `-print` action before `-delete` to confirm search results

	* we can also use predefined actions using `-exec command {} ;` *command* is a name of a command, *{}* is a symbolic representation of current pathname, *;* is delimiter indicating end of the command

	* Using method above, we are executing command for each found file,  in order to improve efficiency, we can run it only once, we need to change delimiter from *;* to *+* (without single quotes around) ( we cna also use `xargs` combined with pipeline to achieve this.






# Commands

* `locate` - find files by name

* `find` - search for files in a directory hierarchy

* `xargs` - build and execute comand lines from standard input

* `touch` - change file times 

* `stat`  - display file or file system status
