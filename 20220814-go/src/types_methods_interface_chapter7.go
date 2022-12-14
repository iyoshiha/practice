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

type Score int
type Converter func(string)Score
type TeamScores map[string]Score
type Person struct {
	FirstName string
	LastName string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d" , p.FirstName, p.LastName, p.Age)
}

func main() {
	p := Person {
		FirstName: "Fred",
		LastName: "Fredson",
		Age : 23,
	}

}