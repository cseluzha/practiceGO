package main

import (
	"fmt"
)

func main() {
	Numbers := []int{2, 3, 5, 7, 8, 9, 12, 22, 50}
	fmt.Println(linearSearch(Numbers, 15))
}

// The linear search iterate through each element of the array
func linearSearch(numbers []int, item int) int {
	for i, v := range numbers {
		if v == item {
			return i
		}
	}
	return -1
}
