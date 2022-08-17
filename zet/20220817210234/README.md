# Processes


* Processes are how Linux organizes the different programs waiting for their turn at the CPU
	
	* Each process is assigned a number called *PID* (process ID), in ascending order
	
	* We can view processes using `ps` (TTY means 'teletype' refers t o the *controlling terminal* for the process) (TIME is metric of how much CPU time process is consuming) (we can set higher prio to process which means it is consuming more CPU time (more resources are used by this process (works *faster/better*) Process like this is called *lessnice*

	* `ps` provides snapshot of processes, to see how process works *live* use `top` command (it refreshes snapshot by default every 3 secs)
	
	* We can place programs to the `bg` (stuff hidden behind the surface) and `fg` (stuff visible on the surface)

	* To run program in the bg, use its name followed with ampersand, e.g. *exeprogramname &*

	* In order to place program to the fg, first use `jobs` command,and take a *jobspec* (job number), then run `fg %jobspec`

	* In order to place program to the bg, press CTRL+Z which stops(pause) a process, then run `bg %jobspec`

**
### Signals

* How to terminate a process?

	* Via `kill` command followed with *PID* / *%jobspec*, `kill` command doesn't kill processes, it sends them *signals*. We can specify which signal we want to send, *default is TERM(terminate)*


### Shutting down system

* How to shutdown a system?

	* `halt`

	* `poweroff`

	* `reboot`

	* `shutdown`

# Commands

* `ps` - show processes

* `jobs` - show list of jobs launched from our terminal

* `fg` - return program to the foreground

* `bg` - return program to the background

* `kill` - send signal to a process

* `killall` - send signal to all processes with *name*








#processes #ps #top #jobs #bg #fg #kill #killall #shutdown #signals #jobspecs
