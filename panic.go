package main

import "os"

func main() {
	panic("danger, Will Robinson")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
