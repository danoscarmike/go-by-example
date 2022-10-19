package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

type container2 struct {
	base
	str string
	num int
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Printf("co={num: %v, str: %v}\n", co.base.num, co.str)
	fmt.Println("describe:", co.describe())

	co2 := container2{
		base: base{
			num: 2,
		},
		str: "other name",
		num: 22,
	}

	fmt.Println()
	fmt.Printf("co2={num: %v, str: %v}\n", co2.num, co2.str)
	fmt.Printf("co2={num: %v, str: %v}\n", co2.base.num, co2.str)
	fmt.Println("describe:", co2.describe())

	type describer interface {
		describe() string
	}

	fmt.Println()
	var d describer = co
	fmt.Println("describer:", d.describe())

	var d2 describer = co2
	fmt.Println("describer2:", d2.describe())

}
