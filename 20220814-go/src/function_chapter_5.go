package main

import (
	"fmt"
)

func main() {
	var maji int = 100
	abc(maji)

}

type aaa int

func abc(a aaa) {
	fmt.Println(a)

}

// you can write your input when multiple input prameters of the same type
func div(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

