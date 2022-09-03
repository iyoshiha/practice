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

// multiple return value
func divAndRemainder(numerator, denominator int) (result int, remainder int, err error) {
	result, remainder = 20, 30
	if denominator == 0 {
		err = errors.New("cannot divie by zero")
		return result, remainder, err
	}
	// we dont have to return named return valued: 
	// this potentially has a problem
	return numerator / denominator, numerator % denominator, err
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

