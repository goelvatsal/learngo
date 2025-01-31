package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Assign the Arrays
//
//  1. Create an array named books
//
//  2. Add book titles to the array
//
//  3. Create two more copies of the array named: upper and lower
//
//  4. Change the book titles to uppercase in the upper array only
//
//  5. Change the book titles to lowercase in the lower array only
//
//  6. Print all the arrays
//
//  7. Observe that the arrays are not connected when they're copied.
//
// NOTE
//  Check out the strings package, it has functions to convert letters to
//  upper and lower cases.
//
// BONUS
//  Invent your own arrays with different types other than string,
//  and do some manipulations on them.
//
// EXPECTED OUTPUT
//   Note: Don't worry about the book titles here, you can use any title.
//
//   books: ["Kafka's Revenge" "Stay Golden" "Everythingship"]
//   upper: ["KAFKA'S REVENGE" "STAY GOLDEN" "EVERYTHINGSHIP"]
//   lower: ["kafka's revenge" "stay golden" "everythingship"]
// ---------------------------------------------------------

func main() {
	books := [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	lower := books
	lower[0] = "kafka's revenge"
	lower[1] = "stay golden"
	lower[2] = "everythingship"

	upper := books
	upper[0] = "KAFKA'S REVENGE"
	upper[1] = "STAY GOLDEN"
	upper[2] = "EVERYTHINGSHIP"

	fmt.Println("books:", books)
	fmt.Println("upper:", upper)
	fmt.Println("lower:", lower)
}
