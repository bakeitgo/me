# UNIX commands rewind

### `ls <options> <directory>` - list files

| Command | Option | Long option | Description |
|---------|--------|-------------|-------------|
| ls | -a | --all | List all files including hidden |
| ls | -A | --almost-all | List all files including hidden **except current and parent directory `.`,`..`**
| ls | -d | --directory | List contents of directory, not directory itself |
| ls | -F | --classify | Appends indicator character to the end of each listed name. e.g. it appends forward slash `/` if the name is directory
| ls | -h | --human-readable | display file sizes in human-readable format rather than in *bytes*
| ls | -l | display results in long format |
| ls | -r | display results in reverse order | 
| ls | -S | sort results by file size | 
| ls | -t | sort by modification time |


* Pagers in UNIX - `less` and `more` -> less is an improvement of an earlier UNIX program called `more`
	* pgup or b -> scroll bakc one page
	* pgdown or space -> scroll forward one page
	* up arrow / down arrow -> scroll up/down one line
	* G -> move to the end of file
	* g -> move to the beginning of a file
	* /characters -> search for next occurence of `characters`
	* n -> search for next occurence of previous search
	* h -> display help screen
	* q -> quit less
	* reset -> reset terminal

### directories in linux

* directories 

	* `/` -> root directory
	* `/bin` -> conains binaries (programs) that must be present for the system to boot and run
	* `/boot` -> contains the Linux kernel, initial RAM disk image ( as large as much drivers needs it at boot time) and the boot loader
	* `/dev` -> special directory which contains *device nodes*, here kernel maintains a list of all devices it understands (device files are abstraction of standard devices. We communicate with them via I/O system calls) There are 2 types of files inside, character and block. Those files are interfaces to character / block devices. The difference between them is how system reads data off them. In character devices, the data is a single character (bytes), in block devices we deal with blocks of data. Character devices are sound cards, serial ports, whereas Block devices are hard disks or USBs. Via `ls -l` we can identify which one is block/character file. block file starts with `b`, character file with `c`

	* `/etc` -> contains a system-wide configuration files, also contains collections of shell script that start each of the system services at boot time. e.g. files here is a /etc/passwd (list of the user accounts, uids, gids)

	* `/home` -> each user is given a directory in /home. Ordinary users can write files only in **THEIR** home directories.

	* `/lib` -> contains shared library files which are used by core programs. (same as DLL in Windows)

	* `/lost+found` -> each formatted partition, will have this directory. It is used to recovery files from a file system corruption event.

	* `/media` -> This directory contains mount point for removable media such as CD's DVD's, USB drives etc. They are mounted automatically after insertion.

	* `/mnt` -> This directory was on old linux systems, it is containing mounting points for mounted devices manually

	* `/opt` -> This directory is used to install *optional* software. Mainly used to hold software products that *might* be installed.

	* `/proc` -> It is a **virtual** file system maintained by the Linux kernel. They are peepholes to kernel itself.

	* `/root` -> home directory for root account

	* `/sbin` -> directory which contains system binaries. These are programs that performs vital system tasks. Generally reserved by superuser
	* `/tmp` -> directory where *temporary* files are stored. Some configuartions can cause to clean this folder after system reset

	* `/usr` -> Contains all programs and uspport files used by regular users

	* `/usr/bin` -> contains exe files installed by Linux distro

	* `/usr/lib` -> libraries for exe files in `/usr/bin`

	* `/usr/local` -> programs for system-wide use are stored here. By default it is empty, until sysadmin puts something here.

	* `/usr/sbin` -> contains sysadmin programs

	* `/usr/share` -> contains all shared data used by `/usr/bin` exe programs

	* `/usr/share/doc` -> contains documentation for installed packages (every installed package will put here its documentation

	* `/var` -> data that is likely to change is stored here (db's, spool files, user mail etc.)

	* `/var/log -> contains log files (various system activity). Should be monitored from time to time. For some systems, we have to be superuser to view them. Most useful ones are `/var/log/messages` and `/var/log/syslog`

***
###  cp, rm Options

| Command | Option | Long Option | Description |
|---------|--------|-------------|-------------|
| cp | -a | --archive | Copy the files / directories and all of their attributes. Normally copies take only default attributes of user performing the copy. |
| cp, rm | -i | --interactive | Before overwriting existing file, propmt user for confirmation. If omitted, cp works silently |
| cp, rm | -r | --recursive | Recursively copy directories / files and their content. Recursively delete files / directories, means if directory contains subdirectories, delete them also. | 
| cp | -u | --update | When copying files, copy only those that not exists or are not updated (modification date is older). |
| cp, rm | -v | --verbose | display messages as the copy is performed | | rm | -f | --force | ignore nonexistent files and do not prompt. **Overrides -i option** |


#ls #commands #unix #linux #gnu #cp #mv #directories #unixdir 
