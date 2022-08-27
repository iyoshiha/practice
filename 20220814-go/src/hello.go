package main

import "fmt"

var global = 20

func main() { 
	x := 10
	x, y := 20, "hello"

	fmt.Println(y)
	fmt.Println(x)
}
