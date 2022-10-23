package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)h", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)h")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatch("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)h")
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("peach punch pinch", "ooo"))

	fmt.Println(string(r.ReplaceAllFunc([]byte("peach punch pinch"), bytes.ToUpper)))
}
