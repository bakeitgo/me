# Structs

* Struct is a field collection

	* It's kinda like object in JavaScript

* In order to access field from struct, we use `.`

* `type` keyword creates a new type with name you specify
```go
package main

import "fmt"

type Vertex struct { // here type is Vertex, it can be used similar to string,int etc. e.g. in function parameters
	X int
	Y int
} // This is how we declare struct

func main() {
	v := Vertex{1, 2} // consuming struct
	v.X = 4 // accessing field
	fmt.Println(v.X)
}
```

### Pointers to structs:

* To access struct field e.g. `X` via pointer `p` we can write (\*p).X. But this notation is not handy. *Golang* let us refer to `X` field via this notation: `p.X`. We are omitting pointer dereference

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

### Struct literals
	
* *Literal struct* means new created value which contains struct by exchangeable values of struct fields

* You can assign values to chosen struct fields, by using *Name:* notation. ** The order of fields in struct doesn't  matter **

* Prefix **&** returns pointer to struct

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // posiada typ Vertex
	v2 = Vertex{X: 1}  // Y:0 jest domniemane
	v3 = Vertex{}      // X:0 oraz Y:0
	p  = &Vertex{1, 2} // posiada typ *Vertex
)

func main() {
	v1 = Vertex{10,10}
	fmt.Println(v1, p, v2, v3)
}
```

# Related to: 

* https://go.dev/doc/


#golang #struct #fields #accessingstruct #structpointer #structliterals


