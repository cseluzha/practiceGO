package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(7))
}

func factorial( n int ) int {
	if n >1{
		return n * factorial(n-1)
	} else {
		return 1
	}
}