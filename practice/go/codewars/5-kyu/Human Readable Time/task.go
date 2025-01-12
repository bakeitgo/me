package main

import (
	"fmt"
)

/*
Write a function, which takes a non-negative integer (seconds) as input and returns the time in a human-readable format (HH:MM:SS)

HH = hours, padded to 2 digits, range: 00 - 99
MM = minutes, padded to 2 digits, range: 00 - 59
SS = seconds, padded to 2 digits, range: 00 - 59
The maximum time never exceeds 359999 (99:59:59)

*/

func HumanReadableTime(seconds int) string {
	if seconds > 360_000 {
		return "99:59:59"
	}
	hours := seconds / 3600
	minutes := seconds % 3600 / 60
	seconds = seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func main() {
	fmt.Print(HumanReadableTime(359_900))
}
