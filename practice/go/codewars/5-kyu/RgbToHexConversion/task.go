package main

import "fmt"

/*
The rgb function is incomplete.
Complete it so that passing in RGB decimal values will result in a hexadecimal representation being returned.
Valid decimal values for RGB are 0 - 255.
Any values that fall out of that range must be rounded to the closest valid value.

Note: Your answer should always be 6 characters long, the shorthand with 3 will not work here.
*/

func valid(x int) int {
	switch {
	case x < 0:
		return 0
	case x > 255:
		return 255
	default:
		return x
	}
}

func main() {
	fmt.Printf("%02X%02X%02X", valid(18), valid(255), valid(255))
}
