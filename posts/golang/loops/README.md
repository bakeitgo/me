# Loops

* How many loops *Golang* have?

	* *Golang* have only one loop, *for* loop.

	* Basic *for* loop have 3 gradients separated by *;*

* How *for* loop works?

	1. Initialization - occurs **once**, before *loop* start e.g. (i := 0)
	1. Conditional expression - checks before **every** *loop* starte.g. (sum < 1000)
	1. Incrementation - occurs after **every** *loop* iteration e.g. (i++)

* Often *Initialization* is short variable / variables declaration
	
	* *Note* Those variables are accessible only inside loop scope

* When loop stop *iteration*?
	
	* When *bool* value will turn to *false* value

	```go
	package main

	import "fmt"

	func main() {
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		fmt.Println(sum)
	}
	```


* BeWARE! 
	* Besides *Java* / *JavaScript* / *C* we dont use `()` parentheses which surround *loop gradients*, but `{ }` curly braces needs to be applied.
2

### Optional gradients in *loop*

* Which gradients are optional?
	
	1. Initialization
	1. Incrementation

	```go
	package main

	import "fmt"

	func main() {
		sum := 1
		for ; sum < 1000; {
			sum += sum
		}
		fmt.Println(sum)
	}
	```
	
	* You can omit *;* while not using gradients listed above.
	
	```go
	package main

	import "fmt"

	func main() {
		sum := 1
		for sum < 1000 {
			sum += sum
		}
		fmt.Println(sum)
	}
	```

### Infinite loop
* If you omit *conditional expression* loop will be *infinite*
	
```go
package main

func main() {
	for {
	}
}
```





## Relates to: 

* https://go.dev/doc/


#golang #for #loop #iteration #infiniteloop 
