package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int64
	name string
}

var allStudent map[int64]*student

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)

	fmt.Print("please input your code:")
	fmt.Scanln(&id)
	fmt.Print("please input your name:")
	fmt.Scanln(&name)

	if _, ok := allStudent[id]; !ok {
		newStu := newStudent(id, name)
		allStudent[id] = newStu
	} else {
		fmt.Print("student existed")
	}
}

func delStudent() {
	var deleteOID int64
	fmt.Println("please input code:")
	fmt.Scanln(&deleteOID)
	if _, ok := allStudent[deleteOID]; !ok {
		fmt.Println("student not exist:")
	} else {
		delete(allStudent, deleteOID)
	}

}

func showStudent() {
	if len(allStudent) != 0 {
		for k, v := range allStudent {
			fmt.Printf("code:%d , name is %s \n", k, v.name)
		}
	} else {
		fmt.Println("student info empty")

	}
}
func main() {
	allStudent = make(map[int64]*student, 50)
	for {
		fmt.Println("welcome to student sys")
		fmt.Println(`
		   1,showStudent
		   2,addStudent
		   3,delStudent
		   4,exit
		`)

		fmt.Print("please input your code:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("you select %d choice \n ", choice)
		switch choice {
		case 1:
			showStudent()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			os.Exit(1)
		}

	}
}
