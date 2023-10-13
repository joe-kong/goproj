package main

import "fmt"

func main() {
	var sum int = 0
	//for的初始表达式，不能用var定义变量的形式，要用:=
	//for循环世界上让程序员写代码的效率提升了，但是底层的执行逻辑没有变。
	for i := 1; i<=5; i++ {
		sum += 1
	}

	fmt.Println(sum)


	//j在for之外定义也是可以的。
	j :=1
	for ;j<=5; j++ {
		fmt.Println(j)
	}

	//死循环
	for{
		fmt.Println("dead roop running")
	}
}