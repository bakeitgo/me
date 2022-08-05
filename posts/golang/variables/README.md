# Variables

### Declaring variables

* How to declare variable in *Golang*?
	* by using `var` instruction, it declares list of variables, like in funcs, first we write *name*, then *type*

* Where we can declare variables?
	* *variables* can be declared in *package* and *function* level

* Declaring variable examples.
	
	```go
	package main

	import "fmt"

	var c, python, java bool // in package scope

	func main() {
		var i int // in function scope
		fmt.Println(i, c, python, java)
	}
	```

* Default initial value of declared variables is: 
	* `0` for `numeric types`
	* `false` for `bool types`
	* `""` (empty string) for `strings`


### Declaring and Initializing variables

* Declared variable can have initializer, *1 initializer for 1 variable* 
	* When initializer is assumed, declaring type of variable is unnecessary
	
	* e.g `var i, j, c, python, java = 1, 2, true, false, "no!"`


### Short variable declarations

* *INSIDE OF A FUNCTION* we can use `:=` instead of `var`
	
	```go
	package main

	import "fmt"

	func main() {
		var i, j int = 1, 2
		k := 3
		c, python, java := true, false, "no!"

		fmt.Println(i, j, k, c, python, java)
	}
	```


* *OUTSIDE OF A FUNCTION* we **have** to use `var`, `func`, **etc.**
	
	```go
	package main

	import "fmt"
	
	var c, python, java := true, false, "no!" 

	func main() {
		var i, j int = 1, 2
		k := 3
		
		fmt.Println(i, j, k, c, python, java)
	}
	```



# Relates to: 

* https://go.dev/doc/


#golang #variables #var #packagelevel #funclevel #shortvariable #variabledeclaration #variableinitialise
