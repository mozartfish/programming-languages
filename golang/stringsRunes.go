package main

import (
	"fmt"
	"unicode/utf8"
)

// - compare rune value to a rune literal
func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

// A go string is a read-only slice of bytes. Go and the standard library treat strings as containers of text encoded in UTF-8.
// In other languages, strings are made up of characters
// In Go, the concept of a character is called a Rune - an integer that represents a Unicode point
// - values enclosed in single quotes are rune literals
func stringsRunes() {

	fmt.Println("STRINGS AND RUNES")
	// - s is a string asigned a literal value representing hello in the Thai language
	// - Go string literals are UTF-8 encoded text
	const s = "สวัสดี"

	// - since strings are equivalent to []byte in Go, len produces the length of the raw bytes stored within
	fmt.Println("len:", len(s))

	// - indexing into a string produces the raw byte value at each index
	// - the loop generates the hex values of all the bytes that constitute the code points in s
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// - the utf8 package provides a way of counting the number of runes in a string
	// - the runtime of RuneCountInString depends on the size of the string because it has to
	// decode each UTF-8 rune sequentially. Some Thai characters are represented by multiple UTF-8 code points
	// so the count of hello in Thai may be surprising
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// - the range loop handles strings specifically and decodes each rune along with its offset
	// in the string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// - utf8 builtin package to decode runes in string
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// pass rune value to a function
		examineRune(runeValue)
	}
}
