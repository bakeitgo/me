package main

import (
	"fmt"
)

/*
ISBN-10 identifiers are ten digits long. The first nine characters are digits 0-9. The last digit can be 0-9 or X, to indicate a value of 10.

An ISBN-10 number is valid if the sum of the digits multiplied by their position modulo 11 equals zero.
*/

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}
	var res int
	for i, r := range isbn {
		if r == 88 && i == len(isbn)-1 {
			res += (i + 1) * 10
			continue
		}
		r = r - 48
		if r < 0 || r > 9 {
			return false
		}
		fmt.Printf("%v", r)
		res += (i + 1) * int(r)
	}
	return res%11 == 0
}

func main() {
	fmt.Println(ValidISBN10("048665088X"))
}
