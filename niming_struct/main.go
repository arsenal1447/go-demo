package main

import "fmt"

type address struct {
	city     string
	province string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address //匿名嵌套结构体
	workPlace
}

type company struct {
	city string
	name string
	address
}

func main() {
	p1 := person{
		name: "zxx",
		age:  19,
		address: address{
			city:     "suzhou",
			province: "anhui",
		},
		workPlace: workPlace{
			city:     "haidian",
			province: "beijing",
		},
	}

	fmt.Println(p1, p1.address.city, p1.workPlace.city)
}
