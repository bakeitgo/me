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

https://go-tour-pl1.appspot.com/methods/1



# Related to: 

* https://go.dev/doc/

#golang #methods