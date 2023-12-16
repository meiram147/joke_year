package main

import (
	"fmt"
	"my_project/example"
)

func main() {
	var (
		year int
		text string
	)
	text = "Введите год своего рождения:"
	fmt.Print(text)
	fmt.Scanln(&year)
	example.YearSet(year)
	fmt.Println(example.GetSay())
}
