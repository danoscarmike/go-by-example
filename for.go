package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("looping until break")
		break
	}

	for k := 0; k <= 10; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}
}
