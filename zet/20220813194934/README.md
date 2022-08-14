# bash / configuring vim

## UNIX Filter

* `:.! command` - place stdout of `command` in the current line

* `:.,.+2!bash` - run code from current line to current line + 2, and use bash to interpret it.

* `nl` - numerise all lines

* `!}` - select couple of lines on which we will make operations. e.g. it selects `.,.+9!`  

* `.` - `.` repeats last used command

* `:help` - opens VIM help

* Aliases - they are not working in the files with (bang bang) (:.!)


* difference between compiled / interpreted code


	* nothing is running on pc when it is not compiled (it is not listed in the process list `ps`)
	
	* if we call something a program mostly we think that it has been compiled

	* main difference is that compiled program is runned as binary file and it is showed in the processes and takes pc memory out, and it is faster

	* interpreted languages need interpreter

* How pc looks for a scripts as binary file and as text file

	* binary file is assuming pc memory and it is displayed as a normal process

* How running program looks like?

	1. Firstly, we look is called program existing by its `PATH`
	1. Secondly, pc looks is user can execute program (have x permission)
	1. Thirdly, it looks is it binary or not, if yes it runs it, if it is an ascii text file, it opens it and reads first line which specifies what is this file, for e.g. is it py or bash script, then it reads it line by line by line by line by line

     1	thING
     2	thING
     3	thING
     4	thING
     5	thING
     6	thING
     7	thING
     8	thING
     9	thING
    10	thING

Sun Aug 14 12:03:25 CEST 2022
***
# Related to:

* https://youtu.be/yu2kO2yb7L8?t=5057

#bash #vim #scripting #scripts
