package main

import (
	"fmt"
	"strings"
)

func High(s string) (result string) {
	var highest rune = 0
	words := strings.Split(s, " ")
	for _, word := range words {
		var wordResult rune = 0
		for _, b := range word {
			scoreFromAscii := b - 97 + 1
			wordResult += scoreFromAscii
			if wordResult > highest {
				highest = wordResult
				result = word
				continue
			}
		}
	}
	return
}

func main() {
	fmt.Println(High("aba ccc aaa"))
}
