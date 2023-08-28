package main

import (
	"fmt"
	"os"
)

func main() {
	const (
		introductionText = "Добрый день. Что бы задать размеры шахматной доски, введи целое положительно не равное 0 число.\n"
		errorText        = "Введены не верные данные, прошу повторить ввод.\n"
		finalText        = "Благодарю тебя мой Дорого Пользователь! Заглядывай ещё! Я буду ждать тебя!"
	)
	var (
		hashtag   = "[#]"
		spacex    = "[ ]"
		sizeBoard int
	)
	for {
		fmt.Printf("%s", introductionText)
		fmt.Fscan(os.Stdin, &sizeBoard)
		if sizeBoard <= 0 {
			fmt.Printf("%s", errorText)
			continue
		}
		break
	}
	for j := 1; j <= sizeBoard; j++ {
		for i := 1; i <= sizeBoard; i++ {
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
	fmt.Printf("%s", finalText)
}
