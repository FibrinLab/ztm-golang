package main

import "fmt"

func main() {
	var myName = "Akan"
	fmt.Println("My name is", myName)

	var name string = "Kathy"
	fmt.Println(name)

	userName := "admin"
	fmt.Println("Username:", userName)

	var sum int
	fmt.Println("The sum is:", sum)

	part1, other := 1, 5
	fmt.Println("Part1:", part1, "Other:", other)

	part2, other := 1, 9
	fmt.Println(part2, other)

	sum = part1 + part2
	fmt.Println("Sum:", sum)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println("Word1:", word1, "Word2:", word2)
}
