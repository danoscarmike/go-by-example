package main

import "fmt"

func plusPlus(a int, b int) int {
	return a + b
}

func plusPlusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	fmt.Println(plusPlus(1, 2))
	fmt.Println(plusPlusPlus(1, 2, 3))
}
