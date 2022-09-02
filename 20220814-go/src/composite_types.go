package main

import (
	"fmt"
)


func main(){
	
	samples := []string{"hello", "apple_n!"}
	outer:
		for ii, sample := range samples {
			for i, r := range sample {
				fmt.Println(ii, i, r, string(r))
				if r == 'l' {
					continue outer
				}
			}
			fmt.Println()

		}
}
