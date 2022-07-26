package main

import (
	"fmt"
)

type myint int


func (m myint) hello() {
	fmt.Println("myint hello")
}

func main() {
	m := myint(20)
	m.hello()
}
