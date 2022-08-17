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


### Default permissions

* How to set default permissions on newly created file?

	* `umask` - it is using octal notation

	* default octals are *0002* `-rw-rw-r--` and *0022* `-rw-r--r--`
	
	* we use to write octal permission mask as four digits number, in case of permission settings.

### Superuser

* When we need to access files which ordinary user cannot access, we can setuid to the superuser, which changes *effective userid* to the *file owner*. Mostly it is used in programs which are owned by superuser. Via this, ordinary user have privileges and acts as superuser (got access tofiles / directories to which normally would be inaccessible. `chmod u+s program` listing files with `ls` shows it like that e.g. `-rwsr-xr-x`

	* We can do it to the *gid* also `chmod g+s dir` listing files with `ls` shows it like that e.g. `drwxrwsr-x`


### Sticky bit

* We can mark exe file as *not swappable*. On files it is ignored, but it works on directory (where inside are files). It prevents users from deleting / renaming files until they are directory owner, file owner, superuser. 

	* `chmod +t dir` listing files with `ls` shows it like that e.g.`drwxrwxrwt`

### Changing identities


* We can change identities, by:

	* log out and log back in as other user

	* use the `su` command

	* use the `sudo` command
**
* The sudo command allows an administrator to set up a configuration file called /etc/sudoers and define specific commands that particular users are permitted to execute under an assumed identity.

* `su` - using with `-l` / `-` (it is the same) means that we are going to load user envioronment and his shell to use.

	* We can exit session via `exit` command

	* With `-c` option, we gonna execute single command via specified user
**
* `sudo` - with this command we dont need to have password to superuser, we are using our password to authenticate (with this command we can run exe files which requires *superuser (root)* privileges.

	* it doesn't starts new shell and doesn't load other user environment, but it can, if we use `-i` option, we can list privileges we gains by sudo by using `-l` option

### Change file owner / group

* `chown [owner][:[group]] file` - change file owner/ group

* `chgrp` - change file group

* by setting gid perms as superuser to the directory, all files will be created with the group same as in directory




### Commands

* `id` - display user identity

* `chmod` - change a file's mode

* `umask` - set the default file permissions

* `su` - run a shell as another user

* `sudo` - execute command as another user

* `chown` - change a file's owner

* `chgrp` - changel a file's group ownership

* `passwd` - change user password













#bash #expansions #permissions #unix #linux #chmod #octal #symbolic #ugoa #rwx #chown #chgrp #passwd #su #fileowner #defaultpermissions #umask #chmod
