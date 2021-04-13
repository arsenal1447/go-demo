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
	// var arr = [3]int{1, 5, 6}
	// fmt.Println("arraySum(arr):", arraySum(arr))
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

	// 切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。
	// 切片表达式中的low和high表示一个索引范围（左包含，右不包含），
	// 也就是下面代码中从数组a中选出1<=索引值<4的元素组成切片s，
	// 得到的切片长度=high-low，容量等于得到的切片的底层数组的容量。
	// a := [5]int{1, 2, 3, 4, 5}
	// s := a[1:3]                                                  // s := a[low:high]  索引1，索引2 对应的值
	// fmt.Printf("s:%v len(s):%v cap(s):%v \n", s, len(s), cap(s)) //s:[2 3] len(s):2 cap(s):4

	// s2 := s[3:4]
	// fmt.Printf("s2:%v len(s2):%v cap(s2):%v \n", s2, len(s2), cap(s2)) //s2:[5] len(s2):1 cap(s2):1

	// a[2:]  // 等同于 a[2:len(a)]
	// a[:3]  // 等同于 a[0:3]
	// a[:]   // 等同于 a[0:len(a)]

	// a[low : high : max]
	// 上面的代码会构造与简单切片表达式a[low: high]相同类型、相同长度和元素的切片。另外，它会将得到的结果切片的容量设置为max-low。
	// 在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0。
	//完整切片表达式需要满足的条件是0 <= low <= high <= max <= cap(a)，其他条件和简单切片表达式相同。
	// t := a[1:3:5]
	// fmt.Printf("t:%v len(t):%v cap(t):%v \n", t, len(t), cap(t)) //t:[2 3] len(t):2 cap(t):4

	// 使用make()函数构造切片

	// make([]T, size, cap)
	// T:切片的元素类型
	// size:切片中元素的数量
	// cap:切片的容量

	// c := make([]int, 2, 10)
	// fmt.Println(c)
	// fmt.Println(len(c))
	// fmt.Println(cap(c))

	// 所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。
	// var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	// s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	// s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	// fmt.Println(s1)
	// fmt.Println(s2)
	// fmt.Println(s3)

	//下面的代码中演示了拷贝前后两个变量共享底层数组，
	//对一个切片的修改会影响另一个切片的内容，这点需要特别注意。
	// s1 := make([]int, 3) //[0,0,0]
	// s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	// s2[0] = 100

	// fmt.Println(s1) //[100 0 0]
	// fmt.Println(s2) //[100 0 0]

	// 切片的遍历方式和数组是一致的，支持索引遍历和for range遍历。
	// s := []int{1, 3, 5}

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(i, s[i])
	// }

	// for index, value := range s {
	// 	fmt.Println(index, value)
	// }

	//注意：通过var声明的零值切片可以在append()函数直接使用，无需初始化。
	var s []int
	s = append(s, 1)
	fmt.Println(s) //[1]

	s = append(s, 2, 3, 4)
	fmt.Println(s) //[1 2 3 4]

	s2 := []int{5, 6, 7}
	fmt.Println(s2) //[5 6 7]

	//Go语言的内建函数append()可以为切片动态添加元素。
	//可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。
	s = append(s, s2...)
	fmt.Println(s) //[1 2 3 4 5 6 7]

}
