package main

import "fmt"

func main() {
	//break 停止正在执行的循环
var sum int = 0
	for i:=0;i<100;i++{
		sum +=i 
	fmt.Println(sum)
	if sum > 300{
		break
	}
	}
}