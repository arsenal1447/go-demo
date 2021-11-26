package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "大帅哥"
	} else {
		talk = "hello"
	}
	return
}

func main() {
	var peo People = &Student{}//可以接收 &Student类型
  
	//var peo People = Student{}//报错，不可以接收 Student类型，可以接收 &Student类型
	think := "abc"
	fmt.Println(peo.Speak(think))
}
