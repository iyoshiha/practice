package main

import "fmt"

func main(){

	// make

	// x := make([]int,5)

	// x = append(x, 8) // [0 0 0 0 0 0 0 0 10]
	// fmt.Println(x)
	// fmt.Println(cap(x))

	y := make([]int, 5, 10)
	fmt.Println(y)
	fmt.Println(cap(y))

	p := make([]int, 0, 10)
	fmt.Println(p == nil) // false
	var q []int;
	fmt.Println(q == nil) // true
	var j []int = []int{};
	fmt.Println(j == nil) // true

	
	
}