package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Println(j, "*", i, "=", i*j)
	// 	}
	// }
	// x := 1
	// for ; x < 10; {
	// 	fmt.Println(x)
	// 	x++
	// }
	fmt.Println(runtime.GOOS)

	fmt.Println("When's Saturday?")
	fmt.Println(time.Now().Weekday())
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