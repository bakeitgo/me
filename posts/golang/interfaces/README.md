# Interfaces

* What is an interface? 

    * Interface type is defined as a set of method signatures. It's like a contract which have set of functions, which is implemented to fulfill it. It provides better abstraction.
    ##
```go
package main
 
import "fmt"
 
// declare interface
type Dog interface { 
    Bark()
}
 
// declare struct
type Dalmatian struct {
    DogType string
}
 
// implement the interface
func (d Dalmatian) Bark() {
    fmt.Println("Dalmatian barking!!")
}
 
func MakeDogBark(d Dog) {
    d.Bark()
}
 
func main() {
    d := Dalmatian{"Jack"}
    MakeDogBark(d)                    // Dalmatian barking!!
}
```

* How interfaces are implemented?

    * Implicitly, there is no keywords like `implements` etc.

    * Interface values with nil underlying values.
    ## 
    If the concrete value inside the interface its nil, function will be called with a nil receiver
    * *Note that an interface value that holds a nil concrete value is itself non-nil.*

```go
package main

import "fmt"

//declare interface
type I interface {
	M()
}

// declare struct
type T struct {
	S string
}
// implement abstract function declared in interface
func (t *T) M() {
	if t == nil { // check is func called with nil type
		fmt.Println("<nil>") // if yes print <nil>
		return
	}
	fmt.Println(t.S) // else print string
}

func main() {
	var i I // declare i which type is interface
	fmt.Println(i) // <nil> is printed out (because value of interface is nil)
	var t *T // declare t which type is struct
	i = t // t (STRUCT) implements i (INTERFACE)
	describe(i)
	i.M() // call M() [abstract function declared in interface] and print nil because value is nil

	i = &T{"hello"} // &T{"hello"}(consuming pointer with "hello") implements i (INTERFACE)
	describe(i)
	i.M() // call M() [abstract function declared in interface] and print hello
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

* Under the hood interface values may be considered as a tuple `(value, type)` - Interface holds a value of a concrete type. Calling a method on an interface value, execute it on its underlying type

* Calling methods on a `nil` interface causes error because there is no type inside the interface tuple to indicate which *concrete* method to call

* Empty interface - empty interface may hold any values of any type. They are used by code that handle values of *any* type

## Type assertion

* What is a *type assertion*?

    * It provides access to an interface value underlying concrete value
    ##
    `t := i.(T)`

    * This statement asserts that interface value holds the concrete type, and assigns this underlying value to the variable `t`

    * If `i` doesn't hold type `T` it trigger a *panic*

    * To *test whether interface value is type T*, type assertion can return two values, *underlying value* and *boolean* which reports is assertion succeeded.
    ##
    `t, ok := i.(T)`

    * If i holds a T, then t will be the underlying value and ok will be true. 

    * If not, ok will be false and t will be the *zero value of type T*, and no panic occurs. 

## Type switches

* What is type switch?

    * Same concept as type assertion, but switch permits several type assertions in series

    * It's like regular switch statement

    * It have same syntax as regular type assertion, but `i.(T)` is replaced with `i.(type)`
```go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

# Related to:

* https://go.dev/doc/

#golang #interfaces

