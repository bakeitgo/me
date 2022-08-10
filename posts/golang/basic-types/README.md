# Basic types

* Which basic types occur in *Golang*?
	* `bool`
	* `string`
	* `int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr`
	* `byte` -> alias for `uint8`
	* `rune` -> alias for `int32`
	* `float32 float64`
	* `complex64 complex128`

* `uint` can only have positive values, `int` doesn't care 

* `var` can be formuled in blocks like in `import`
	```go
	package main

	import (
		"fmt"
		"math/cmplx"
	)

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	func main() {
		fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
		fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
		fmt.Printf("Type: %T Value: %v\n", z, z)
	}
	```

* Types `int uint uintptr` generally have 32bits length at 32-bits systems and 64bits length at 64-bits systems
	* If you want *total* value, you should use `int`, unless you need for some reason value in specific bit length.


### Type conversion

* Expression `T(v)` convert value `v` to `T` type
	* e.g 
		1. `var i int = 42`
		1. `var f float64 = float64(i)`
		1. `var u uint = uint(f)`
	* or easier:
		1. `i := 42`
		1. `f := float64(i)`
		1. `u := uint(f)`


### Type inference

* When you declare variable without type (for e.g using i := 1 or var i = 1) Variable type will be inferred
	* e.g `
	var count int
	nextcounter := 1 // nextcounter is in same type as count
	`
* But it dangerous because, if you dont declare type, *Golang* will infer type from right side like this:
	* `i := 42 // int
	   f := 3.124 // float64
	   g := 0.867 + 0.5i // complex128
	   `

### Operators

```go
package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
}
```
### bit operators

```go
package main

import (
	"fmt"
)

func main() {
	a := 10 // 1010
	b := 3 // 0011
	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	fmt.Println(a ^ b) // 1001
	fmt.Println(a &^ b) // 0100
}
```

### bit shifting

```go
package main

import (
	"fmt"
)

func main() {
	a := 8 // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0

}
```

# Related to: 

* https://go.dev/doc/


#golang #types #string #bool #int #float #complex #typeinference #inference #typeconversion #conversion



