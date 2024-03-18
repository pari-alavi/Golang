package main

import (
	"fmt"
	"strings"
)

func Reversing(input string) string {

	reverseUnit := func(word string) string {
		// converting the word into a slice of bytes
		letters := []byte(word)
		for i := 0; i < len(letters)/2; i++ {
			j := len(letters) - 1 - i
			letters[i], letters[j] = letters[j], letters[i]
		}
		//converting the slice of bytes back into a string
		return string(letters)
	}

	//spliting the input of string into words
	words := strings.Fields(input)
	// iterating over the words
	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverseUnit(word)
		}
	}
	//joining the words back into a string
	return strings.Join(words, " ")
}

func main() {
	input := "This is a sample string with words of different lengths number 85962158"
	result := Reversing(input)
	fmt.Println("Original:", input)
	fmt.Println("Reversed:", result)
}
