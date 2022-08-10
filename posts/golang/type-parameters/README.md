# Type parameters

* What is a type parameter?

    * It is designed to work with multiple types in functions, it appear between brackets
    ##
    `func Index[T comparable](s []T, x T) int`

    * This declaration means that `s` is a slice of any type `T` that fulfills the built-in constraint `comparable`. `x` is also value of the same *Type*

* What is `comparable` constraint?

    * It is useful because it makes possible to use the `==` and `!=` operators on values of the type. In this example, we use it to compare a value to all slice elements until a `match` is found. This `Index` function works for any type that supports comparison.

```go
package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
```

* Generic types

    * In addiction to generic functions, *Golang* also supports generic types. Type can be parameterized with a type parameter.
    ##
    `type List[T any] struct {next *List[T] val T}`






# Related to: 

* https://go.dev/doc/

#golang #type-parameters #type parameters