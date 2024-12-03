package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x // p is a pointer to x

	fmt.Println("Value of x:", x)          // Output: 10
	fmt.Println("Address of x:", p)        // Output: memory address of x
	fmt.Println("Value at address p:", *p) // Output: 10

	*p = 20                           // Modify the value at the address p
	fmt.Println("New value of x:", x) // Output: 20
}
