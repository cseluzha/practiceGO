package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func sqrt(x float64) string {
	// if no necessary the () in the IF statement, but the braces {} are required
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//You can define  a value into statement, it's a short version
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	// Call to function
	fmt.Println(add(42, 13))

	// Statement FOR
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// The init and post statements are optional.

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	sum2 := 0
	for i := 0; i < 10; i++ {
		sum2 += add(i, sum2)
	}
	fmt.Println(sum2)

	// Without conditions is a forever For statement
	/*
		for {
		}
	*/

	// Statement IF
	fmt.Println(sqrt(2), sqrt(-4))
	// call to If with a short statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	// call to If and else
	fmt.Println(
		pow1(3, 2, 10),
		pow1(3, 3, 20),
	)

	// Statement SWITCH
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	//Switch evaluation order
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//Switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// Statement DEFER
	/*
		A defer statement defers the execution of a function until the surrounding function returns.

		The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.		
	*/

	//Use of defer

	defer fmt.Println("world")

	fmt.Println("hello")

	/*
	Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	*/
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
