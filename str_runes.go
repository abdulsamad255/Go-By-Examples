package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// A Unicode string (Thai language)
	// Each character uses multiple bytes
	const s = "สวัสดี"

	// -------------------------------
	// len(s) returns number of BYTES
	// -------------------------------
	fmt.Println("Len:", len(s))

	// -------------------------------
	// Print each byte in hexadecimal
	// This does NOT represent characters
	// -------------------------------
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// -------------------------------
	// Count actual Unicode characters (runes)
	// -------------------------------
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// -------------------------------
	// Range over string
	// idx = byte index
	// runeValue = Unicode character (rune)
	// -------------------------------
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// -------------------------------
	// Manually decode runes using utf8
	// -------------------------------
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {

		// Decode first rune from remaining string
		runeValue, width := utf8.DecodeRuneInString(s[i:])

		// Print rune and its starting byte index
		fmt.Printf("%#U starts at %d\n", runeValue, i)

		// width tells how many bytes this rune uses
		w = width

		// Examine specific rune values
		examineRune(runeValue)
	}
}

// -------------------------------
// Function to compare rune values
// -------------------------------
func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
