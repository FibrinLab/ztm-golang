package main

import "fmt"

func double(num int) int {
	return num + num
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet function")
}

func main() {
	greet()
	dozen := double(6)
	fmt.Println(dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println(bakersDozen)

}
