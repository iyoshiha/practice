package main

import "fmt"

func main(){

	var x string = "hello world!"
	var b byte = x[6]
	var y = []int{1,2,5:10,3,4,18:100}

	fmt.Println("===== x =====")
	fmt.Println(string(b))
	fmt.Println("===== y =====")
	fmt.Println(y)
	fmt.Println("===== make =====")
	
	var l rune = 'あ'
	fmt.Println(l)
	var n int = 65
	var nn = string(n) // int to string is not hobidden but not recommended
	fmt.Println(nn)

	var nnn string = "hello, ☺ A"
	var bs []byte = []byte(nnn)
	var rs []rune= []rune(nnn)
	fmt.Println(nnn)
	fmt.Println(bs)
	fmt.Println(rs)

}