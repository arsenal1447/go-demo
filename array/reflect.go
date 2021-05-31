package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var a float32 = 3.14
	// reflectType(a) //type:float32

	// var b int64 = 100
	// reflectType(b) //type:int64

	// var a *float32
	// var b myInt
	// var c rune
	// reflectType(a) //type: kind:ptr
	// reflectType(b) //type:myInt kind:int64
	// reflectType(c) //type:int32 kind:int32

	// type person struct {
	// 	name string
	// 	age  int
	// }

	// type book struct{ title string }
	// var d = person{
	// 	name: "北京boy",
	// 	age:  18,
	// }
	// var e = book{title: "我来学习go语言"}
	// reflectType(d) //type:person kind:struct

	// reflectType(e) //type:book kind:struct

	//通过反射获取值
	// var a float32 = 3.14

	// var b int64 = 100

	// reflectValue(a) //type is Float32,value is 3.140000
	// reflectValue(b) //type is int64,value is 100
	// c := reflect.ValueOf(10)
	// fmt.Printf("type c :%T\n", c) //type c :reflect.Value

	// 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。
	// var a int64 = 100
	// // reflectSetValue1(a) //panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
	// reflectSetValue2(&a)
	// fmt.Println(a)

	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil()) //var a *int IsNil: true
	fmt.Println("Nil IsValid:", reflect.ValueOf(nil).IsValid())  //Nil IsValid: false

	//实例化一个匿名结构体
	b := struct{}{}
	//尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())  //不存在的结构体成员: false
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid()) //不存在的结构体方法: false
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("boy")).IsValid()) //map中不存在的键: false

}

// 在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。
func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

type myInt int64

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Int()从反射中获取整型的原始值，然后通过 float32()强制类型转换
		fmt.Printf("type is Float32,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Int()从反射中获取整型的原始值，然后通过 float64()强制类型转换
		fmt.Printf("type is int64,value is %f\n", float64(v.Float()))
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
