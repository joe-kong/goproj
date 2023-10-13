package main

import "fmt"

// recover:程序中出现错误后，程序被中断，无法继续执行

//错误处理和捕获机制

//go语言吹球代码优雅，引入了defer+recover机制来处理错误
//两个都是内置函数

func main() {
test()
fmt.Println("上面的错误正常执行。。")
fmt.Println("开始执行下面的程序")
}

func test() {
	//利用defer和Recover机制来捕获错误
	defer func ()  {
		// 调用recover的内置函数，可以捕获错误
		err := recover()
		// 如果没有错误，返回值为零值
		if err != nil{
			fmt.Println("错误已经处理")
			fmt.Println("Err是:",err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}