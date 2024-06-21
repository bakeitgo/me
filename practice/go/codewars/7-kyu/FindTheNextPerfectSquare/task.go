package main

import "fmt"

var iterator int64 = 1

func RecursiveFindNextSquare(sq int64) int64 {
	isPerfectSqrt := iterator*iterator == sq
	if isPerfectSqrt {
		return (iterator + 1) * (iterator + 1)
	}
	if iterator*iterator > sq {
		return -1
	}
	iterator += 1
	return RecursiveFindNextSquare(sq)
}

func FindNextSquare(sq int64) int64 {
	var iterator int64 = 1
	for {
		sqrt := iterator * iterator
		if sqrt == sq {
			return (iterator + 1) * (iterator + 1)
		}
		if sqrt > sq {
			return -1
		}
		iterator++
	}
}

func main() {
	var sqrt int64 = 121
	fmt.Println(RecursiveFindNextSquare(sqrt))
}

/*

You might know some pretty large perfect squares. But what about the NEXT one?

Complete the findNextSquare method that finds the next integral perfect square after the one passed as a parameter. Recall that an integral perfect square is an integer n such that sqrt(n) is also an integer.

If the argument is itself not a perfect square then return either -1 or an empty value like None or null, depending on your language. You may assume the argument is non-negative.

Examples:(Input --> Output)

121 --> 144
625 --> 676
114 --> -1 since 114 is not a perfect square


*/
