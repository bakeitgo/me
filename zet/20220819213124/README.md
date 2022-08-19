# Bash


### Variables

* Constants are written in uppercase

	* in order to provide immutability use `declare -r` (-r -> read-only option)

* Variables are written in lowercase

* The first character of a variable name must be either a letter or an underscore.

* Spaces and punctuation symbols are not allowed.

* Variable names may consist of alphanumeric characters (letters and numbers) and underscore characters.

* Shell doesn't care about types of variables *TREATS ALL AS **STRINGS***

* Multiple variable assignments can be written in single line
	
	* `a=5 b="a string"`

* In order to combine variable within string we use curly braces:

```bash
name="john"

echo ${name}1 // prints to stdout john1

```

### Functions

* How define a function?

```bash
function name {
	commands
	return
} // or like that

name () {
	commands
	return
}
/ to invoke function we use its name
```

* Shell function needs to be declared before initialization (must be located higher than initialization)

* Declaring variable inside shell function is in local scope (works onlyin this scope), other variables are global scope

### Flow control

```bash
x=5

if [[ "$x" -eq 5 ]]; then
		echo "x equals 5."
else
		echo "x does not equal 5."
fi
```

### Exit status

* By default everything which returns 0 is successful action, otherwise it fails (range 0..255)

	* How to check it? `echo $?` displays last returned value

* Expressions works same as functions (0 - false 1 - true)

* `if` is frequently used with `test` option (performs variety checks / comparisons)





#function #exitstatus #bash #unix #flowcontrol #variables
