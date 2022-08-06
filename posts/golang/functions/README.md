# Functions

###  What is a function?
* Function is used to group code. They main advantage is reusability.

###  Is function able to take 0 or >1 args?
* Yes.

###  Types in *Golang* are declared after variable name e.g (x int)

###  Func syntax:
```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(40, 10))
}
```

###  When >2 args in func param are in same type, you can omit writing typein every variable, instead you have to write type in last arg
	
```go
package main

import "fmt"

func add(x, y int) int {
	return x + y
	}

func main() {
	fmt.Println(add(40, 10))
        }
```
* Func listed above is taking 3 args, with *int* type

###  Functions can return more than 1 value
```go
func swap(x, y string) (string,string) {
	return y, x
}

// Note about function declaration above (after listed args), we declare what function returns, in this example its *(string,string)*
	
//  To consume *swap* func, do this:
func main() {
	a, b := swap("hello", "world")
fmt.Println(a, b)
}
	
// Now *a* variable is string of value "hello", *b* is "world"
```
###  Named returned values:
* In *Golang* returned values can be named, if they have names, they behave like variables defined in the top of a function. `func add (sum int) (x,y int)` - **values in parentheses after func param, is list of what function returns**

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

* In example above, values returned by **return** are 2. *x,y*. It is called **naked** return

* Note: **naked** return, should be used only in short functions, otherwise it impact readability.


### Function values

* *Functions* are also values. They can be provided as same as other values.

* Values which are *functions*, can be used as other *function* arguments and values which *func* can return

```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```


### Function Closures

* What is a closure?

	* Closure is a utility, which provides access to *variables* which are outside **current scope**

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

# Related to: 
* https://go.dev/doc/



#functions #function #types #funcsyntax #func #args #types #nakedreturn #return #functionasargument #functionasparameter #argument #parameter #closures
