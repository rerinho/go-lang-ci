package main

import "fmt"

func main() {
	fmt.Println(sum(10, 10))
}

// It returns the sum of 2 integers
func sum(a int, b int) int {
	return a + b
}
