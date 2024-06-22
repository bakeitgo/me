package main

type command struct {
	c rune
}

var Increment = command{'i'}
var Decrement = command{'d'}
var Square = command{'s'}
var Output = command{'o'}

func Parse(data string) (deadfish []int) {
	value := 0
	for _, input := range data {
		switch input {
		case Increment.c:
			value += 1
		case Decrement.c:
			value -= 1
		case Square.c:
			value *= value
		case Output.c:
			deadfish = append(deadfish, value)
		}
	}
	return deadfish
}

/*

Write a simple parser that will parse and run Deadfish.

Deadfish has 4 commands, each 1 character long:

i increments the value (initially 0)
d decrements the value
s squares the value
o outputs the value into the return array
Invalid characters should be ignored.

Parse("iiisdoso") == []int{8, 64}

*/
