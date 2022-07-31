# How to connect to our VM? (Secure Shell) -!

* What is NAT (Network Address Translation) and Bridged
* What is an IP address?
* What is `git-bash` and why would I need it?
* What about `putty` for ssh access? - It's useful because we doesn't need admin privileges to connect via ssh
* Why can't i use 127.x.x.x address? - It's a loopback address, it's always redirecting to the our PC
* How can you test ssh connection to *localhost*?
* What is a "user" and "userid"?
* How can i see stuff scrolling off screen? need to use *pipe* `| less` || `| more` (`q` to quit) `/ssh` 
search for all packages with `ssh` in name
* What package need sto be installed to connect with ssh? OPENSSH
* What is the correct `ssh` connection address?
* Remote access considerations? - You need to be in same network / SECURITY (if you want to set up serverr ouitside in the cloud, you gotta be hacked
* Setting up a Digital Ocean account
* How to setup GCP "nano" server for free?
* Running VM headless and connecting with ssh
  * via GUI
  * From (DOS) CLI
* narc project (spying on all file changes)


## Commands
*  `ip a` - show all ip addresses
*  `clear` - clears the screen
*  `which ssh` - display full path to ssh program
*  `type ssh` - ?
*  `ipconfig /a` - show all ip addresses on windows
*  `users` - short name of all logged in users
*  `who` - display who is logged in and how
*  `last` - summary of last logged in users
*  `whoami` - print effective user name/ID
*  `id` - display user and group IDs (here you can escalate privileges) [shows permission groups]
*  `w` - display logger version of who is logged in
*  `<Ctrl>-c` - interrrupt whatever (exit)
*  `<Ctrl>-d` - send "end of data/file"
*  `sudo apt install openssh-server` - install ssh server (if ssh doesn't works)
*  `vboxxmanage startvm <VMNAME> --type headless` - start *vm* headless

## Related

* https://www.virtualbox.org/manual/ch06.html
* https://www.washingtonpost.com/technology/2022/05/19/hacking-cfaa-justice-policy/


## Quick RECAP

* install package - `sudo apt install {{package-name}}`
* upgrade all package - `sudo apt upgrade`
* update all packages (requires to install update for specified package after) - `sudo apt update` then `sudo apt install {{package-name}}`
* use apt-get in scripts, because it keeps backward compatibility as long as he can
* sudo apt autoremove -- removes all unused packages


#ssh #VM #openssh #apt #apt-get #ip #ipconfig #user #userid #who #w #last #users #id #permission #groups #root #adm #putty
