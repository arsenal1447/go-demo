package main

import "fmt"

// 求数组[1, 3, 5, 7, 8]所有元素的和
func sum(array [5]int) int {
	sum := 0
	for _, v := range array {
		sum = sum + v
	}
	return sum
}

func main() {
	var array = [5]int{1, 3, 5, 7, 8}
	fmt.Println("sum(array):", sum(array))
	// fmt.Println("find(array):", subscript(array))

	var testArr2 = [5]int{2, 7, 4, 1, 6}
	find(testArr2, 8)
}

// 找出数组中和为指定值的两个元素的下标，
// 比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
// func find(array [5]int) int {
// 	// var array = [5]int{1, 3, 5, 7, 8}
// 	length := len(array)
// 	const num = 8

// 	for i := 0; i < length-1; i++ {
// 		for j := i + 1; j < length; j++ {
// 			if array[i]+array[j] == num {
// 				fmt.Println("(%d, %d)\n", i, j)
// 				return array[i]
// 			}
// 		}
// 	}

// }

func find(arr [5]int, sum int) {
	for i, v := range arr {
		for n, val := range arr {
			if i == n {
				break
			}

			if v+val == sum {
				fmt.Println("元素相加和为", sum, "的元素下标有:{index:", i, "value:", v, ",index:", n, "value:", val, "}")
			}
		}
	}
}
