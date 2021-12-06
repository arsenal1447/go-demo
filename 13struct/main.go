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
	fmt.Printf("%p \n", p2) //0xc000028040

	var a int
	a = 100
	b := &a
	fmt.Printf("type a %T , type b %T \n", a, b) //type a int , type b *int
	//a的16进制地址
	fmt.Printf(" a p=%p ,  b p=%p \n", &a, b) // a 0xc000016090 ,  b 0xc000016090
	fmt.Printf(" b p= %p \n", &b)             //  b 0xc00000e030
	fmt.Printf(" b v= %v ", b)                //  b 0xc000016090

	var p3 = Person{
		name:   "zxx",
		gender: "man",
	}
	fmt.Printf(" p3 v= %v ", p3) //  p3 v= {zxx man}

	p4 := Person{
		"chow",
		"woman",
	}
	fmt.Printf(" p4 v= %v ", p4) //p4 v= {chow woman}

}

func f1(p Person) {
	p.gender = "woman1"
}

func f2(p *Person) {
	// (*p).gender = "woman2"
	p.gender = "woman2"
}
