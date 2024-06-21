package main

import (
	"fmt"
	"strings"
)

func reverse(str string) (reversed string) {
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return
}

func SpinWords(str string) string {
	var spinnedWords []string
	words := strings.Split(str, " ")
	for _, word := range words {
		if len(word) >= 5 {
			word = reverse(word)
		}
		spinnedWords = append(spinnedWords, word)
	}

	return strings.Join(spinnedWords, " ")
} // SpinWords

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
