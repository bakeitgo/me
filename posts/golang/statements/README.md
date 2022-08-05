# Statements


### If statement

* If statement looks like *for* loop
	* Conditional expression doesn't need to be surrounded with parentheses, but still needs block - `{ }` curly braces


	```go
	package main

	import (	
		"fmt"
		"math"
		)

	func sqrt(x float64) string {
		if x < 0 {
			return sqrt(-x) + "i"
		}	
		return fmt.Sprint(math.Sqrt(x))
	}

	func main() {
		fmt.Println(sqrt(2), sqrt(-4))
	}
	```


### Short instruction in *If* statement

* Like in *for* loop, we can Initialise variables before conditional expression
	
	```go
	package main

	import (
		"fmt"
		"math"
	)

	func pow(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		}
		return lim
	}

	func main() {
		fmt.Println(
			pow(3, 2, 10),
			pow(3, 3, 20),
		)
	}
	```


	* *Note* Variables declared in `if` statement, are accessible only  in `if` block scope, and `else` block.


### Switch statement

* Switch in *Golang* is same like in *Java* / *JavaScript* / *C* etc. bBut with some **exceptions**.
	
	* You dont need to provide `break` after every condition, this is provided automatically

	* Conditions doesn't need to be *constants*

	* Compared values doesn't need to be *integers*

	```go
	package main

	import (
		"fmt"
		"math"
	)

	func pow(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		} else {
			fmt.Printf("%g >= %g\n", v, lim)
		}
		// can't use v here, though
		return lim
	}

	func main() {
		fmt.Println(
			pow(3, 2, 10),
			pow(3, 3, 20),
		)
	}
	```

* Switch are performed from **top** to **bottom** and stops when condition coincides with the checked value

	```go
	package main

	import (
		"fmt"
		"runtime"
	)

	func main() {
		fmt.Print("Go runs on ")
		switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			// freebsd, openbsd,
			// plan9, windows...
			fmt.Printf("%s.\n", os)
		}
	}
	```



* Switch without main condition is same like `switch true`

	* It can be used when we want to use `booleans` in cases

	```go
	package main

	import (
		"fmt"
		"time"
	)

	func main() {
		t := time.Now()
		switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
			fmt.Println("Good evening.")
		}
	}
	```


### Defer 

* Defer - *delayed execution*

	* Defer delays function execution until surrounding function returns.
	* Deferred call args are evaluated immediately, but the function call is not executed until surrounded function returns.

	```go
	package main

	import "fmt"

	func main() {
		defer fmt.Println("world")

		fmt.Println("hello")
	}
	```


	* In example above first `hello` is printed, then `world`

	* Deferred function calls are pushed onto a stack, When a surrounded function returns, deferred calls are executed in *lifo* order

	```go
	package main

	import "fmt"

	func main() {
		fmt.Println("counting")

		for i := 0; i < 10; i++ {
			defer fmt.Println(i)
		}

		fmt.Println("done")
	}	
	```


## Related to: 

* https://go.dev/doc/


#golang #statement #if #switch #ifelse #conditions

