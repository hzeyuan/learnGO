package main

import "fmt"

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

func main() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Printf("FizzBuzz")

		case i%FIZZ == 0:
			fmt.Printf("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
