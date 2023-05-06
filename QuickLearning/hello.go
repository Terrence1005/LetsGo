package main

import (
	"fmt"
	"math"

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
