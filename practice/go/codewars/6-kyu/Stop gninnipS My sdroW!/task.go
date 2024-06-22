package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {
	aTokens := strings.Split(str, " ")
	for i, word := range aTokens {
		if len([]rune(word)) > 4 {
			aTokens[i] = reverse(word)
		}
	}
	return strings.Join(aTokens, " ")
} // SpinWords

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(SpinWords("Hey fellow warriors"))
}

/*

Write a function that takes in a string of one or more words, and returns the same string,
but with all words that have five or more letters reversed (Just like the name of this Kata).
Strings passed in will consist of only letters and spaces.
Spaces will be included only when more than one word is present.

Examples:

"Hey fellow warriors"  --> "Hey wollef sroirraw"
"This is a test        --> "This is a test"
"This is another test" --> "This is rehtona test"

*/
