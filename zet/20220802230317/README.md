# Finding files

* How do I link to a file or directory?
* What's a regular expression and why do I care?
* What's a glob?
* How do i find stuff?
* UNIX philosophy, filters, pipesm, etc.
* Difference between *whereis*, *find*?
	* So *find* is what you use when you want to search by particular criteria and also manipulate files. It has many more options than locate so allows far more fine-grained control of results. It is slow because it performs the requested test(s) on every file to see if it matches.

	* *locate* is used to scan the whole system quickly for something - you might use this when you have no idea where something is, or when you want to find all related files scattered across various unknown places. It's fast because it uses a binary database to index the system. To get new files to show up, first run *sudo updatedb* (the database it updated once per day by *cron*

	* the *whereis* command simply returns the location of the executables, the man pages and the sources of a program (see man whereis)
	* *locate* and *whereis* needs a db
* How do I save output to a file?
	* `<command> > file.txt`

* How do I find all the files in a directory and subdirectories?
	* `find .`

* Find only the  files *recursively*
	* `find . -type f`

* Find only directories beginning with *.*
	* `find . -type d -name ".*"`

* Find all files that have been created or modified in last hour?
	* `find . -mmin -60`

* Find all files that are bigger than 10MB
	* `find . -size 10M`

* Find all files with *setuid* perms on computer?
	* ``
	* https://stackpointer.io/unix/unix-linux-find-setuid-files/532/

* How do I silence errors?
	* `/dev/null` - everything which is send there disappears
	* You cannot recover it
	* What if we want to read this stuff which we wanted to filter to the `/dev/null`?
		* We need to redirect it to the other file `someshit 2>/dev/null >/tmp/foo` (you can name it as you want)
		* You can still redirect it to the terminal, but need tofind your current pts by `w` command, then instead of `>/tmp/foo` write `>/dev/pts/{n}`

* Diff between `/tmp` and `/dev/shm`?
	* `/tmp` is stored in a disk partition and `/dev/shm` is stored in RAM

* What is an inode?
	* An inode is a data structure that stores all the information about a file except its name and its actual data.
	* They are unique only on partition level, in two partitions you can have two files with same inode
	* `ls -i` - command which shows inodes of files
	* Inodes stores metadata about the file it refers to. This metadata contains all the information about the said file.
    		* Size
    		* Permission
    		* Owner/Group
    		* Location of the hard drive
    		* Date/time
    		* Other information



## Commands
* `find .` - sort of the same as `ls -l1`
* `find . | <pager>` -  find files with <pager(more | less)>
* `find . -ls` - sort of the same as `ls -l1`

# Relates to: 
* https://youtu.be/z7nluhUWYRs?t=4666
#files #find #locate #whereis #exercises
