package main

import "fmt"

//数组的注意事项
func main(){
	
	// 1st
	var arr1 [3] int = [3] int {3,6,9}
	fmt.Println(arr1)//[3 6 9]
    fmt.Printf("数组的类型为: %T\n",arr1)

	var arr2 [6] int = [6] int {3,6,9,1,4,7}
	fmt.Println(arr2)//[3 6 9]
    fmt.Printf("数组的类型为: %T \n",arr2)

	var arr3 = [3]int{3,6,9}
	//test(arr3)
	test(&arr3)//指针传递 传入arr3的地址
	fmt.Printf("arr3: %v", arr3)//[3 6 9] 而不是[7 6 9 ]

}

// func test(arr [3]int){
// 	//这个数组实在新的栈中，另外开辟一个新的内存空间，不会更新原来的参数的值。
// 	// 如果想在其他函数中，去修改原来的数组，可以使用引用传递（指针方式）
// 	arr[0]  =7
// }

func test(arr *[3]int){
	//这个数组实在新的栈中，另外开辟一个新的内存空间，不会更新原来的参数的值。
	// 如果想在其他函数中，去修改原来的数组，可以使用引用传递（指针方式） 如下
	//arr[0]  = 7
	 (*arr)[0]  =7

}