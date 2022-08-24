package main

import (
	"algorithms/internal/utils"
	"fmt"
)

/*
 A linked list is a linear collection of data elements whose order is not given by their physical placement in memory.
 Instead, each element points to the next.
 It is a data structure consisting of a collection of nodes which together represent a sequence
*/

func main() {
	fmt.Println("Linked List")
	list := utils.New()
	fmt.Println(list)
	list.PushBack("Element 1")
	p := list.PushBack("Element 2")
	list.PushBack("Element 3")
	list.Remove(p)
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
