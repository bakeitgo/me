# Networking

* In order to check is network working, we use `ping domainname` command

* `traceroute domainname` - see all *hops* traffic takes to get from the file system to a specified host

* `ip a` - shows network interfaces and routing table

* `ftp` - transfer programs over network, it communicates ftp servers that let us to upload / download files over a  network

	* `ftp fileserver` - connect to the ftp server fileserver

	* `lcd Desktop` - change directory on local system to ~/Desktop

	* `get filename` - get file filename

	* `bye` - log off ftp session (`quit` / `exit` also works)

* ftp is not secure because it sends password with clean text

* we can download files from *ftp / web servers* using `wget`

	* via `wget` we can download files in bg, and recursively

### ssh

* In old days in Linux we were using `rlogin` / `telnet` to remote login to a system, it has a huge vulnerability same as `ftp`, it sends clear text over a network, this is a main reason why `ssh` has been created.

* ssh solves two basic problems of secure communication with a remote host:

	* Authenticates remote host via fingerprint for RSA key (prevents mitm attacks)

	* Encrypts communication

* By default it connects to port 22.

* To enable system receive remote connections, it needs to have *OPENSSH-server* installed, configure, running and must allow incoming network cnnections on TCP port 22.

* Besides opening a shell connection, we can execute a single command on a remote system, and display it results on local system

	* In order to omit expansion / put file on remote system  we need to wrap commands which we want to execute in single quotes ''

	* OpenSSH pckg provides also `scp` and `sftp`. `scp` works same as `cp` command to copy files (secure copy). `sftp` works same as `ftp` but it uses encrypted tunnel to copy files (it doesnt requires ftp server, only needs SSH server.

# Commands

* `ping` - send an echorequest to network hosts

* `traceroute` - print the route packets trace to a network host

* `ip` - show/ manipulate routing, devices, policy routing, tunnels

* `netstat` - print network connections, routing tables, interrface stats, masquerade connections, multicast memberships

* `ftp` - internet file transfer program

* `wget` - non-interactive network downloader

* `ssh` - ssh client (remote login)



#ssh #ftp #wget #networking #ip #netstat #traceroute #ping #unix
