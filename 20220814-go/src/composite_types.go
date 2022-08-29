package main

import "fmt"

func main(){

	// struct 
	type person struct {
		name	string
		age		int
		pet		string
	}
	var totalWins = map[string]int{} // empty map

	fmt.Println("===== person =====")
	fmt.Println(person)
	fmt.Println("===== totalWins =====")
	fmt.Println(totalWins)
	fmt.Println("===== make =====")
}