package main

import (
	"fmt"
)

func fibonacchi(nx int){
	n := 0
	n1 := 1
	var sum int
	
	for i:=0; i < nx; i++ {
		sum = n + n1
		n = n1
		n1 = sum
		fmt.Println(sum)
	}
}
func main() {
	fibonacchi(20)
}