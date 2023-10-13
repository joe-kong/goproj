package main

import "fmt"

//单分支
func main(){
	//实现功能　如果口罩的库存少于三十个，提示　库存不足
	var count int = 10
	//单分支
	if count<30 {
		fmt.Println("sorry 口罩库存不足")
	}

	//if后面的表达式，返回结果一定是true或者是false
	//如果返回结果为true，那么大括号的代码就会执行，如果是false那么大括号的代码就不执行
	//if后面一定要有空格，和条件表达式分割开来，大括号一定不能省略

	//条件表达式左右的小括号是建议省略的 在golang里面，if后面可以并列的加入变量的定义。
	if count1:=20;count1<30 {
		fmt.Println("sorry 口罩库存不足.....")
	}

}