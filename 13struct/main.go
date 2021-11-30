package main

import "fmt"

type Person struct {
	name, gender string
}

func main() {
	var p Person
	p.name = "zhangsan"
	p.gender = "man"
	f1(p)
	fmt.Println(p.gender) //man

	f2(&p)
	fmt.Println(p.gender) //woman2

	var p2 = new(Person)
	fmt.Printf("%T \n", p2) //*main.Person
	p2.name = "zxx"
	fmt.Println(p2)         //&{zxx }
	fmt.Printf("%x \n", p2) //&{7a7878 }

}

func f1(p Person) {
	p.gender = "woman1"
}

func f2(p *Person) {
	// (*p).gender = "woman2"
	p.gender = "woman2"
}
