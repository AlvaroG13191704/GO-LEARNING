package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sum(expresion string) int {
	// 4+4+5+6+7+8+9+10
	var result int
	stringArray := strings.Split(expresion, "+")

	// sum the values
	for _, value := range stringArray {
		// convert value to float64
		num, err := strconv.Atoi(value) // this func return two values (the value and the error)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error converting string to int")
			return 0
		}
		result += num
	}

	return result
}

func main() {
	var expresion string
	// var result flaot64

	fmt.Print("=>")
	fmt.Scan(&expresion)

	fmt.Println(sum(expresion))

}
