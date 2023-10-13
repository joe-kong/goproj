package main

import "fmt"

func main() {

	var score int = 75
	fmt.Println(score)
	//switch score/10 {
	switch 6 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	case 5:
		fmt.Println("E")
	case 4 ,3 ,1 ,2 ,0:
		fmt.Println("E")
	default://可以放在任何位置，不一定在最后
		fmt.Println("???")
	}

	//switch后面是一个表达式 这个表达式的结果一次跟case进行比较，满足结果后返回对应的处理
	//default是用来兜底的 其他case都不能走的情况下就会走default分支
	
	//fallthrough关键字 可以执行穿透  只能到下一个分支 不会执行所有的剩余分支

}