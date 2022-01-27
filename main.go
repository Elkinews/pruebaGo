package main

import (
	"fmt"

	"example.com/prueba1/selfFuncs"
)

func main() {
	sentence := "ajitraga la lagartija"
	reversedText := selfFuncs.ReverseSentence(sentence)
	fmt.Println("hello world", reversedText)
	fmt.Println(selfFuncs.SupressSpaces(reversedText))

	if selfFuncs.PalindromeDetect(sentence) {
		fmt.Println("is palindrome")
	}
}
