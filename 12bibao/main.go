package main

import (
	"fmt"
)

func main() {
	f1, f2 := calc(10)

	fmt.Println(f1(4), f2(4))
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
