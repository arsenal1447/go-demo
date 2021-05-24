package main

import (
	"fmt"
	"io"
	"os"
)

// os.Open()函数能够打开一个文件，返回一个*File和一个err。
// 对得到的文件实例调用close()方法能够关闭文件

// Read方法定义如下：

// func (f *File) Read(b []byte) (n int, err error)
// 它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF。 举个例子：

func main() {
	// // 只读方式打开当前目录下的 xx.log 文件
	// file, err := os.Open("./xx.log")
	// if err != nil {
	// 	fmt.Println("open file failed, err:", err)
	// 	return
	// }
	// // 为了防止文件忘记关闭，我们通常使用defer注册文件关闭语句。
	// defer file.Close()

	// var tmp = make([]byte, 128)
	// n, err := file.Read(tmp)
	// if err == io.EOF {
	// 	fmt.Println("文件读完了")
	// 	return
	// }

	// if err != nil {
	// 	fmt.Println("read file failed, err:", err)
	// 	return
	// }
	// fmt.Printf("读取了%d字节数据\n:", n)
	// fmt.Println(string(tmp[:n]))

	//************************ for 循环读取文件 *******************************
	// 使用for循环读取文件中的所有数据。
	// file, err := os.Open("./xx.log")
	// if err != nil {
	// 	fmt.Println("open file failed ,err", err)
	// 	return
	// }

	// defer file.Close()
	// var content []byte
	// var tmp = make([]byte, 128)

	// for {
	// 	n, err := file.Read(tmp)
	// 	if err == io.EOF {
	// 		fmt.Println("文件读完了")
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("read file failed ,err", err)
	// 		return
	// 	}
	// 	content = append(content, tmp[:n]...)
	// }
	// fmt.Println(string(content))

	//************************ bufio读取文件 *******************************

	// file, err := os.Open("./xx.log")
	// if err != nil {
	// 	fmt.Println("open file failed ,err", err)
	// 	return
	// }
	// defer file.Close()

	// reader := bufio.NewReader(file)
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		if len(line) != 0 {
	// 			fmt.Println(line)
	// 		}
	// 		fmt.Println("文件读完了")
	// 		break
	// 	}

	// 	if err != nil {
	// 		fmt.Println("read file failed ,err", err)
	// 		return
	// 	}
	// 	fmt.Println(line)
	// }

	//************************ ioutil读取整个文件 *******************************
	//io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入。

	// content, err := ioutil.ReadFile("./xx.log")
	// if err != nil {
	// 	fmt.Println("open file failed ,err", err)
	// 	return
	// }
	// fmt.Println(string(content))

	//************************ 文件写入操作 *******************************

	// file, err := os.OpenFile("./aa.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	// if err != nil {
	// 	fmt.Println("open file failed ,err", err)
	// 	return
	// }

	// defer file.Close()
	// str := "hello world"
	// file.Write([]byte(str))          //写入字节切片数据
	// file.WriteString("\n hello sun") //直接写入字符串数据

	//************************ bufio.NewWriter *******************************
	//os.O_WRONLY 只写
	//os.O_CREATE 创建文件
	//os.O_RDONLY 只读
	//os.O_RDWR 读写
	//os.O_TRUNC 清空
	//os.O_APPEND 追加
	// perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。

	// file, err := os.OpenFile("bb.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	// if err != nil {
	// 	fmt.Println("open file failed ,err", err)
	// 	return
	// }
	// defer file.Close()

	// writer := bufio.NewWriter(file)
	// for i := 0; i < 10; i++ {
	// 	writer.WriteString("hello world\n") //将数据先写入缓存
	// }
	// writer.Flush() //将缓存中的内容写入文件

	//************************ ioutil.WriteFile *******************************
	// str := "hello world"
	// err := ioutil.WriteFile("cc.log", []byte(str), 0666)
	// if err != nil {
	// 	fmt.Println("write file failed ,err", err)
	// 	return
	// }

	_, err := copyFile("dst.txt", "xx.log")
	if err != nil {
		fmt.Println("copy file failed ,err", err)
		return
	}
	fmt.Println("copy done!")

}

// 借助io.Copy()实现一个拷贝文件函数。
func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("write %s failed ,err:%v.\n", srcName, err)
		return
	}
	defer src.Close()

	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("write %s failed ,err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //调用io.Copy()拷贝内容

}
