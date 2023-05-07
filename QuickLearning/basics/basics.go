package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"time"

	"example.com/types"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	//for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// optional init

	sum2 := 0
	for sum2 < 1000 {
		sum2 += sum
	}

	//infinite loop
	// for{

	// }

	fmt.Println(sum2)

	fmt.Println(sqrt(2), sqrt(-16), pow(3, 2, 10), pow(3, 3, 20))

	Runtime()
	Time()
	SwitchTrue()
	DeferCheck()
	DeferStackCheck()
	types.Pointer()
	fmt.Println(types.Vertex{1, 2})

	v := types.Vertex{1, 2}
	v.X = 5

	fmt.Println(v)
}

// If statement
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt((x)))
}

// if with short statement
// Variables declared by the statement are only in scope until the end of the if.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		//v can also be accessed in else scope
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// Compile error, v only in if scope
	// return v

	return lim
}

// exercise Newton's method
func SqrtNewton(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

// Switch statement
// break state need in other language is automatically provided
// switch cases need not to be constants

func Runtime() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

// Switch case evalute top down
func Time() {
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
}

// Switch without case is true

func SwitchTrue() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Defer statement
// defer statement defers the execution of a function until the surrounding function returns
// The defered call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns

func DeferCheck() {
	defer fmt.Println("world!")

	fmt.Print("Hello ")
	fmt.Print("to this ")
}

// Defer calls are pushed on to a stack, executed in last in first out order

func DeferStackCheck() {
	fmt.Println("Checking numbers:")

	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		defer fmt.Println(s)
	}
	//This will show before
	fmt.Println("End of numbers")
}
