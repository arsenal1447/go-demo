package main

import (
	"fmt"
	"sync"
)

func main() {

	// //接下来我们在调用hello函数前面加上关键字go，也就是启动一个goroutine去执行hello这个函数。
	// //这一次的执行结果只打印了main goroutine done!，并没有打印Hello Goroutine!。为什么呢？
	// go hello()
	// fmt.Println("main goroutine done")

	// //所以我们要想办法让main函数等一等hello函数，最简单粗暴的方式就是time.Sleep了。
	// //执行上面的代码你会发现，这一次先打印main goroutine done!，然后紧接着打印 Hello Goroutine!。
	// //首先为什么会先打印main goroutine done!是因为我们在创建新的goroutine的时候需要花费一些时间，而此时main函数所在的goroutine是继续执行的。
	// time.Sleep(time.Second)

	// 在Go语言中实现并发就是这样简单，我们还可以启动多个goroutine。让我们再来一个例子： （这里使用了sync.WaitGroup来实现goroutine的同步）
	//Hello goroutine 9
	// Hello goroutine 5
	// Hello goroutine 4
	// Hello goroutine 7
	// Hello goroutine 6
	// Hello goroutine 1
	// Hello goroutine 2
	// Hello goroutine 0
	// Hello goroutine 3
	// Hello goroutine 8
	//多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello1(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束

}

func hello() {
	fmt.Println("Hello goroutine")
}

//在Go语言中实现并发就是这样简单，我们还可以启动多个goroutine。让我们再来一个例子： （这里使用了sync.WaitGroup来实现goroutine的同步）
var wg sync.WaitGroup

func hello1(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello goroutine", i)
}
