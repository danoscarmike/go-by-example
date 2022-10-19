package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	res1 := sum(1, 2)
	fmt.Println(res1)

	res2 := sum(1, 2, 3)
	fmt.Println(res2)

	nums := []int{1, 2, 3, 4}
	res3 := sum(nums...)
	fmt.Println(res3)
}
