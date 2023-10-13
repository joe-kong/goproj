package main

import "fmt"

func main(){

	//实现功能　如果口罩的库存少于三十个，提示　库存不足
	//定义口罩数量
	var count int = 70

	if count<30 {
		fmt.Print("库存不足")
	}else{
		fmt.Print("库存足")
	}
	//双分支一定会二选一走其中一个分支
}