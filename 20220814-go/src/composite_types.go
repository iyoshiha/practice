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
	ran["majimanji"] = 1010
	fmt.Println(len(ran))
	
	// var emp map[string]int
	// emp["err"] = 3 // panic assignment to nil
	// fmt.Println(emp)

	
	ages := make(map[int][]string, 10)
	ages[1] = []string{"kondo"}
	
	fmt.Println((ages[8]))
	fmt.Println((ages))

	// ok idiom 
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println("ok idiom\n", v, ok)	
	v, ok = m["world"]
	fmt.Println("ok idiom\n", v, ok)	
	v, ok = m["nothing"]
	fmt.Println("ok idiom\n", v, ok)	
	delete(m, "hello")
	fmt.Println("delete function\n", m)	


}