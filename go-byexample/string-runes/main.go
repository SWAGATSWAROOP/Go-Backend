package main

import (
	"fmt"
	"unicode/utf8"
)

// A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”. In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point. This Go blog post is a good introduction to the topic.

//Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within

// Indexing into a string produces the raw byte values at each index. This loop generates the hex values of all the bytes that constitute the code points in s.

func main() {
	const a = "स्वागत"

	fmt.Println("Len: ", len(a))

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}

	fmt.Println()

	// We cannot declare named function inside another function but we can declare anonymous function
	examineRune := func(r rune) {
		if r == 't' {
			fmt.Println("found tee")
		} else if r == 'a' {
			fmt.Println("Found a")
		}
	}

	fmt.Println("Rune Count: ", utf8.RuneCountInString(a))

	for idx, Value := range a {
		fmt.Println("The value is ", Value, " and index is ", idx)
	}

	fmt.Println("To get actual value we can use utf8.DecodeRuneInString")

	/*
		for i, w := 0, 0; i < len(a); i += w
		i is the byte index into the string (not character index).
		w is the width in bytes of the last rune.
		After each iteration, i += w moves i forward by the number of bytes in the last rune.
		This ensures you're stepping over each complete UTF-8 character, regardless of its byte size (1–4 bytes).

		🧠 utf8.DecodeRuneInString(a[i:])
		This decodes the next rune (Unicode code point) starting at index i.
		It returns:
		runeValue → the actual rune
		width → the number of bytes this rune occupies
		Example:
		For the string "Swagat":

		Each ASCII character like S, w, a, etc. is 1 byte wide in UTF-8. So:

		On first iteration: i = 0, a[0:] = "Swagat", runeValue = 'S', width = 1
		Then i += 1 → i = 1, decode 'w', and so on.
		If the string had multi-byte Unicode characters (like emojis or characters from Hindi, Japanese, etc.), the width would be >1 and the loop would still work perfectly.

		🖨️ fmt.Printf("%#U starts at %d\n", runeValue, i)
		%#U formats the rune as U+XXXX 'char', e.g., U+0053 'S'
		i shows the starting byte index of the rune in the original string
	*/
	for i, w := 0, 0; i < len(a); i += w {
		runeValue, width := utf8.DecodeRuneInString(a[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

}
