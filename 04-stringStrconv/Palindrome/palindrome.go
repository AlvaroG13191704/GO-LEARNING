package main

import (
	"fmt"
	"strings"
)

func reverse(word string) string {
	// convert to array
	aray := strings.Split(word, "")
	// reverse the array
	for i, j := 0, len(aray)-1; i < j; i, j = i+1, j-1 {
		aray[i], aray[j] = aray[j], aray[i]
	}
	// convert to string
	return strings.Join(aray, " ")
}

func isPalindrome(word string) {
	// Convert to lowercase and uppercase
	word = strings.ToLower(word)
	word = strings.ToUpper(word)
	fmt.Println(word)

	// Replace characters
	newWord := strings.ReplaceAll(word, " ", "") //with replace we defined the number of characters to be replaced
	fmt.Println(newWord)
}

func main() {

	// call the func
	isPalindrome("Deleveled")
	reverse("Deleveled")
	fmt.Print(reverse("Hola"))
}
