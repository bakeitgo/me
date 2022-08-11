# Goroutines

* What is a goroutine?

    * It's an lightweight thread managed by the *Go* runtime.

* How to use it?

    * As simply as it sounds. Using `go` keyword before function call

* Goroutines runs in the same address space, so access to shared memory must be synchronized. The `sync` pckg provides useful primitives, but you won't need them much as there are other exists.


## Channels

* What is a channel?

    * Channels are typed conduit through which you can send and receive values, using channel operator `<-`
    ##
```go
ch <- v // send v to channel ch
v := <-ch // Receive from ch, and assign value to v
// DATA FLOWS IN THE DIRECTION OF THE ARROW
```

* How to use channel?

    * They must be declared and initialized: `ch := make(chan int)`

* How do they work?

    * By default, they send and receive blocks until the *other side is ready*. This allows *goroutines* to synchronize without explicit locks or condition variables


```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c) // sum s[0] s[1] s[2]
	go sum(s[len(s)/2:], c) // sum s[3] s[4] s[5]
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

* Code above sums the number in slices, and sends them to the `c` channel in separated threads. After computation in `func sum` is completed, it calculates the final value in `fmt.Println`

## Buffered channels

* What is a buffered channel?

    * It describes how many channels can be opened, passed as second argument in `make(chan int, 100)`

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
    // ch <- 3
    // fmt.Println(<-ch) // it prints: fatal error: all goroutines are asleep - deadlock! why? because we opened more channels than specified in buffer length
}
```

## Closing channels

* Why we want to close channel?

    * A sender can close channel, why? To indicate that no more values will be sent.

* How to check it?

    * Receiver can test is channel closed, how? By assigning a second parameter to expression: after

```go
v, ok := <-ch // ok is false when there are no more values to receive and the channel is closed.
```
##
```go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // receives values from the channel repeatedly until its closed
		fmt.Println(i)
	}
}

```

* Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a *range* loop. 

* Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic. 


## Select

* What is a select?

    * It's like a `switch` statement but designed to work with goroutines

    * It creates a logic for multiple communication operations

* How *select* works?

    * It blocks until one of its cases can run, then it executes that case.
    
    * If multiple cases are ready, it choose one at random.


```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

* What if all cases are not ready?

    * In order to handle this, we can use `default` block.

    * `default` case in `select` runs if no other case is ready

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```
# Related to:

* https://go.dev/doc/



#golang #goroutines #concurrency