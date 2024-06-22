package main

func FindMissingLetter(chars []rune) (missing rune) {
	prevCharAscii := chars[0]
	for _, char := range chars[1:] {
		if prevCharAscii++; prevCharAscii != char {
			break
		}
	}
	return
}
