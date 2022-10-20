package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var p = fmt.Println

func main() {
	p("Replace:", s.Replace("fooooool", "o", "0", 2))
	p("Contains:", s.Contains("fooooool", "0"))
	p("Count:", s.Count("fooooool", "o"))
	p("Compare:", s.Compare("fooooool", "foll"))
	p("Contains any:", s.ContainsAny("fooooool", "asdf"))
	p("Contains rune:", s.ContainsRune("fooooool", '\u2618'))
	p("Contains rune:", s.ContainsRune("fooo☘️oool", '\u2618'))

	before, after, found := s.Cut("fooool", "oooo")
	p("Cut:", found, before, after)

	p("EqualFold:", s.EqualFold("fooool", "fOoOoL"))

	for _, field := range s.Fields("fo o   oo      l") {
		p(field)
	}

	f := func(r rune) bool {
		return !unicode.IsLetter(r)
	}

	p("FieldsFunc:", s.FieldsFunc("foo☘️ool", f))

}
