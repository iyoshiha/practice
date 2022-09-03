package main

import (
	"fmt"
)

func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2,4,5,6,8))
	a := []int{4,3}
	fmt.Println(addTo(3,a...))
	fmt.Println(addTo(3,[]int{1,2,3,4,5}...))

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

