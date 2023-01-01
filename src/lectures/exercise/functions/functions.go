//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
)

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func add(a, b, c int) int {
	return a + b + c
}

// * Write a function that returns any number
func number() int {
	return 5
}

// * Write a function that returns any two numbers
func returnTwo() (int, int) {
	first := 5
	second := 6

	return first, second
}

//* Add three numbers together using any combination of the existing functions.

//  * Print the result
//* Call every function at least once

func main() {
	greet("John")
	add(1, 2, 3)
	rand := number()
	fmt.Println(rand)
	returnTwo()

	addThree := add(number(), number(), number())
	fmt.Println(addThree)

}
