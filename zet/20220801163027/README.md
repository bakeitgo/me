# Editing files, SSH

* Where are all my files and why are they there?
* How do I remove add and rename files and directories?
* Dangers of *setuid* ?
* Console is a hardware, or a symulation of a hardware
* set UID - special permission (only able to be settled up on binary file) [it runs like the user who owns the file]
over the wire (practise suid [escalating privileges])
web servers are used to run as root (suid)
* What are permissions and why do I care? 
  * -rwxrw-r-- (2-4) file owner (5-7) group file belongs (8-10) every1 else
* Change permissions of a file (chmod)  (permissions are called octal, every octal is (my,group,world)[-rwxrwxrwx]
 
* Why IP Address changes? DHCP
* If file got <1kb, it is saved in inode table
(it is loaded instantly when you go into directory where files with <1kb are located)
* `Host` it is PC which is the root, like my PC right now is HOST, when i connect to VM via ssh, then VM is HOST
* What is a home directory? 
  * Its a directory where user can put his files
  * we need a user becauase he got less privileges than root
  * 


## Commands
* `ls -ld` - look directory permission on current directory
* `stat foo` - see all the details about the `foo` inode
* `sudo adduser foo` - interactively add a user named `foo` (not RedHat)
* `sudo deluser foo` - interactively del a user named `foo` (not RedHat)
* `sudo passwd foo` - change the password for foo
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
* `sudo groupadd <groupname>` add new permission group
* `sudo usermod -aG <groupname> <username>` add user to a group
* `ls -l $(which sudo)` - list perms for `sudo` command (dont use backticks, they *cant nest*)
* `echo foo` - write foo in stdout
* `cat foo` - write content of foo to stdout
* `which foo` - print th efull file path to the executable foo
* `chown foo rando` - change ownership of *foo* to *rando*
* `[u]ser [g]roup [o]thers` - list of users shortcuts to add permission
* `mv foo other` change file/dir *foo* name to *other*
* `cp foo other` - copy file/dir *foo* name to *other* | move
* `rmdir <name>` - remove directory *only if empty*
* `rm <name>` - remove file with <name>
* `rm -rf <name>` - force file remove with <name>
* `sudo su - <username>` - login as username without logging off
* `sudo su -` - login as a root without logging off
* `touch <name>` - create file with <name>
* `mkdir <name>` - create dir with <name>


## Relates to:

* https://youtu.be/ba0v04Q9Fqk?t=8244
* https://linuxize.com/post/how-to-add-user-to-group-in-linux/
* https://www.geeksforgeeks.org/permissions-in-linux/
##
#UNIX Timeshare #adduser #stat #Host #permissions #editingpermissions #chmod #sudo #passwd #deluser #addgroup #modifygroup #details #stat
