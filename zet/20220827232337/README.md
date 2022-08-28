# Bash

* Compound commands

	* (list) is executed in the subshell environment

	* { list; } is executed in current shell environment

	* ((expression)) - if value of expression is non-zero *exit status* is *0*, otherwise *1*

	* [[ expression ]] - conditional expression *(returns 0 || 1 depends on its evaluation*

	* `nocasematch` - shell option which disables cases of alphabetic characters

	* `=~` - matches based on PERE 

* Parameters

* Entity that stores value

	* Special parameters - see `man bash` e.g. `$#` expands number of positional params in decimal

	* Positional parameters - parameters which are assigned from the shell arguments when it is invoked, can be reassigned via `set` command. Starts from $1...$n *$0 is the name of the script itself*


##  Shell variables

* Environment variables 

	* Environment variables - effective way to pass info about current environment to script. At execution time, program receives array of env variables in *key=value* format.

	* Environment variables cna be temporarily changed, by prefixing command with *key=value*. *Naming convention* to env variables used by shell = Uppercase letters, digits, underscore characters. Cannot begin with digit. *Naming convention* to env variables only used by app = lowercase characters

	* `printenv` -- display env variables

	* Variable can be altered via `env` command


* Special variables

	* Only accessible to the shell execution environment. Limited to the current instance of shell
	
	* More here: [About env variables](https://www.shell-tips.com/bash/environment-variables/#gsc.tab=0)


## Related to:

* https://www.shell-tips.com/bash/environment-variables/#gsc.tab=0




#compoundcommands #expressions #nocasematch #=~ #PERE #POSIX #REGEX #list #shellvariables #positionalparameters #specialparameters #parameters #environmentvariables
