# bash

## Pipelines

* When using |& stdout and stderr is redirected to the *command2* (shorthand for *2>&1*

* In pipeline *stdout* is connected to *command2* before redirection, *stderr* is connected after redirection

* Pipelines in default *returns* with the last command *exit status*, but if *pipefail* is enables, it returns with the *rightmost* command which returned **0**

* If the *!* precedes a pipeline, it negates *return status*

* All commands are terminated before pipeline *returns exit status*

* When *time* precedes command, it displays execution time to the output for *real*, *user*, *sys*

* Each command in pipeline is executed in *separate process* ***(subshell)***

* List is a sequence of one or more *pipelines*

	* separated by *;*, *&*, *&&*, or *||*

	* *&&*, *||* have equal precedence, followed by *;*, *&*, *<newline>* which have equal precedence

* When a command is terminated by *&*, it is runned asynchronously, *it is runned in background, and shell doesn't wait until it finish, **shell returns exit status 0 in this mode*** 

* Commands separated by *;* are executed ***sequentially*** *shell waits until command finish execution***

* *&&* and *||* are executed with *left associativity*

	* command1 && command2 -> first runs command1, if command1 exit status is 0 (*true*) then command2 is executed

	* command1 || command2 -> first runs command1, if command1 exit status is != 0 (*false*) then command2 is executed

# References

* `man bash` - till line *354*

#bash #pipelines #redirections #stdout #stderr #sequential #asynchronous #; #& #termination #or #and #|| #&& #execution #subshell


