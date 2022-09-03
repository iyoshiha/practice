package main

import (
	"fmt"
	"os"
	"errors"
)

func main() {

	result, remainder, err := divAndRemainder(5,2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder, err)
}

// blank return (naked return )
// we must not use blank return since it's really confusing for code readers  
func divAndRemainder(numerator, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divie by zero")
		return 
	}
	result, remainder = numerator / denominator, numerator % denominator
	return 
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out,  base+v)
	}
	return out
}

func variadicInputParameters() {
}

// you can write your input when multiple input prameters of the same type
func div(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

