# regex, filtering, stuff



* How do I only show top five lines?
	* PIPELINES `<command> | <nextcommand>`
	* `head -5 <filename>` - by default is 10 lines

* How do I only show bottom five lines?
	* `tail -5 <filename>` - by default is 10 lines

* How to read files reverse sorted?
	* `tac <filename>` - reverse cat of file

* How to avoid using `cat`?
	* It's bad to use it in bash script
	* Why?

* Unix philosophy - 
	* Keep things modular
	* Make each program do one thing well. To do a new job, build afresh program.
	* Expect the output of every program, to become an input to another. Dont clutter output with extraneous information. Filter stderr / stdout
	* Design and build software, even operated system to be tried early, ideally within weeks - Usage of Agile meth. in 1978

* How to append lines to a file?
	* use `>>` instead of `>` e.g cat <filename> >> /dir/<filename>

* Difference between `more` and `less`
	* `more` is older xD
	* in new systems there is no `more` anymore


## Commands: 
* `head -n <file>` - read n top lines of a file
* `tail -n <file>` - read n bottom lines of a file
* `tac <file>` - print lines of a file in *reversed* order (bottom-top)
* `wc -l` - print count of lines
* `nl` - add line numbers to output
* `>` - (over) write to file
* `>>` - append (concat) lines to end of a file
* `|` - pipeline e.g *<command> | <secondcommand>* (get <command> *stdout* to <secondcommand> *stdin*
* `<` - send file to *stdin* e.g *<command> < <filename>* (get filename  *stdout* to *<command> stdin*
* `tee` - save *stdout* to a file, and send *stdout* to pipeline


#stdout #stdin #stderr #UNIX #tee #file #overwrite #>> #> #append #concat #readtoplines #readbottomlines #tail #head #tac #reversereadfile #tee


