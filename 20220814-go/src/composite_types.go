package main

import "fmt"

func main(){

	// make

	// x := make([]int,5)

	// x = append(x, 8) // [0 0 0 0 0 0 0 0 10]
	// fmt.Println(x)
	// fmt.Println(cap(x))

	x := []int{1,2,3,4}
	fmt.Println(cap(x))
	y := x[:2]
	y = append(y,10)
	// x = append(x,20)
	y = append(y,10,10)
	y = append(y,10)
	
	var capa = [5]int{}

	fmt.Println("===== x =====")
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("===== y =====")
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
	fmt.Println("===== make =====")
	copy(capa[0:], x)
	fmt.Println((capa))
	
}