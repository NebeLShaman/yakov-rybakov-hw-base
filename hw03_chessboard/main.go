package main

import (
	"fmt"
)

func main() {
	var (
		hashtag     string = "[#]"
		spacex      string = "[ ]"
		lenghtBoard        = 8
		heightBoard        = 8
	)

	for j := 1; j <= heightBoard; j++ {
		for i := 1; i <= lenghtBoard; i++ {
			if i%2 == 0 && j%2 == 0 {
				fmt.Printf("%s", hashtag)
			} else if i%2 != 0 && j%2 != 0 {
				fmt.Printf("%s", hashtag)
			} else {
				fmt.Printf("%s", spacex)
			}
			if i == lenghtBoard {
				fmt.Printf("%s", "\n")
			}
		}
	}
}
