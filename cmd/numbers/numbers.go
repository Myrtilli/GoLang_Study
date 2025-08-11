package main

import (
	"fmt"
)

func main() {
	numbers := []int {1, 2, 3, 4, 5, 6}
	pos_numbers := make([]int, 0, 10)
	for _, i := range numbers{
		if (i % 2 == 0){
			pos_numbers = append(pos_numbers, i)
		}
	}
	fmt.Println(numbers)
	fmt.Println(pos_numbers)
}