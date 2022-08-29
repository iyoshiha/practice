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
	bob := person{}

	fmt.Println("===== person =====")
	fmt.Println(fred)
	fmt.Println("===== totalWins =====")
	fmt.Println(bob)
	fmt.Println("===== make =====")
}