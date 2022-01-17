package main

import (
	"fmt"
	"time"
)

func process(c chan bool) {
	fmt.Printf("received: %v\n", <-c)
}

func checkchannel() {
	rc := make(chan bool)
	defer close(rc)
	go process(rc)
	rc <- true
	fmt.Print("done...")
}

func readvalues() {
	c := make(chan int)
	defer close(c)

	go func(c chan int) {
		for v := range c {
			fmt.Println(v)
		}
	}(c)

	for i := 0; i < 5; i++ {
		c <- i
	}
}

func bufferedchannel() {
	ch := make(chan int, 5)
	defer close(ch)

	go func(c chan int) {
		for v := range c {
			fmt.Println(v)
		}
	}(ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func process1(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "hello1"
	}
}
func process2(ch chan string) {
	for {
		time.Sleep(2 * time.Second)
		ch <- "hello2"
	}

}

func testforselect() {
	output1 := make(chan string)
	output2 := make(chan string)
	go process1(output1)
	go process2(output2)
loop:
	for {
		select {
		case s1, ok := <-output1:
			if !ok {
				break loop
			}
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
		}
	}
}

func main() {
	//checkchannel()
	//readvalues()
	//bufferedchannel()
	testforselect();
}
