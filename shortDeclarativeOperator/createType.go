package main

import (
	"fmt"
)
var a int
type hotdog int
var b hotdog

func main() {
	a = 42
	b = 22
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)


}

