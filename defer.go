package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f, "Testing write to file...")
}

func createFile(name string) *os.File {
	fmt.Println("creating file:", name)
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, data string) {
	fmt.Println("writing to file:", data)
	fmt.Fprintf(f, "%s\n", data)
}

func closeFile(f *os.File) {
	fmt.Println("closing file:", f.Name())
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
