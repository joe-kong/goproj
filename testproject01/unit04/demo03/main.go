package main

import "fmt"

func main(){
	//实现功能 根据给出的学生的分数，判断学生的等级
	// >=90 A >=80 B  >=70 C >=60 D  <60 E

	// //方式一 利用if单分支实现
	var score int = 76
	// if score >=90{
	// 	fmt.Println("level A")
	// }
	// if score >=80 && score <90 {
	// 	fmt.Println("level B")
	// }
	// //...
	// if score <60{
	// 	fmt.Println("level E")
	// }


	//多分支实现 优点 走了一个分支 其他的分支就不会再去判断执行
	if score >=90{
		fmt.Println("level A")
	}else if score>=80{
		fmt.Println("level B")
	}else if score >=70{
		fmt.Println("level C")
	}else if score >=60{
		fmt.Println("level D")
	}else if score <60{
		fmt.Println("level E")
	}

}