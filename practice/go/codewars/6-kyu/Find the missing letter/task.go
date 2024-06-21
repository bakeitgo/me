package main

func FindMissingLetter(chars []rune) (missing rune) {
	var prevCharAscii byte
	for _, char := range chars {
		if prevCharAscii != 0 && byte(char)-prevCharAscii > 1 {
			missing = rune(char - 1)
			break
		}
		prevCharAscii = byte(char)
	}
	return
}
