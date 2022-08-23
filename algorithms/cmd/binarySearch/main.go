package main

import (
	"fmt"
)

func main() {
	elements := []int{2, 3, 5, 7, 8, 9, 12, 22}
	fmt.Println(binarySearch(elements, 9))
}

// The binary search It is an efficient algorithm to find an element in an ordered list of elements.
func binarySearch(elements []int, item int) int {
	low := 0
	high := len(elements)  -1
	fmt.Println("Item: ", item)
	fmt.Println("Elements: ", elements)

	for low <= high {
		mid := (low + high) / 2
		guess := elements[mid]
		fmt.Printf("Low: %d High: %d mid: %d \n", low, high, mid)
		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
