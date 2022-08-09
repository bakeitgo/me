# Errors

* How *Golang* express error state?

    * Via `error` values

    * It's a built-in interface, like `fmt.Stringer`


    ##
```go
type error interface {
    Error() string
}
```

* Functions often return an `error` value, and calling code should handle them by checking whether it equals `nil`

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

* **A `nil` error denotes success**

* **A `non-nil` error denotes failure**

```go
package main

import (
	"fmt"
	"time"
)
//declare and initialize struct
type MyError struct {
	When time.Time
	What string
}
// implement Error() interface, customize printed string by stringer
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
// return error with timestamp and message
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil { // check error occurs
		fmt.Println(err) // print it // 
        // at 2009-11-10 23:00:00 +0000 UTC m=+0.000000001, it didn't work
	}
}
```

```go
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
    // Note: A call to fmt.Sprint(e) inside the `Error` method will send program to an infinite loop, we can avoid this by converting value to float64
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	if x < 0 { // check is x negative, if yes return non-nil value (failure)
		return x, ErrNegativeSqrt(x)
	}
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z, nil // if everything is ok return z and nil value (success)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```

# Related to: 

* https://go.dev/doc/


#golang #errors