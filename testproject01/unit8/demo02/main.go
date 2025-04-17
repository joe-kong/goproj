package main

import "fmt"

// 切片类似与java中的list????
// 定义切片的三种方式
// 从数组中切
// 通过make函数创建 var slicename []type = make([]type,len,[cap])] 1，切片的类型，2，长度 3，容量
// 定义一个切片，直接就指定具体数组，使用原理类似make的方式。
func main() {
	slice := make([]int, 4, 20) //slice :  [0 0 0 0]
	fmt.Println("slice : ", slice)

	fmt.Println("slice元素的个数 : ", len(slice))
	fmt.Println("slice元素的容量 : ", cap(slice))

	slice[0] = 66
	slice[1] = 88
	slice[3] = 99
	//slice[4] = 99 //panic: runtime error: index out of range [4] with length 4

	fmt.Println("slice : ", slice)
	//make底层创建一个数组，对外不可见，所以不可以直接操作这个数组，要通过slice间接的访问各个元素，不可以直接对数组进行维护和操作

	//3，定义一个切片，直接就指定具体的数组，使用原理类似make

	slice2 := []int{1, 4, 7}
	fmt.Println("slice2 : ", slice2)
	fmt.Println("slice2元素的个数 : ", len(slice2))
	fmt.Println("slice2元素的容量 : ", cap(slice2))

}
