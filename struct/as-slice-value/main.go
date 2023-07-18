package main

import "fmt"

/**

 */

type Foo struct {
	modified bool
}

func main() {
	sliceWithStruct()
}

func sliceWithStruct() {
	s := make([]Foo, 1)

	fmt.Println("Slice holding STRUCT")

	fmt.Printf("  Before change: %t\n", s[0].modified)
	s[0].modified = true

	fmt.Printf("  After change: %t\n", s[0].modified)
}
