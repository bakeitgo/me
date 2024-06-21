package main

// outside of this package field r would be encapsulated
// we will imitate this behaviour with getter
type vowels struct {
	r [5]rune
}

func (v vowels) IsVowel(r rune) (isVowel bool) {
	for _, vowel := range v.r {
		if r == vowel {
			isVowel = true
			break
		}
	}
	return
}

var (
	VOWELS = vowels{[5]rune{'a', 'e', 'i', 'o', 'u'}}
)

func GetCount(str string) (count int) {
	for _, r := range str {
		if VOWELS.IsVowel(r) {
			count++
		}
	}
	return
}

func main() {
}

/*

Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.


*/
