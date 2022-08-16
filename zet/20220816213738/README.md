# UNIX permissions / bash


### Bash 

* always double quote expansions / variables  to prevent word splitting


### Permissions

* ugoa - user group others all

	* user contains uid (userid)

	* group contains gid (groupid)

* User accounts arer defined in the */etc/passwd* file

* Groups are defined in */etc/group* file

* *Note!* Besides regular user accounts, there are accounts for **superuser** *(uid 0)*

* Access rights to files are defined in *rwx (read write execute)*

	* How to check them? via `ls -l <file>` command

	* It lists *file type, permissions, hardlinks, useritbelongs, groupitbelongs, size, date, filename*

	* Permissions are listed: `-rw-rw-r--` (first hyphen is a file type) and rest are permissions in order: (Owner, Group, World)
	
	* Possible file types: `-` regular file, `d` - directory, `l` symbolic link, `c` - character special file, `b` - block special file

	* Permission attributes:

	| Attribute | Files | Directories |
	|-----------|-------|-------------|
	| r | Allows a file to be opened and read | If the `x` attribute is also set, it allows to list directory contents |
	| w | Allows a file to be written to or truncated | Allows filesinside a directory to be *created* *deleted* *rernamed* if the `x` attrirbute is also set |
	| x | Allows file to be executed (programs need have readable permission also | Allows directory to be entered |
	
	* To change permissions of a file use `chmod` command (only file owner can do it) (it can be used using symbolic / octal representation)


	* `umask` - set default permission





### Commands

* `id` - display user identity

* `chmod` - change a file's mode

* `umask` - set the default file permissions

* `su` - run a shell as another user

* `sudo` - execute command as another user

* `chown` - change a file's owner

* `chgrp` - changel a file's group ownership

* `passwd` - change user password













#bash #expansions #permissions #unix #linux #chmod #octal #symbolic #ugoa #rwx
