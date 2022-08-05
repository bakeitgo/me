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

https://go-tour-pl1.appspot.com/moretypes/11




# Related to: 

* https://go.dev/doc/


#golang #arrays
