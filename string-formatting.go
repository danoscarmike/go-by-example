package main

import (
	"fmt"
	"os"
)

var foobar string = "foobar"

type FooBarBaz struct {
	Foo, Bar, Baz string
	Fizz, Bang    int
	FUBAR         bool
}

func main() {

	p := fmt.Printf

	fbb := FooBarBaz{"foo", "bar", "baz", 1432, 62, true}
	p("%v\n", fbb)
	p("%+v\n", fbb)
	p("%#v\n", fbb)

	p("Type: %T\n", fbb)
	p("bool: %t\n", fbb.FUBAR)
	p("int: %d\n", fbb.Fizz)
	p("bin: %b\n", fbb.Fizz)
	p("char: %c\n", fbb.Bang)
	p("hex: 0x%04x\n", fbb.Fizz)

	p("float: %.2f\n", float64(fbb.Bang))
	p("scientific1: %.3E\n", float64(fbb.Fizz))
	p("scientific2: %.3e\n", float64(fbb.Fizz))

	p("string: %s\n", fbb.Foo)
	p("quotes: %q\n", fbb.Bar)
	p("str-hex: 0x%08x\n", fbb.Baz)

	p("pointer: %p\n", &fbb)

	p("width: |%6d|%6d|\n", fbb.Fizz, fbb.Bang)
	p("width2: |%08.2f|%08.2f|\n", float64(fbb.Fizz), float64(fbb.Bang))
	p("left: |%-6d|\n", fbb.Bang)

	s := fmt.Sprintf("sprintf: a %s", fbb.Foo)
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: this is %s%s\n", fbb.Foo, fbb.Bar)

}
