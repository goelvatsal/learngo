package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Math Tables
//
//  Create division, addition and subtraction tables
//
//  1. Get the math operation and
//     the size of the table from the user
//
//  2. Print the table accordingly
//
//  3. Correctly handle the division by zero error
//
//
// BONUS #1
//
//  Use strings.IndexAny function to detect
//    the valid operations.
//
//  Search on Google for: golang pkg strings IndexAny
//
// BONUS #2
//
//  Add remainder operation as well (remainder table using %).
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Usage: [op=*/+-] [size]
//
//  go run main.go "*"
//    Size is missing
//    Usage: [op=*/+-] [size]
//
//  go run main.go "%" 4
//    Invalid operator.
//    Valid ops one of: */+-
//
//  go run main.go "*" 4
//    *    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    2    3    4
//    2    0    2    4    6    8
//    3    0    3    6    9   12
//    4    0    4    8   12   16
//
//  go run main.go "/" 4
//    /    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    0    0    0
//    2    0    2    1    0    0
//    3    0    3    1    1    0
//    4    0    4    2    1    1
//
//  go run main.go "+" 4
//    +    0    1    2    3    4
//    0    0    1    2    3    4
//    1    1    2    3    4    5
//    2    2    3    4    5    6
//    3    3    4    5    6    7
//    4    4    5    6    7    8
//
//  go run main.go "-" 4
//    -    0    1    2    3    4
//    0    0   -1   -2   -3   -4
//    1    1    0   -1   -2   -3
//    2    2    1    0   -1   -2
//    3    3    2    1    0   -1
//    4    4    3    2    1    0
//
// BONUS:
//
//  go run main.go "%" 4
//    %    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    0    1    1    1
//    2    0    0    0    2    2
//    3    0    0    1    0    3
//    4    0    0    0    1    0
//
// NOTES:
//
//   When running the program in Windows Git Bash, you might need
//   to escape the characters like this.
//
//   This happens because those characters have special meaning.
//
//   Division:
//     go run main.go // 4
//
//   Multiplication and others:
//   (this is also necessary for non-Windows bashes):
//
//     go run main.go "*" 4
// ---------------------------------------------------------

func main() {
	argOne := os.Args[1]
	argTwo := os.Args[2]

	if len(os.Args) != 2 {
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}

	if strings.IndexAny(argOne, "*") == 0 {
		n, _ := strconv.Atoi(argTwo)
		mulTable(n)
	} else if strings.IndexAny(argOne, "/") == 0 {
		n, _ := strconv.Atoi(argTwo)
		divTable(n)
	} else if strings.IndexAny(argOne, "+") == 0 {
		n, _ := strconv.Atoi(argTwo)
		addTable(n)
	} else if strings.IndexAny(argOne, "-") == 0 {
		n, _ := strconv.Atoi(argTwo)
		subTable(n)
	} else {
		fmt.Println("Invalid operator.\nValid ops one of: */+-")
		return
	}

}

func mulTable(n int) int {
	fmt.Printf("%5s", "*")

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= n; j++ {
			fmt.Printf("%5d", j*i)
		}
		fmt.Println()
	}
	return 0
}

func divTable(n int) int {
	fmt.Printf("%5s", "/")

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= n; j++ {
			if j == 0 {
				fmt.Printf("%5v", "0")
				continue
			}
			fmt.Printf("%5d", i/j)
		}
		fmt.Println()
	}
	return 0
}

func subTable(n int) int {
	fmt.Printf("%5s", "-")
	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= n; j++ {
			fmt.Printf("%5d", i-j)
		}
		fmt.Println()
	}
	return 0
}

func addTable(n int) int {
	fmt.Printf("%5s", "+")

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= n; j++ {
			fmt.Printf("%5d", j+i)
		}
		fmt.Println()
	}
	return 0
}
