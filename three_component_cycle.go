package main

import "fmt"

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
	
}
