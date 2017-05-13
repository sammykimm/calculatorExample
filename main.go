package main

import (
	"fmt"
	"time"
)

func main() {
	// var name = "star"
	// number  := 48
	fmt.Println("Hello",validate("hello", 5))
	fmt.Println("time", time.Now())
}

func validate(input string, number int) int{
	if input == "hello"{
		return 4 * number
	}
	return 0 * number
}

