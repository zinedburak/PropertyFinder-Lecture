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
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	var (
		argLen = len(os.Args) - 1
		name1  = os.Args[1]
		name2  = os.Args[2]
		name3  = os.Args[3]
	)

	fmt.Println("There are", argLen, "people !")
	fmt.Println("Hello great", name1, "!")
	fmt.Println("Hello great", name2, "!")
	fmt.Println("Hello great", name3, "!")
	fmt.Println("Nice to meet you all.")
}