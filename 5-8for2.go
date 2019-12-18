package main

import "fmt"

func main() {
	var i int = 5
	for i >= 0 {
		i = i - 1
		fmt.Printf("the variable i is now %d\n", i)
	}
}
