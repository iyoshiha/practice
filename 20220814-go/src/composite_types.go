package main

import "fmt"

func main(){

	// var x [3]int
	var y = [3]int{1,2,3}
	var z = [...]int{1,2,3}
	var k [2][3]int 
	k[0] = y
	k[1] = z
	var x = [21]int{1,5:4,6,1:100,15}

	var s []int;
	var ss = []int{2,3,4,5}
	// ture
	fmt.Println(y == z);
	fmt.Println(len(k[0]));
	fmt.Println(x)
	s = append(s,4)
	s = append(s,ss...)
	fmt.Println(cap(s))
	s = append(s,4)
	fmt.Println(cap(s))
	
}