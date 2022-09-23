package main

import(
	. "fmt"
)

func main() {
	// true 
	var x *int
	Println(x == nil)

	// false
	x = new(int)
	Println(x == nil)

	// false
	var y int
	x = &y
	Println(x == nil)
}
