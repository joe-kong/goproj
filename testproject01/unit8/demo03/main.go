package main

import (
	"fmt"
)

func main() {
	//定义切片

	slice := make([]int, 4, 20)
	slice[0] = 66
	slice[1] = 88
	slice[2] = 99
	slice[3] = 100

	// 方式1: 普通for循环
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]的值是：%v \n", i, slice[i])
		fmt.Printf("slice[%v]的地址是：%v \n", i, &slice[i])

	}

	//方式2:for-range循环

	for i, v := range slice {
		fmt.Printf("下标为： %v 的值是：%v \n", i, v)
		fmt.Printf("下标为： %v 的地址是：%v \n", i, &slice[i])

	}

}
