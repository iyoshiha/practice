package main

import (
	"fmt"
	"strconv"
	"errors"
)

type opFuncType func(int, int) int

func funcvaricheck() {
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"5"},
	}
	for i, expression := range expressions {
		fmt.Println("start :", i)
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		onFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := onFunc(p1, p2)
		fmt.Println(result)
	}
}


func main() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 10},
	}
	sort.Slice(people, func(i, j int) bool{
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool{
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

}
func add(i int, j int) int {return i + j}
func sub(i int, j int) int {return i - j}
func mul(i int, j int) int {return i * j}
func div(i int, j int) int {return i / j}

// var opMap = map[string]func(int, int) int{
var opMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

// blank return (naked return )
// we must not use blank return since it's really confusing for code readers  
func divAndRemainder(numerator, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divie by zero")
		return 
	}
	result, remainder = numerator / denominator, numerator % denominator
	return 
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out,  base+v)
	}
	return out
}

func variadicInputParameters() {
}

// you can write your input when multiple input prameters of the same type
func divk(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

