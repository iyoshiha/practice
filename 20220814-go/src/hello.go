package main

import "fmt"

var global = 20

func main() { 
	m := 80
	const x int = 10
	const (
		idKey = "id"
		namekey = "name"
	)
	const z = 20 * 10

	fmt.Println(x)
	fmt.Println(m)
	fmt.Println(idKey)
	fmt.Println(namekey)
	fmt.Println(z)

}
