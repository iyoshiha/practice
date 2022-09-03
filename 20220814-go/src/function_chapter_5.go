package main

import (
	"fmt"
)

func main() {
	result := div(5,2)
	fmt.Println(result)

}

// you can write your input when multiple input prameters of the same type
func div(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

