package main

import "fmt"

func arraySum(x [3]int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 切片的长度和容量
// 切片拥有自己的长度和容量，
// 我们可以通过使用内置的len()函数求长度，
// 使用内置的cap()函数求切片的容量。

// 切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。它有两种变体：
// 一种指定low和high两个索引界限值的简单的形式，
// 另一种是除了low和high索引界限值外还指定容量的完整的形式。

func main() {
	var arr = [3]int{1, 5, 6}
	fmt.Println("arraySum(arr):", arraySum(arr))
	// 声明切片类型
	// var a []string              //声明一个字符串切片
	// var b = []int{}             //声明一个整型切片并初始化
	// var c = []bool{false, true} //声明一个布尔切片并初始化
	// var d = []bool{false, true} //声明一个布尔切片并初始化

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(a == nil) //true
	// fmt.Println(b == nil) //false
	// fmt.Println(c == nil) //false
	// fmt.Println(d == nil) //false

}
