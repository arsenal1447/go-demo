package main

import (
	"fmt"
)

func main() {
	// zxx(100)
	// zxx1(f2, 30, 40)
	// ret := zxx2(f2, 30, 40)
	// ret()

	res := add(100)
	res2 := res(200)

	fmt.Println(res2)

	res3 := f3(f2, 10, 20)
	f1(res3)

}

// func zxx(x int) {
// 	tmp := func() {
// 		fmt.Println(x)
// 	}
// 	tmp()
// }

// func f2(x, y int) {
// 	fmt.Println("this is f2")
// 	fmt.Println(x + y)
// }

// func zxx1(x func(int, int), m, n int) {
// 	// tmp := func() {
// 	// 	x(m, n)
// 	// }
// 	// tmp()
// 	x(m, n)
// }

// func zxx2(x func(int, int), m, n int) func() {
// 	tmp := func() {
// 		x(m, n)
// 	}
// 	return tmp
// }

func add(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}
func f3(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}
