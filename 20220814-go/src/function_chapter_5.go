package main

import (
	"fmt"
	"strconv"
	"errors"
	"os"
	"log"
	"io"
	"context"
	"database/sql"
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

type Person_ struct {
	FirstName string
	LastName string
	Age int
}

func makeMult(base int) func(int) int{
	return func(factor int) int {
		return base * factor
	}
}

func deferExample() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
		}
	}
}

type person struct {
	age int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func main5() {
	p := person{}
	i := 2
	s := "Hello"
	gosa
	modifyFails(i, s, p)
	

}

func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
		tx,err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		defer func() {
			if err == nil {
				err = tx.Commit()
			}
			if err != nil {
				tx.Rollback()
			}
		}()
		_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1")
		if err != nil {
			return err
		}
		// use tx to do more database inserts here
		return nil
	}

func orderDefer() int {

	defer fmt.Println("end of orderDefer")
	return 10

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

