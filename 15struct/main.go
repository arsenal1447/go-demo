package main

//https://www.liwenzhou.com/posts/Go/10_struct/#autoid-2-4-1
import (
	"encoding/json"
	"fmt"
)

//一个结构体中可以嵌套包含另一个结构体或结构体指针，就像下面的示例代码那样。
type Address struct {
	Province   string
	City       string
	CreateTime string
}

type User struct {
	Name    string
	Gender  string
	Address Address
	Email
}

//上面user结构体中嵌套的Address结构体也可以采用匿名字段的方式，例如：
type People struct {
	Name    string
	Gender  string
	Address //匿名字段
}

type Email struct {
	Account    string
	CreateTime string
}

func main() {
	user1 := User{
		Name:   "zxx",
		Gender: "man",
		Address: Address{
			Province: "anhui",
			City:     "suzhou",
		},
	}
	//user1 = main.User{Name:"zxx", Gender:"man", Address:main.Address{Province:"anhui", City:"suzhou"}}
	fmt.Printf("user1 = %#v\n", user1)
	var peo People
	peo.Name = "zxx2"
	peo.Gender = "man"
	peo.Address.Province = "beijing" //匿名字段默认使用类型名作为字段名
	peo.City = "beijing"             // 匿名字段Address可以省略

	//当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。
	//peo = main.People{Name:"zxx2", Gender:"man", Address:main.Address{Province:"beijing", City:"beijing"}}
	fmt.Printf("peo = %#v\n", peo)

	//嵌套结构体的字段名冲突
	//嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名。
	var user3 User
	user3.Name = "zxx3"
	user3.Gender = "man"
	user3.Address.CreateTime = "2000"
	user3.Email.CreateTime = "2000"
	//user3 = main.User{Name:"zxx3", Gender:"man", Address:main.Address{Province:"", City:"", CreateTime:"2000"},
	//Email:main.Email{Account:"", CreateTime:"2000"}}
	fmt.Printf("user3 = %#v\n", user3)

	//Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}

	d1.move() //乐乐 会动
	d1.wang() //乐乐 会汪汪汪

	//结构体与JSON序列化
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "man",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}

	fmt.Printf("json:%s\n", data)
}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s 会动\n", a.name)
}

func (d *Dog) wang() {
	fmt.Printf("%s 会汪汪汪\n", d.name)
}

type Dog struct {
	Feet int8
	*Animal
}

type Student struct {
	ID     int
	Name   string
	Gender string
}

type Class struct {
	Title    string
	Students []*Student
}
