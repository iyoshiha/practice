package main

import "fmt"

func main(){

	// map with string key and int values
	var nilMap map[string]int // nil map
	var totalWins = map[string]int{} // empty map

	fmt.Println("===== nilMap =====")
	fmt.Println(nilMap)
	fmt.Println("===== totalWins =====")
	fmt.Println(totalWins)
	fmt.Println("===== make =====")

	teams := map[string][]string{
		"Orcas": []string{"Fred", "Ralph"},
		"Lions": []string{"ired", "asalph"},
		"kittens": []string{"ilred", "opalph"},
	}
	
	fmt.Println(teams["Orcas"])
	ran := map[string]int{
		"mynum":4,
		"Komnum":9,
	}
	fmt.Println(ran["mynum"])
	
	

}