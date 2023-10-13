package main

import "fmt"

//数组的内存分析
func main(){
	//声明数组
	var arr[3] int16
	fmt.Printf("the len of arr: %v", len(arr))

	fmt.Println(arr)

	fmt.Printf("the address of arr : %p", &arr)
	fmt.Println("")
	fmt.Printf("the first element address of arr %p ", &arr[0])
	fmt.Printf("the 2nd element address of arr %p ", &arr[1])
	fmt.Printf("the 3rd element address of arr %p ", &arr[2])

// 数组每个空间占用的字节数取决与数组的类型
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Printf("2the first element address of arr %p ", &arr[0])
	fmt.Printf("2the 2nd element address of arr %p ", &arr[1])
	fmt.Printf("2the 3rd element address of arr %p ", &arr[2])
}