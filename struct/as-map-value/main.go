package main

import "fmt"

/**
How can I mutate a struct as value map?

Structs are passed by value, including in function arguments and also map values
In order to persist struct changes from receivers, we need to have a map of struct
pointers instead a simple map of structs because:

(&mp[key]).Toggle() -> is an invalid operation

f := mp[key] -> creates a copy
(&f).Toggle() -> mutates a copy and not original mp[key]
*/

type Foo struct {
	toggle bool
}

func (f *Foo) Toggle() {
	f.toggle = !f.toggle
}

func main() {
	mapWithStruct()
	mapWithPointerToStruct()
}

func mapWithStruct() {
	key := "key"
	mp := make(map[string]Foo)
	mp[key] = Foo{}

	fmt.Println("Map holding STRUCT")
	f := mp[key]
	fmt.Printf("  Before toggle: %t\n", mp[key].toggle)
	(&f).Toggle()
	fmt.Printf("  Struct after toggle: %t\n", f.toggle)
	fmt.Printf("  Map struct after toggle: %t\n", mp[key].toggle)
	fmt.Printf("  Same reference? %t\n", f == mp[key])
}

func mapWithPointerToStruct() {
	key := "key"
	mp := make(map[string]*Foo)
	mp[key] = &Foo{}

	fmt.Println("Map holding POINTER to struct")
	f := mp[key]
	fmt.Printf("  Before toggle: %t\n", mp[key].toggle)
	f.Toggle()
	fmt.Printf("  Struct after toggle: %t\n", f.toggle)
	fmt.Printf("  Map struct after toggle: %t\n", mp[key].toggle)
	fmt.Printf("  Same reference? %t\n", f == mp[key])
}
