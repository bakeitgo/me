# Environment

* What is stored in env?

	* Env variables and shell variables, shell variables are bits of data placed there by bash, env variables are everything else

	* `set` - display both shell / env variables and shell functions(output is sorted alphabetically)

	* `echo $VARIABLE` - print variable to stdout

	* `alias` - display list of aliases

* env have built-in variables, e.g. SHELL, HOME, PAGER, OLDPWD, PWD, USER

### Packages

* All distros depends mostly on package managers, some of them are checked by distro team, some not. We describe two levels of package tools

	* Low level package tool, which handles task such as installing / removing packages, by this tool we can install from source other than repo

	* High level package tool, which perform dependency resolution and metadata searching (`apt-cache search search_string`), by this method we can install packages from repo with dependencies



# Commands

* `printenv` - print part or all of the environment

* `set` - set shell options

* `export` - export envto susbequently executed programs

* `alias` - create alias for command

* `apt-cache search search_string` - search repo for package (it base on name / desc)

* `apt-get update` - update pckgs 

* `apt-get install pckg_name` - install package

* `dpkg -i pckg_file` - install / upgrade package from source other than repo

* `apt-get remove pckg_name` - remove package

* `dpkg -l` - list installed pckgs

* `dpkg -s` - list is specified pckg installed

* `apt-cache show pckg_name` - display info about installed pckg

* dpkg -S filename` - determine which packages is responsible for installed file



#env #environment #packages #pckg #apt #dpkg
