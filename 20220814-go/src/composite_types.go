package main

import (
	"fmt"
	"math/rand"
)


func main(){
	
	choosingIfAndSwitch()
}

func choosingIfAndSwitch() {
n := rand.Intn(10);
	fmt.Println("n =", n)
	switch  {
	case n == 0:
		fmt.Println("that's too low")
	case n > 5:
		fmt.Println("that's too big", n)
	default:
		fmt.Println("that's a good number", n)
	}
}

func blankSwitches() {

	words := []string{"hi","salutations", "hello"}
	for _,word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is the exqxtly the right word!")
		}
	}

}

func standardSwitch() {

	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologlist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1,2,3,4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6,7,8,9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}


func missingLabelSwitch() {

	loop:
		for i := 0; i < 10; i++ {
			switch {
			case i%2 == 0:
				fmt.Println(i, "is even")
			case i%3 == 0:
				fmt.Println(i, "is divisible by 3 but not 2")
			case i%7 == 0:
				fmt.Println("exit the loop")
				break loop
			default:
				fmt.Println(i, "is boring")
			}
		}
}