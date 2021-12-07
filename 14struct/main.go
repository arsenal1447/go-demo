package main

import (
	"fmt"
	"unsafe"
)

//类型定义
type NewInt int

//类型别名
type MyInt = int

type person struct {
	name string
	city string
	age  int8
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

type student struct {
	name string
	age  int
}

func main() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T \n", a) //type of a:main.NewInt

	fmt.Printf("type of b:%T \n", b) //type of b:int

	var user struct {
		Name string
		Age  int
	}

	user.Age = 18
	user.Name = "zxx"

	fmt.Printf("type of user:%#v\n", user) //type of user:struct { Name string; Age int }{Name:"zxx", Age:18}

	//使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
	p5 := person{
		name: "zxx",
		city: "beijing",
		age:  18,
	}

	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"zxx", city:"beijing", age:18}

	//也可以对结构体指针进行键值对初始化，例如：
	p6 := &person{
		name: "zxx",
		city: "beijing",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"zxx", city:"beijing", age:18}

	//当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	p7 := &person{
		name: "zxx",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"zxx", city:"", age:0}

	//初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
	// 1.必须初始化结构体的所有字段。
	// 2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 3.该方式不能和键值初始化方式混用
	p8 := &person{
		"zxx", "beijing", 18,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"zxx", city:"beijing", age:18}

	n := test{
		1, 2, 3, 4,
	}
	//结构体占用一块连续的内存。
	fmt.Printf("n.a %p\n", &n.a) //n.a 0xc000016110
	fmt.Printf("n.b %p\n", &n.b) //n.b 0xc000016111
	fmt.Printf("n.c %p\n", &n.c) //n.c 0xc000016112
	fmt.Printf("n.d %p\n", &n.d) //n.d 0xc000016113

	//空结构体是不占用空间的。
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) //0

	m := make(map[string]*student)

	stus := []student{
		{name: "zxx1", age: 18},
		{name: "zxx2", age: 19},
		{name: "zxx3", age: 20},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	//zxx1 => zxx3
	// zxx2 => zxx3
	// zxx3 => zxx3
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
