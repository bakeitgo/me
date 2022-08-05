# Pointers

* What is a pointer?
	
	* Pointer points on memory address where some value is stored.
	
* What about pointer usage?
	
	* Pointer is useful when you want to assign one variable to another variable. This lets you  manipulate one variable using any variable which points to the same memory address.

* How to use pointer?

	* `*T` is a pointer which points to *T* type. If you consume it like this: `var p *int`, then it's initial value (*zero value*) = `nil` 

* How to point to another variable? 
	
	* `i := 42
	   p = &i`
	
* How to consume pointer variable?
	
	* By using `*` operand. `fmt.Println(*p)

	* Assigning new value to some variable using pointer. `\*p = 21` (Now `i` value is 21) It's called *dereferencing*


# Related to: 

* https://go.dev/doc/


#golang #pointer #memoryaddress #variable
