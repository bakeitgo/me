# UNIX stuff

### file linking

* ln - create  links (Symbolic / hard)
	* ln - create hard link
	* ln -s - create symbolic link

* hard link contains same inode, using hard link we can manipulate file. deleting hard link deletes only hard linked file, not all files with same inode.

* it is better practice to use soft links

* **hard link cannot reference directories**

* when creating soft links, file to which we want to link, need to have **path led from directory where link will be stored** e.g. if file with name bar is in foo directory and we want to place link in top of foo directory, we need to: `ln -s ./foo/bar ./bar-soft-link`

### Commands

| Command | Description |
|---------|-------------|
| type | indicate how a command name is interpreted |
| which | display exectuable program path (only exe files!) |
| help | get help for shell builtins |
| man | Display a command's manual page |
| apropos | Display a list of appropriate commands (displays what man page we can display about thing) |
| info | Display a command info entry (man alternative) |
| whatis | display one line manual page descriptions |
| alias  | create an alias for a command |

* What tha hell are commands? ( 4 possibilities )

	* an exe file

	* a command built into the shell itself e.g. `cd`

	* a shell function (miniature shell script)
	
	* an alias - alias is command named by ourselves, e.g. we can alias `cd` with `moveto`, and use `moveto` instead of `cd`

* Redirecting stdout to a file

	* `>` overrides file if it exists

	* `>>` appends to a file

* how to redirect stderr?

	* stdin, stdout,stderr are indicated by numbers: 0,1,2 respectively.
	* to redirect stderr, we need to use 2> instead of >

* how to redirect both?

	* e.g. ls -l /bin/usr > ls-output.txt 2>&1

	* above, we firstly redirect stdout, then stderr



#stderr #stdout #stdin #redirection #commands #hardlink #softlink #symboliclink #link #unix #gnu #linux



