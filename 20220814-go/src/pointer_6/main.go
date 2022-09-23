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
	Println(x)
	Println(*x)

	// false
	// you cant use an & before a primitive time 
	// when you need a primitive type pointer
	// declare a variable and point to it
	var y int
	x = &y
	Println(x == nil)
}
