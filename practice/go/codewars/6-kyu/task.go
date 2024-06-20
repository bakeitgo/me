package main

// Two Sum
func TwoSum(numbers []int, target int) [2]int {
	if len(numbers) < 2 {
		return [2]int{}
	}
	for firstIndex, firstNumber := range numbers {
		for secondIndex, secondNumber := range numbers {
			if firstIndex == secondIndex {
				continue
			}
			sum := firstNumber + secondNumber
			if sum == target {
				return [2]int{firstIndex, secondIndex}
			}
		}
	}
	return [2]int{}
}

func main() {

}

/* NOTE:

The kata there is a sentence: "..It should find two different items in the array"

The point there is about finding different values in context of array (different indices of values, not values itself)

Description should be fixed to clarify it

*/

/*
Write a function that takes an array of numbers (integers for the tests) and a target number.
It should find two different items in the array that, when added together, give the target value. The indices of these items should then be returned in a tuple / list (depending on your language) like so: (index1, index2).

For the purposes of this kata, some tests may have multiple answers; any valid solutions will be accepted.

The input will always be valid (numbers will be an array of length 2 or greater, and all of the items will be numbers;
target will always be the sum of two different items from that array).

*/

// NOTE: There is a bug in this kata,
