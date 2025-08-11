package main

import (
	"fmt"
)

func main() {
	numbers := []int {1, 2, 3, 4, 5, 6}
	posNumbers := []int {}
	for _, i := range numbers{
		if (i % 2 == 0){
			posNumbers = append(posNumbers, i)
		}
	}
	fmt.Println(numbers)
	fmt.Println(posNumbers)
}