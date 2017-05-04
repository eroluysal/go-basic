package main

import "fmt"

func main() {
	sum := 1

	// Semicolon is required as much as the number of arguments
	for ; sum < 1000; {
		sum += sum
	}

	fmt.Println(sum)
}
