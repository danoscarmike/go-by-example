package main

import "fmt"

func interator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := interator()
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())

	nextInt2 := interator()
	fmt.Println("nextInt2:", nextInt2())

}
