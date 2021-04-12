package main

import (
	"fmt"
)

func main() {
	// var testArray [3]int

	// var numArray = [3]int{1, 2}
	// var numArray1 = [...]int{1, 2}

	// var cityArray = [3]string{"北京", "上海", "深圳"}
	// var cityArray1 = [...]string{"北京", "上海", "深圳"}

	// fmt.Println(testArray)

	// fmt.Println(numArray)
	// fmt.Printf("type of numArray:%T\n", numArray)
	// fmt.Println(numArray1)
	// fmt.Printf("type of numArray1:%T\n", numArray1)

	// fmt.Println(cityArray)
	// fmt.Printf("type of cityArray:%T\n", cityArray)
	// fmt.Println(cityArray1)
	// fmt.Printf("type of cityArray1:%T\n", cityArray1)

	// b := [...]int{1: 1, 3: 5}
	// fmt.Println(b)
	// fmt.Printf("type of a %T", b)

	// var a = [...]string{"北京", "上海", "深圳"}
	// // 方法1：for循环遍历
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// // 方法2：for range遍历
	// for index, value := range a {
	// 	fmt.Println(index, value)
	// }

	// 二维数组的定义
	// c := [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }

	// fmt.Println(c)
	// fmt.Println(c[2][1])

	// 二维数组的遍历
	// d := [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }

	// for _, v1 := range d {
	// 	for _, v2 := range v1 {
	// 		fmt.Printf("%s\t", v2)
	// 	}
	// 	fmt.Println()
	// }

	// 多维数组只有第一层可以使用...来让编译器推导数组长度。例如：
	//支持的写法
	// e := [...][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	//不支持多维数组的内层使用...
	// f := [3][...]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }

	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}

// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
