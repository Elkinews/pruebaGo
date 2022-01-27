package selfFuncs

import "strings"

//ReverseSentence , funcs that take an string and reorder all his characters , converting the lates character in the first
func ReverseSentence(sentence string) string {
	var textReverse string

	for i := len(sentence) - 1; i >= 0; i-- {
		textReverse += string(sentence[i])
	}
	return textReverse

}

// SupressSpaces , function that supress the space character of a string variable
func SupressSpaces(sentence string) string {
	var SupressTextSpace string
	for i := 0; i < len(sentence); i++ {
		if string(sentence[i]) != " " {
			SupressTextSpace += string(sentence[i])
		}
	}
	return SupressTextSpace
}

func PalindromeDetect(sentence string) bool {
	var sentenceToLower string = strings.ToLower(sentence)
	sentenceToLower = SupressSpaces(sentenceToLower)

	if ReverseSentence(sentenceToLower) != sentenceToLower {
		return false
	} else {
		return true
	}

}
func MultAB(a float32, b int) float32 {
	return a * float32(b)
}
