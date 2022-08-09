# Methods

* What is a method?

    * Method is equivalent of a *class* in other languages

    * It is a function with special **receiver** argument

    * Receiver appears in its own argument list between the `func` keyword and the method name `func (x someStructType) funcName`

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
} // In this example Abs() method has a receiver of type Vertex named v
  // Remember, its just a function, still we can write it like this:

func AbsButANormalFunc(v Vertex) float64 {
    return math.Sqrt(v.X * v.X + v.Y*v.Y)
}
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // func with receiver (class equivalent)
    fmt.Println(AbsButANormalFunc(v)) // normal function
}
```

* We can declare a method on *non-struct* types also

    * You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int). 


## Pointer receivers

* What does it mean?

    * If  you want to modify the value to which receiver points, you have to use it. 

    * Note: With a value receiver, we operate on copy (same memory address)

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) { // *Name < we point to the Vertex struct(same memory address) by what scale method operate on pointer value.
	v.X = v.X * f
	v.Y = v.Y * f
}

// rewrite scale as function: 

func ScaleWithNormalFunction(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) // scale it using method with value pointer
    ScaleWithNormalFunction(&v, 100) // scale it using normal function
	fmt.Println(v.Abs())
}
```
* Note: Function needs to take value as a pointer receiver, Method doesn't care about it.

    * Function: `var v Vertex ScaleFunc(v, 5)  // Compile error!`       `ScaleFunc(&v, 5) // OK`

    * Method: `var v Vertex v.Scale(5)  // OK`
              `p := &v p.Scale(10) // OK`

    * It's because method with pointer receiver always take pointer. e.g. *Golang* interprets value `v.Scale` always as pointer `(&v).Scale`


* Which one to choose? Value or pointer receiver?

    * There are 2 reasons to use *pointer receiver*

    1. When we want to modify the value to which receiver points

    1. To avoid copying the value on each method call (on each call value gets different memory address, if we dont point to this specific value (specific memory address))

    * Check example below:

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

* In example above, both Abs and Scale are with receiver type, but only Scale method needs it (cause it changes value to which receiver points)

* In general all methods should have either value or pointer receivers, but not a *mix* of both
https://go-tour-pl1.appspot.com/methods/1



# Related to: 

* https://go.dev/doc/

#golang #methods