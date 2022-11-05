package main

import "fmt"

func isEven(number int) bool {
	return number%2 == 0
}

func main() {
	fmt.Println(isEven(64)) // true
	fmt.Println(isEven(43)) // false 
}
