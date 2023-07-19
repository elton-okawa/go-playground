package main

import (
	"fmt"
)

/**
How to return chunks of a slice?
*/

var size int = 3

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(s); i += size {
		end := min(i+3, len(s))

		c := s[i:end]
		fmt.Printf("chunk %d: %v\n", i/size, c)
	}
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}
