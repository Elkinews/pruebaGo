package selfFuncs

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
