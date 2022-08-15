# UNIX

### Pipelines / Filters

* What is a pipeline?

	* it reads data from stdin and sends it to stdout, we can consume this stdout via *pipeline* `|` that means stdout is sending to stdin to another program

	* couple of pipelines are described as ***Filters***


### Commands

| Command | Description |
|---------|-------------|
| uniq | Report or omit repeated lines |
| wc | Print line, word, byte counts |
| grep | print lines which matches pattern |
| head | print lines counting from top | 
| tail | print lines counting from bottom |
| tee | read from stdin, print to stdout, and save to file |



### How the shell sees commands?

* Expansion

	* using `*` in command e.g. `echo *` doesn't prints `*` to the stdout, the `*` means match every character in a current directory and print everything, but how? shell expands any qualifying character before invoking command, it means echo behaved as expected. **Remember a wildcard means match everything**

	* using `*` with `ls` command works as *pathname expansion* it prints all directories, subdirectories

	* using `~` means home directory

	* using `$((expression))` performs arithmetic operation, expressions can be nested e.g.  $(($((expression)) * $((expression))))

	* using {} expansion, we create a couple of operations. e.g. `ls {.,..}` list files in . directory and .. echo Number_{1,..5} prints to stdout `Number_1 Number_2 Number_3 Number_4 Number_5`, second option is called *postscript*, first is *preamble*.

	* *Note!* we can print range of letters in reversed alphabetical order {Z..A}

	* Command substition allows us to use output of command as an expansion. e.g. echo $(ls), ls -l $(which cp), *in older shells backticks`` works instead of $(expr)*


* Quoting

	* echo this is a         test, prints this is a test (shell removes extra whitespaces)

	* if we call an undefined variable,

	* Using *double quotes* word splitting, pathname / tilde / brace expansion is omitted and carried out as normal characters. *Note!* that parameter, arithmetic, command substitutions still works in ""

	* Using *single quotes* we omit all expansions.

	* We can omit a *single expansion* via using backspace character "\"


* Backslash escape sequences

| Escape sequence | Description |
|-----------------|-------------|
| \a | Bell (alert that causes the computer to beep |
| \b | Backspace |
| \n | Newline (on Unix like systems produces line feed |
| \r | Carriage return | 
| \t | Tab |

* In order to use backslash escape sequences we need to use dollar sign then wrap it in single quotes:

	* $'\a' 
	* $'\b' 
	* $'\n' 
	* $'\r' 
	* $'\t'





