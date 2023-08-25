package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		hashtag     = "[#]"
		spacex      = "[ ]"
		lenghtBoard int
		heightBoard int
	)
input_data:
	fmt.Println("Введите целое положительное число больше 0 определяющее длинну доски:")
	fmt.Fscan(os.Stdin, &lenghtBoard)
	fmt.Println("Введите целое положительное число больше 0 определяющее высоту доски:")
	fmt.Fscan(os.Stdin, &heightBoard)
	switch {
	case lenghtBoard <= 0 || heightBoard <= 0:
		fmt.Printf("%s", "Ширина или высота не может быть равна или меньше 0, укажите корректные размеры.\n")
		goto input_data
	}

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
