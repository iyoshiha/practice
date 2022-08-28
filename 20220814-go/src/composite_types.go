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

	// set 
	fmt.Println("======== set ======== ") 

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals{
		fmt.Println(intSet[v])
		intSet[v] = true
		fmt.Println(v)
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

}