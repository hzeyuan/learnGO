package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is :%d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is : %c \n", ix, str[ix])
	}
	str2 := "日语"
	fmt.Printf("the length of str is : %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c", ix, str2[ix])
	}
}
