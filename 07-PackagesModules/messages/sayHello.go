package messages

import "fmt"

func SayHello() { // With uppercase first letter, this function is exported and public
	fmt.Println("Hello from the Messages package")
	onlyThisPackage()
}

const message = "Hello from the constant"

func onlyThisPackage() { // With lowercase first letter, this function is not exported and private
	fmt.Println(message)
}
