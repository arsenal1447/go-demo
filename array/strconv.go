package main

// os.Open()函数能够打开一个文件，返回一个*File和一个err。
// 对得到的文件实例调用close()方法能够关闭文件

// Read方法定义如下：

// func (f *File) Read(b []byte) (n int, err error)
// 它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF。 举个例子：

func main() {
	//Atoi()函数用于将字符串类型的整数转换为int类型，函数签名如下。
	// s1 := "100"
	// i1, err := strconv.Atoi(s1)
	// if err != nil {
	// 	fmt.Println("can't convert to int")
	// } else {
	// 	fmt.Printf("type:%T value :%#v \n", i1, i1) //type:int value :100
	// }

	//Itoa()函数用于将int类型数据转换为对应的字符串表示，具体的函数签名如下。
	//【扩展阅读】这是C语言遗留下的典故。C语言中没有string类型而是用字符数组(array)表示字符串，所以Itoa对很多C系的程序员很好理解。
	// i2 := 200
	// s2 := strconv.Itoa(i2)
	// fmt.Printf("type : %T value :%#v \n", s2, s2) //type : string value :"200"

	//ParseBool()  返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	//ParseInt() 返回字符串表示的整数值，接受正负号。
	//ParseUint() 类似ParseInt但不接受正负号，用于无符号整型。
	//ParseFloat() 解析一个表示浮点数的字符串并返回其值。

	//type:bool value :true
	// type:int64 value :-2
	// type:float64 value :3.144
	// type:uint64 value :0x2
	// b, err := strconv.ParseBool("true") //type:bool value :true
	// if err != nil {
	// 	fmt.Println("can't convert to int")
	// } else {
	// 	fmt.Printf("type:%T value :%#v \n", b, b)
	// }

	// f, err := strconv.ParseInt("-2", 10, 64) //type:int64 value :-2
	// if err != nil {
	// 	fmt.Println("can't convert to int")
	// } else {
	// 	fmt.Printf("type:%T value :%#v \n", f, f)
	// }

	// c, err := strconv.ParseFloat("3.1415", 64) //type:float64 value :3.1415
	// if err != nil {
	// 	fmt.Println("can't convert to int")
	// } else {
	// 	fmt.Printf("type:%T value :%#v \n", c, c)
	// }

	// d, err := strconv.ParseUint("2", 10, 64) //type:uint64 value :0x2
	// if err != nil {
	// 	fmt.Println("can't convert to int")
	// } else {
	// 	fmt.Printf("type:%T value :%#v \n", d, d)
	// }

	// FormatBool() 根据b的值返回”true”或”false”。

	// FormatInt() 返回i的base进制的字符串表示。 base必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。

	// FormatUint() 是FormatInt的无符号整数版本。

	// FormatFloat()  函数将浮点数表示为字符串并返回。

	// a1 := strconv.FormatBool(true)
	// a2 := strconv.FormatInt(-2, 16)
	// a3 := strconv.FormatUint(2, 16)
	// a4 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	// //a1:true
	// // a2: -2
	// // a3:	2
	// // a4: 3.1415E+00
	// fmt.Printf("a1:%v \n a2:%v \n a3:%v \n a4:%v", a1, a2, a3, a4)

}
