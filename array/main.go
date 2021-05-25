package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("./hello.tmpl")

	if err != nil {
		fmt.Println("create template failed , err: ", err)
		return
	}
	// 自定义一个夸人的模板函数
	kua := func(arg string) (string, error) {
		return arg + "真帅", nil
	}

	// 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
	tmpl, err := template.New("hello").Funcs(template.FuncMap{"kua": kua}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	user := UserInfo{
		Name:   "beijijng boy",
		Gender: "man",
		Age:    18,
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, user)
}

// 将上面的main.go文件编译执行，然后使用浏览器访问http://127.0.0.1:9090 就能看到页面上显示了“beijing boy”。
// 这就是一个最简单的模板渲染的示例
// 同理，当我们传入的变量是map时，也可以在模板文件中通过.根据key来取值。

func main() {
	// http.HandleFunc("/", sayHello)
	// http.HandleFunc("/", tmplDemo)
	http.HandleFunc("./templates", index)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed , err: ", err)
		return
	}

}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("create template failed, err:", err)
		return
	}

	user := UserInfo{
		Name:   "beijing",
		Gender: "man",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

// 然后使用template.ParseGlob按照正则匹配规则解析模板文件，然后通过ExecuteTemplate渲染指定的模板：

// 如果我们的模板名称冲突了，例如不同业务线下都定义了一个index.tmpl模板，我们可以通过下面两种方法来解决。
// 在模板文件开头使用{{define 模板名}}语句显式的为模板命名。
// 可以把模板文件存放在templates文件夹下面的不同目录中，然后使用template.ParseGlob("templates/**/*.tmpl")解析模板。
func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*.tmpl")

	if err != nil {
		fmt.Println("create template failed ,err", err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index.tmpl", nil)
	if err != nil {
		fmt.Println("render template failed ,err", err)
		return
	}

}

func xss(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")

	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	jsStr := `<script>alert('hahaha')</script>`
	err = tmpl.Execute(w, jsStr)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

}
