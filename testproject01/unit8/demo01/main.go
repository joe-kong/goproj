package main

import "fmt"

// 切片类似与java中的list????
func main() {
	//定义一个数组
	//var intarr [6]int = [6]int{3, 6, 9, 1, 4, 7}
	intarr := [6]int{3, 7, 9, 1, 4, 7}
	fmt.Println("sintarr元素的个数 : ", len(intarr))
	fmt.Println("intarr元素的容量 : ", cap(intarr))
	// var intarr3 [6]int = [6]int{1,2,3,4,5,6}
	//切片构建在数组之上，是数组切的一个片段。
	// 定义一个切片的名字slice，[]动态变化的数组长度不写，int类型，intarr是原数组
	//var slice []int = intarr[1:3]
	// [1:3]切片，切除一段片段 索引:从1开始，到3结束 不包含3 [1，3)
	// var slice []int = intarr[1:3] //是数组的一个片段的引用
	slice := intarr[3:5]

	fmt.Println("intearr : ", intarr)
	fmt.Println("slice : ", slice)
	fmt.Println("slice元素的个数 : ", len(slice))
	fmt.Println("slice元素的容量 : ", cap(slice)) //容量动态变化 数组的长度-切片开始的位置

	// 切片结构体包含三部分  底层数组的指针  切片的长度 切片的容量
	fmt.Printf("数组中下标为1位置的地址: %p\n", &intarr[1])
	fmt.Printf("数组中下标为1位置的值: %d\n", intarr[1])
	fmt.Printf("数组中下标为3位置的地址: %p\n", &intarr[3])

	fmt.Printf("slice中下标为0位置的地址: %p\n", &slice[0])
	fmt.Println()

	//slice[1] = 16

	slice[0] = 16
	fmt.Println("intearr : ", intarr)
	fmt.Println("slice : ", slice)

	//定义切片的三种方式
	//从数组中切
	//通过make函数创建 var slicename [type = make([],len,[cap])]
	//定义一个切片，直接就指定具体数组，使用原理类似make的方式。
}
