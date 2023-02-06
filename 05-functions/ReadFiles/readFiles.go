package main

import (
	"fmt"
	"os"
)

func main() {

	// Function recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Error: ", error)
		}
	}()
	// evalute error
	if file, error := os.Open("test.txt"); error != nil {
		// print error
		panic(error) // this panic function stops the execution of the program

	} else {
		defer func() { // a necessary defer function to close the file
			fmt.Println("Closing file")
			file.Close()
		}()

		content := make([]byte, 254)              // create a byte array to store the content of the file
		long, _ := file.Read((content))           // read the file and store the content in the byte array
		contentToString := string(content[:long]) // convert the byte array to string

		fmt.Println(contentToString)
	}

}
