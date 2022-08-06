# Arrays


*  What is an array?

	* It's a table which contains values with specified *type*

*  How to declare array?

	* *Type* `[n]T` is array which contains *n* values with *T* type

	* e.g ` var a [10]int` - declaring array which have *10* `integers`

	* *Note:* Array length *[10]* is **part of type** which means it cannot be changed. It seems to be *uncomfortable*, but *Golang* have a lot of ways to *deal with it* so don't worry.

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

### Slices

* What are *slices*?
	
	* Arrays have specified length right? So *slices* are elastic in use *view* of some *array elements* which *length* we can *dynamically change* In practice, *slices* are more *often* used than *arrays*

* How we read *slice*?
	
	* Type *[]T* is *slice* of elements with type *T*. The *[]T* we read as "It's an array *slice* with values of type *T*"

* How we create *slice*?

	* *Slice* is created by choosing *two indexes* **low** && **high** separated by colon **:**. It looks like this: `a[low : high]`

	* Expression above `a[low : high]` determine *range* of *slice*.*Slice* contains elements from *bottom boundary* to *top boundary* **(excluding last element)**

	* e.g. We have array `primes := [6]int{2, 3, 5, 7, 11, 13}` and we create *slice* `var s []int = primes[0:3]`. *Slice* now have values = [2, 3, 5] **Remember accessing array elements begins from 0 index** so *slice* have *0, 1, 2* index-based elements from *array*

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:3]
	fmt.Println(s)
}
```


* What *slices* really are?
	
	* *Slices* are **references** to *arrays*, it doesn't collect any data, it refer to selected piece of *array* 

	* In case it is *reference*, changing value of *slice* element, changing *array* element also.

	* *Slices* which points to the same *array*, gonna see changes also.

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

### Slice literals

* *Literal slice* looks like *Literal array* without specified *length*

	* *literal array* looks like: `[3]bool{true, true, false}

	* *literal slice* looks like: `[]bool{true, true, false}

```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

### Slice defaults

* When creating *slices* we can omit **bottom** and **top** *boundary* and *slice* default values will be used.

	* Default value for **bottom** *boundary* is `0`

	* Default value for **top** *boundary* is `array length`

* For array `var a [10]int`

	* Those *slices* are the **same**
	1. a[0:10]
	1. a[:10]
	1. a[0:]
	1. a[:]

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
```

### Slices length and capability

* What is *slice* length and capability?
	* *Length* is a number of elements *slice* point

	* *Capability* is an overall length of *slice* **counting from first *slice* element**

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// cut slice to give it zero length
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Expand slice length
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Delete from slice first 2 values
	s = s[2:]
	printSlice(s) // len=4 cap=4 [5 7 11 13]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

### Nil slices

* What is *Nil slice* initial (zero) value?

	* *Initial* **zero** value of *Nil slice* is *nil*

	* *Nil slice* *length* and *capability* equals **0** and don't point to any *array*

```go
package main

import "fmt"

func main() {
	var s []int // Declare nil slice
	fmt.Println(s, len(s), cap(s))
	if s == nil { // check content
		fmt.Println("nil!")
	}
}
```

### Creating *slice* using `make`

* How to create *array* with *dynamic length*?

	* Using built-in function `make`

* What `make` function creates?

	* Function `make` creates *array* which *every element* **initial** value is **zero** and returns *slice* which **points** to this *array*
	##
	`a := make([]int, 5)  // len(a)=5`

	* To define *capability* we need to provide *Third argument* to the `make` func
	## 
	`b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5`
	##
	`b = b[:cap(b)] // len(b) = 5, cap(b) = 5`
	##
	`b = b[1:]      // len(b) = 4, cap(b) = 4`

```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```

### *Slices* of *Slices*

* What types *Slices* can contain?
	
	*  *Slices* can contain any type, including **other slices**

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

### Appending to *slice*

* How to add new element to a *slice*?

	* *Golang* have built-in function: 
	##
	`func append(s []T, vs ...T) []T`

	* First param is *slice* with type *T*, second param are values with type *T* which we want to **append** to a *slice*

	* Result of **append** func is *slice* and all *elements* added

	* *Note:* If *array* to which *slice* points is too short, a **new, longer** *array* will be created. Returned *slice* will point to newly created *array*

```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// Slice grows if this is needed
	s = append(s, 1)
	printSlice(s)

	// We can add more elements than 1
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

### For loop with *slices*

* How to use *slice* in **for** loop?

	* The keyword *range* let us iterate on *slice* or *map*

	* When iterating on *slice* on **every step** 2 values are returned. First is *element index*. Second is *element copy value*

	* You can *omit* **index** or **element copy value** by assigning **_** to them. If you want to get only *index* you can *omit* **element copy value variable**
	##
	`for i, _ := range pow // omit value`
	##
	`for _, value := range pow // omit index`
	## 
	`for i := range pow // omit value`
```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```


# Related to: 

* https://go.dev/doc/


#golang #arrays #slices
