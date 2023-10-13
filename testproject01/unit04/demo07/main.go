package main

import "fmt"

func main() {
	var str string = "hello golang 你好"

	// //1 普通for循环
	// for i:=0; i<= len(str)-1; i++{
	// 	fmt.Printf("%c \n",str[i])
	// }

	//2 for range
	for i, value :=range str{
		fmt.Printf("index is %d ,detail value is %c \n",i ,value)
	}
	// for range 结构是GO语言特有的一种迭代结构，在许多情况下都非常有用，forrange 可以遍历数组，切片，字符串，map以及通道，for range语法上类似于其他语言中的foreach语句，一般形式为
	// for key;val:=range coll{...}
}