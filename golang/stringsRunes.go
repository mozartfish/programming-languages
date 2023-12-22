package main

import (
	"fmt"
	"unicode/utf8"
)

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
		// Rune literal 
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

// A Go string is a read-only slice of bytes
// The language and the standard library treat strings specially
// as containers of text encoded in UTF-8
// In other languages strings are made up of characters. In Go, the
// concept of characters is called a RUNE -> integer representing a Unicode
// point.

func stringsRunes() {
	// hello in Thai represented in UTF-8 encoded text
	const s = "สวัสดี"

	// length of the raw bytes in s
	fmt.Println("Len:", len(s))

	// print hex values of all bytes that make up the code points
	// in s
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune Count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// Pass rune value to a function
		examineRune(runeValue)
	}

}
