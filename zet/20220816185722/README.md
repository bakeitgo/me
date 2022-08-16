# UNIX / Bash

* bash script is same as you would write it in cli, but as separated file.

* dont use #!/usr/bin/*env* bash ITS A HUGE VULNERABILITY TO USE SCRIPTS AS *env*, it looks for the first bash occurence in /usr/bin and uses it, so if someone will substitute this executable, you will run it

* `shellcheck` - checks is your script write good

* Comments in bash -  `# some comment`

* Dont use `./` in paths

### Commands

* `printf` - print string in formatted way (use `man` to see possible format control sequences

* `clear` - clear terminal screen

* `history` - display or manipulate the history list




#shellcheck #comments #clear #history #printf #shebang #path
