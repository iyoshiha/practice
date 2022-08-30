package main

import (
	"fmt"
)


func main(){
	
	// nested code sucks  

// 	for i := 1; i <= 100; i++ {
// if i%3 == 0 {
// if i%5 == 0 {
// fmt.Println("FizzBuzz")
// } else {
// fmt.Println("Fizz")
// }
// } else if i%5 == 0 {
// fmt.Println("Buzz")
// } else {
// fmt.Println(i)
// }
// }


	// short is encourages in go 

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	// for - range 

	evenVals := []int{2,4,6,8,10,12}
	for _, v:= range evenVals {
		fmt.Println(v)
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k,v := range m {
			fmt.Println(k,v)
		}
	}

	fmt.Println("===========")

	samples := []string{"hello", "apple_n!"}
	for p, sample := range samples {
		for i, r := range sample {
			fmt.Println(p, " loop\n", i, r, string(r))
		}
		fmt.Println()
	}


}
