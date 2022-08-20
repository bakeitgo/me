# Bash

* Use `command -v` instead of `which` in scripts

	* `which` has never been in POSIX standard

	* https://unix.stackexchange.com/questions/85249/why-not-use-which-what-to-use-then

	* `which` doesn't list is something alias

* Hardcode control path to not be !PWNED
	
	* e.g. `export PATH=/usr/sbin:/usr/bin:/bin

* Dont use `#!/bin/env`

* *Aliases*  doesn't work in scripts

	* Unless you define aliases to work as default cmnds

* Precede anything with backslash `\` to disable aliases, functions

















#bash #scripting
