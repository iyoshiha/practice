package main

import "fmt"

type MyfuncOpts struct {
	FirstName string
	LastName string
	Age int
}

func main(){

	fmt.Println(div(100, 0))
	fmt.Println(MyfuncOpts{FirstName:"ba",Age:10})



}

func myfunc(opts MyfuncOpts) int{

	return opts.Age 
}

func div(numerator int, denominator int) int {
	if denominator == 0{
		return 0
	}
	return numerator / denominator
}