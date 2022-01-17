package main

import (
	"fmt"
	"time"
	"sync"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i,s)
	}
}


func main() {	
	var wg sync.WaitGroup

	process := func(item string) {
	  fmt.Printf("processing %v \n", item)
	}
	items := []string{"item1", "item2", "item3", "item4", "item5"}
  
	wg.Add(len(items))
	for _, item := range items {
	  go func() {
		defer wg.Done()
		process(item)
	  }()
	}
  
	wg.Wait()
}
