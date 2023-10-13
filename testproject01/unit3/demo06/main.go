package main

import "fmt"

func main(){
	// 实现功能‘键盘录入学生的年龄，姓名，成绩，是否ＶＩＰ

	var age int
	fmt.Println("请录入学生的年龄:")
	//传入ＡＧＥ的地址的目的‘在Ｓｃａｎｌｎ函数中给，对地址中的值进行改变的时候，实际外面的ａｇｅ被影响了
	// fmt.Scanln(&age)//录入数据的时候，类型一定要匹配，因为底层会自动判定类型

	var name string
	// fmt.Println("学生の姓名を入力してください：　")
	// fmt.Scanln(&name)

	var score float32
	// fmt.Println("学生の成績を入力してください：　")
	// fmt.Scanln(&score)

	var isVIP bool
	// fmt.Println("ISVIPを入力してください：　")
	// fmt.Scanln(&isVIP)

	//将上述数据在控制台下打印输出‘
	// fmt.Printf("学生的年龄为：　%v,姓名为　%v,成绩为　%v,是否为ＶＩＰ　%v,",age,name,score,isVIP)


	// 方式２:fmt.Scanf()
	fmt.Println("请录入学生的年龄，姓名，成绩，是否ＶＩＰ，使用空格进行分隔")
	fmt.Scanf("%d %s %f %t",&age,&name,&score,&isVIP)
	//将上述数据在控制台下打印输出‘
	fmt.Printf("学生的年龄为：　%v,姓名为　%v,成绩为　%v,是否为ＶＩＰ　%v,",age,name,score,isVIP)


}