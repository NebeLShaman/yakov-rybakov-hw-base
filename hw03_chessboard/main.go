package main

import (
	"fmt"
)

func main() {
	var (
		hashtag     = "[#]"
		spacex      = "[ ]"
		lenghtBoard = 8
		heightBoard = 8
	)

	for j := 1; j <= heightBoard; j++ {
		for i := 1; i <= lenghtBoard; i++ {
			switch {
			case i%2 == 0 && j%2 == 0:
				fmt.Printf("%s", hashtag)
			case i%2 == 1 && j%2 == 1:
				fmt.Printf("%s", hashtag)
			default:
				fmt.Printf("%s", spacex)
			}
		}
		fmt.Printf("%s", "\n")
	}
}
