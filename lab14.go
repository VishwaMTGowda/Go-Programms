package main

import (
	"fmt"
)

func tower(n int, source, temp, destination rune) {
	if n == 0 {
		return
	}
	tower(n-1, source, destination, temp)
	fmt.Printf("\nMove disc %d from %c to %c", n, source, destination)
	tower(n-1, temp, source, destination)
}

func main() {
	var n int
	fmt.Println("Enter the number of discs:")
	fmt.Scan(&n)
	tower(n, 'A', 'B', 'C')
}
