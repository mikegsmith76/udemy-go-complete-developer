package main

import "fmt"

func main() {
	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range n {
		if n%2 == 0 {
			fmt.Printf("Number %v is even\n", n)
			continue
		}
		fmt.Printf("Number %v is odd\n", n)
	}
}
