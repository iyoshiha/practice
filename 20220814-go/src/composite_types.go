package main

import "fmt"

func main(){

	// struct 
	type person struct {
		name	string
		age		int
		pet		string
	}

	var fred person
	bob := person{"julia", 40, "cat"}

	fmt.Println("===== person =====")
	fmt.Println(fred)
	fmt.Println("===== totalWins =====")
	fmt.Println(bob)
	fmt.Println("===== make =====")

	beth := person{
		age: 30,
		name: "beth",
	}

	fmt.Println(beth)
	
	type firstPerson struct {
		name string
		age int
	}

	type secondPerson struct {
		name string
		age int
	}

	var g struct {
		name 	string
		age		int
	}

	f := firstPerson{"bob", 50}
	g = f
	fmt.Println(g)



}