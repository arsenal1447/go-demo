package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person
	p1.name = "海淀帅哥"
	p1.city = "北京"
	p1.age = 19
	fmt.Printf("p1=%v\n", p1)  //p1={海淀帅哥 北京 19}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"海淀帅哥", city:"北京", age:19}

	// 在定义一些临时数据结构等场景下还可以使用
	// 匿名结构体。
	var user struct {
		Name string
		Age  int
	}
	user.Name = "帅哥"
	user.Age = 19
	fmt.Printf("%#v\n", user) //struct { Name string; Age int }{Name:"帅哥", Age:19}

}
