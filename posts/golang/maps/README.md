# Maps

* What is a *map*?

    * Map assign to *keys* specified *values*



* What is initial (zero) value of *map*?

    * It's a `nil` value

* How to create a *map*?

    * Using func `make` or via `var`

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

### Map literals

* What are *map* literal?

    * *Map* literals are same as *struct* literals, but in *map* we need to provide *keys*

    * If type of *element* **returned by map** is same as *type name*, you can omit providing it. (commands in code)
```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	}, 
	"Google": Vertex{
		37.42202, -122.08408,
	},

    // "Bell Labs": {40.68433, -74.39967},
    // "Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```

### Map modification

* How to manipulate *map*?

    * *Inserting* or *Update* element in *map*
    ## 
    `m[key] = elem`

    * *Get* element
    ##
    `elem = m[key]`

    * *Delete* element
    ##
    `delete(m, key)`

    * *Check* is *map* have specified *key* via *assigning* result to 2 variables. If *map* have *key*, second variable have value `true` otherwise `false`. If *map* doesn't have *key*, then *elem* have 
    initial (zero) value (depends on *map element type*). **If `first variable` | `second variable` has not been declared, you can use *short form of declaration* `elem, ok := m[key]`**
    ##
    `elem, ok = m[key]`

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Odpowiedź"] = 42
	fmt.Println("Wartość:", m["Odpowiedź"])

	m["Odpowiedź"] = 48
	fmt.Println("Wartość:", m["Odpowiedź"])

	delete(m, "Odpowiedź")
	fmt.Println("Wartość:", m["Odpowiedź"])

	v, ok := m["Odpowiedź"]
	fmt.Println("Wartość:", v, "Obecna?", ok)
}
```

* https://go-tour-pl1.appspot.com/moretypes/19

# Related to: 

* https://go.dev/doc/


#golang #map 