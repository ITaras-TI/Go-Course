package main

import (
	"fmt"
)

// problem 3
func main() {

	x := 7

	if x%7 == 0 {
		fmt.Println(x * 2)
	} else {
		fmt.Print("oops!")
	}

}
