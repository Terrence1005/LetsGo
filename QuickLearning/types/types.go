package types

import "fmt"

//Pointers -> A pointer holds the memory address of a value
// type *T is a pointer to a T Value, it's zero value is nil
// The & operator generates a pointer to its operand
// The * operator denotes the pointer's underlying value
// Know as "deferencing" or "indirecting"
// No pointer arithmetic

func Pointer() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p) // read i through the pointer
	fmt.Println(p)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}

// Struct collection of fields
type Vertex struct {
	X int
	Y int
}
