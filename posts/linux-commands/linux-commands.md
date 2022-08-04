# Linux commands

* `pwd` / print working directory
* `ls` / show files in current folder
* `cd` / move to the another folder
* `set -o noclobber` - stop from blowing away files
* `ls -a` / show all files including hidden
* `ls -al` / show all files including hidden begin with `.` long named
* `shutdown -h now` / `poweroff` / `init 0` // shutdown pc now
* `mv` / rename source to dest, or move source to directory
  * Be careful! If you do mv <filename1> <filename2> first file gonna
  * overwrite second file!
* `sudo` / do it as administrator
* `mv -i foo other` - change filename but protect against overwrites
* `scp foo target:` - copy *foo* from *host* to remote target *home dir*
(def)
* `apt` / advanced packaging tool package manager
* `apt upgrade` // upgrade *all* packages to latest version
* `apt-get` when you are running a script because its supporting backward compatibility as much as he can
* `apt update` - update all the sources for packages
* `apt search ^neo` - search for all packages starting with neo in name
* `hostname` - display name of host
* `cd foo` - change into foo directory
* `cd`, `cd ~` - change back to the home directory
* `cd ..` - change to the parent directory
*  `cd ../..` - change into parent of the parent directory
*  `cd -` - change to previous chosen directory
*  `cd /` - change to the root directory
*  `sudo apt install neofetch` - install `neofetch` package and its dependencies
*  `sudo apt remove neofetch` - uninstall `neofetch` package
*  `sudo apt autoremove` - automatically remove unused packages
* `ls -ld` - look directory permission on current directory
* `stat foo` - see all the details about the `foo` inode
* `sudo adduser foo` - interactively add a user named `foo` (not RedHat)* `sudo deluser foo` - interactively del a user named `foo` (not RedHat)* `sudo passwd foo` - change the password for foo
* `passwd` - change own password
* `chmod -r` - remove read permission from all users(me,group,world)
* `chmod -w` - remove write permission from all users(me,group,world)
* `chmod -x` - remove execute permission from all users(me,group,world)
* `chmod +r` - add read permission from all users(me,group,world)
* `chmod +w` - add write permission from all users(me,group,world)
* `chmod +x` - add execute permission from all users(me,group,world)
* `chmod u+r` - add read permission to user (me)
* `chmod g+r` - add read permission to user (group)
* `chmod ug+r` - add read permission to users (me,group)
* `chmod ugo+rwx` - add read write execute permissions to (me,group,others)
* `sudo groupadd <groupname>` add new permission group
* `sudo su - ` - effectively login as root without logging out
* `sudo su - foo` - effectively login as foo user
* `su - <name>` - effectively login as <name> user (change current user
* `stat -c '%a'` - see octal permissions
* `sudo usermod -aG <groupname> <username>` add user to a group
* `ls -l $(which sudo)` - list perms for `sudo` command (dont use backticks, they *cant nest*)
* `echo foo` - write foo in stdout
* `cat foo` - write content of foo to stdout
* `which foo` - print th efull file path to the executable foo
* `chown rando foo` - change ownership of *foo* file to *rando* user
* `[u]ser [g]roup [o]thers` - list of users shortcuts to add permission
* `mv foo other` - change file/directory *foo* name to *other* (or move)
* `cp foo other` - copy file/directory *foo* to other directory
* `touch` - create new txt file or update last modified
* `rmdir <name>` - remove an *empty* <name> directory
* `rm -rf <name>` - remove file / directory and everything in it
* `grep jill /etc/passwd` - list only line containing *jill*
* `ls -inl` - show list of files but with number of uid/gid
* `chown -R user:group <olddir>` - recursively change ownership/group
* `mv foo other` change file/dir *foo* name to *other*
* `cp foo other` - copy file/dir *foo* name to *other* | move
* `rmdir <name>` - remove directory *only if empty*
* `rm <name>` - remove file with <name>
* `rm -rf <name>` - force file remove with <name>
* `sudo su - <username>` - login as username without logging off
* `sudo su -` - login as a root without logging off
* `cp -ar foo other` - copy *foo* to *other* keeping timestamp
* `tr <set1> <set2>` - replace set1 characters with set2 characters e.g.`echo linuxize | tr 'lin' 'red'` will print to stdout *reduxize*



# Relates to: 
* http://www.linuxcommand.org

#commands #unix #linux #cheatsheet #shell #bash #permissions #dir #file #directory #owner #users #groups #others
