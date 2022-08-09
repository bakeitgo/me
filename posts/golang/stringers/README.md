# Stringers

* What is a stringer?

    * Interface defined by the `fmt` package

    * A `Stringer` is a type that can describe itself as a `string`, It's used by `fmt` pckg to print values

    ##
```go
type Stringer interface {
    String() string
}
```
```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```



# Related to:

* https://go.dev/doc/


#golang #stringers