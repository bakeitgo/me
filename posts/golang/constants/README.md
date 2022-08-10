# Constants

* Constants are declared same as variables but using `const` instead of `var`

* Constants types can be: 
	* `string`
	* `text characters`
	* `bool`
	* `number types`

* They cannot be *initialised in runtime*

* They can be *shadowed* in inner scope

* Type inferrence is working here same as in variables `const a = 42 // type is int`

	* But, with `const` type is flexible when it comes to operations like *add* e.g. `var b int16 = 30` then `a + b // 72`, when we will try to do it with `var a = 42`, error will be printed (mismatched types int and int16)


```go
package main

import (
	"fmt"
)

const a int16 = 27

func main() {
	const a int = 14
	fmt.Printf("%v, %T\n", a, a) // 14 is printed
}
```

### Special symbol *iota*

* What is `iota`?

	* It's a counter what we can use when we create numerated constants

```go
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const ( // we can omit writing iota next to a constant, compiler interprets it.
	_ = iota // ignore first valuie by assigning to blank identifier (memory address saver)
	a		 // 1
	b		 // 2
	c		 // 3
)

const ( // when we declare next const block, iota will reset to initial value
	a2 = iota // 0
)

const (
	_ = iota + 5// 0
	a		 	// 6
	b 			// 7
	c			// 8
)
```

* Note: **CONSTANTS CAN'T BE DECLARED USING :=**

### Numeric constants

	* Numeric constants are values with *high precision*
	* Constant declared without type, initialise type via context














# Related to: 

* https://go.dev/doc/

#golang #constant #const #constants #types #type



