package main

import "fmt"

//切片类似与java中的list????
func main(){
	var intarr [6]int = [6]int {3,6,9,1,4,7}

	//切片构建在数组之上
	// 定义一个切片的名字slice，[]东爱变化的数组长度不写，int类型，intarr是原数组
	// [1:3]切片，切除一段片段 索引:从1开始，到3结束 不包含3 [1，3)
	// var slice []int = intarr[1:3] //是数组的一个片段的引用
	slice := intarr[3:5]

	fmt.Println("intearr : ",intarr)
	fmt.Println("slice : ",slice)
	fmt.Println("slice元素的个数 : ",len(slice))
	fmt.Println("slice元素的容量 : ",cap(slice))//容量动态变化 数组的长度-切片开始的位置

// 切片结构体包含三部分  底层数组的指针  切片的长度 切片的容量

	fmt.Printf("数组中下表为1weizhi的地址: %p", &intarr[1])
	fmt.Printf("slice中下表为0weizhi的地址: %p", &slice[0])
	//slice[1] = 16
	
	slice[0] = 16
	fmt.Println("intearr : ",intarr)
	fmt.Println("slice : ",slice)

	//定义切片的三种方式
	//从数组中切
	//通过make函数创建 var slicename [type = make([],len,[cap])]
	//定义一个切片，直接就指定具体数组，使用原理类似make的方式。
}