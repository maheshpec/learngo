package main

import "fmt"

func main() {
	var a int
	var b int
	var alocation *int
	var x interface{}

	a = 0x23
	b = 0x123
	alocation = &a

	fmt.Printf("hello, world %x \n", a*b)
	fmt.Printf("What's up %d\n", *alocation)

	switch x.(type) {
	case nil:
		fmt.Println("Case is nil")
	}
    a, b = swap(a, b)
	fmt.Printf("Added value %d %d", a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}
