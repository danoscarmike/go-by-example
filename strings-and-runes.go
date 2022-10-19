package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี"
	const s2 = "tree"

	fmt.Println("Len: ", len(s))
	fmt.Println("Len: ", len(s2))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	for i := 0; i < len(s2); i++ {
		fmt.Printf("%x ", s2[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s2))

	for idx, runeValue := range s {
		fmt.Printf("s: %#U begins at rune %d\n", runeValue, idx)
	}

	for idx, runeValue := range s2 {
		fmt.Printf("s2: %#U begins at rune %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString on s")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		w = width
		fmt.Printf("Examining rune starting at index %d\n", i)
		examineRune(runeValue)
	}

	fmt.Println("\nUsing DecodeRuneInString on s2")
	for i, w := 0, 0; i < len(s2); i += w {
		runeValue, width := utf8.DecodeRuneInString(s2[i:])
		w = width
		fmt.Printf("Examining rune starting at index %d\n", i)
		examineRune(runeValue)
	}

}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
