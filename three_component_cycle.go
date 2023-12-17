package main

import (
	"fmt"
)

func main() {
	v := 0
	for i := 1; i < 10; i++ {
		v++
	}
	fmt.Println(v)
	array := [3]int{1, 2, 3}
	for arrayIndex, arrayValue := range array {
		fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
	}
	for {
		var name string
		fmt.Println("hello whats are your name?")
		fmt.Scanln(&name)
		if name == "Meiram" {
			fmt.Printf("hello %s", name)
			break
		}
	}

	group := 0
	for i := 0; i < 20; i++ {
		switch {
		case i%2 == 0:
			if i%10 == 0 {
				group++
				break
			}
			fmt.Printf("%02d: %d\n", group, i)
		default:
		}
	}
}
