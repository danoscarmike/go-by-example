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

	p("HasPrefix:", s.HasPrefix("fooool", "foooo"))
	p("HasSuffix:", s.HasSuffix("fooool", "ooool"))
	p("Index:", s.Index("fooool", "oooo"))
	p("IndexAny:", s.IndexAny("fooool", "jkl;"))
	p("IndexFunc:", s.IndexFunc("foo☘️ool", f))
	p("IndexRune:", s.IndexRune("foo☘️ool", '\u2618'))
	p("Join:", s.Join([]string{"f", "o", "☘️", "l"}, "_"))
	p("LastIndex:", s.LastIndex("fooool", "oo"))
	p("LastIndexAny:", s.LastIndexAny("fooool", "fol"))

	g := func(r rune) rune {
		switch r {
		case '\u2618':
			return 'o'
		case '\u0020':
			return '!'
		}
		return r
	}
	p("Map:", s.Map(g, "fo☘️l "))

	p("Repeat:", s.Repeat("fool ", 3))
	p("Replace:", s.Replace("fool", "o", "☘️", -1))
	p("Split:", s.Split("fo-ol", "o"))
	p("SplitAfter:", s.SplitAfter("fool", "o"))
	p("SplitN:", s.SplitN("foool", "o", 4))
	p("Title:", s.Title("fool"))
	p("ToLower:", s.ToLower("fOoL"))
	p("ToTitle:", s.ToTitle("fOoL Hardy"))
	p("ToUpper:", s.ToUpper("fOoL Hardy"))
	p("Trim:", s.Trim(" fooool-", " -"))
	p("TrimPrefix:", s.TrimPrefix("prefool", "pre"))
	p("TrimSuffix:", s.TrimSuffix("foolsufferer", "sufferer"))
	p("TrimSpace:", s.TrimSpace("   fooooool    "))
}
