// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Print the runes
//
//  1. Loop over the "console" word and print its runes one by one,
//     in decimals, hexadecimals and binary.
//
//  2. Manually put the runes of the "console" word to a byte slice, one by one.
//
//     As the elements of the byte slice use only the rune literals.
//
//     Print the byte slice.
//
//  3. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only decimal numbers.
//
//  4. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only hexadecimal numbers.
//
//
// EXPECTED OUTPUT
//   Run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	const word = "console"

	var bWord []byte
	var dWord []byte
	var hWord []byte
	for _, v := range word {
		fmt.Printf("%q\n", string(v))
		fmt.Printf("\tdecimal: %d\n", v)
		fmt.Printf("\thex: %x\n", v)
		fmt.Printf("\tbinary: %b\n", v)
		bWord = append(bWord, byte(v))
		dWord = append(dWord, byte(v))
		hWord = append(hWord, byte(v))
	}
	fmt.Println("printed using binary:", string(bWord))
	fmt.Println("printed using decimal:", string(dWord))
	fmt.Println("printed using hexadecimal:", string(hWord))
}
