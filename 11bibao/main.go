package main

import (
	"fmt"
	"strings"
)

func main() {
	jpgFunc := checkSuffix("jpg")
	txtFunc := checkSuffix("txt")

	fmt.Println(jpgFunc("aaa.jpg"))
	fmt.Println(jpgFunc("bbb"))

	fmt.Println(txtFunc("ccc.txt"))
	fmt.Println(txtFunc("dddd"))
}

func checkSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
