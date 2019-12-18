package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Printf("First example with -1")
	ret1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Println("error,return values are ", ret1, err1)
	} else {
		fmt.Printf("it's ok! return values are:", ret1, err1)
	}
	if ret2, err2 := MySqrt(5); err2 != nil {
		fmt.Println("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", ret2, err2)
	}
	// named return variables:
	fmt.Println(MySqrt2(5))
}
func MySqrt(f float64) (float64, error) {
	//return an error as second parameter if invalid input
	if f < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	//otherwise use default square root function
	return math.Sqrt(f), nil
}

//name the return variables - by default it will have 'zero-ed' values i.e. numbers are 0, string is empty, etc.
func MySqrt2(f float64) (ret float64, err error) {
	if f < 0 {
		//then you can use those variables in code
		ret = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number!")
	} else {
		ret = math.Sqrt(f)
		//err is not assigned, so it gets default value nil
	}
	//automatically return the named return variables ret and err
	return
}

// error_returnval.go
