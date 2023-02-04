package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ { // the same as other languages (i++ is the same as i = i + 1)
		fmt.Println(i)
	}

	// for loop with break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break // the break statement terminates the loop
		}
		fmt.Println(i)
	}

	// for loop with continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue // the continue statement skips the current iteration of the loop
		}
		fmt.Println(i)
	}

	// for loop with continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			goto end // the goto statement transfers control to the labeled statement
		}
		fmt.Println(i)
	}
end:
	fmt.Println("done")

	// for as while
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// for each
	s := []string{"a", "b", "c"}
	for index, value := range s {
		fmt.Println("index:", index, "value:", value)
	}
}
