package main

import "fmt"

//全局变量
var n7=100
var n8=9.8


//设计者认为全局变量可以一次性的声明
var(
	n9=500
	n10="netty"
)

func main(){
	//第一种 变量的使用方式‘指定变量的类型 并且赋值
	var num int =15
	fmt.Println(num)
	//第2种 变量的使用方式‘指定变量的类型 但不赋值 使用默认值
	var num2 int 
	fmt.Println(num2)
   
	 //第三种 如果没有写变量的类型 那么根据等号后面的值进行判断变量的类型*自动类型推断
	 var num3="tom"
	 fmt.Println(num3)

	 //第四种 省略var 用:=
	 sex :="男"
	 fmt.Println(sex)

	 fmt.Println("-----------------------------------------")
	 var n1,n2,n3 int
	 fmt.Println(n1)
	 fmt.Println(n2)
	 fmt.Println(n3)

	 var n4,name,n5=10,"jack",7.8 
	 fmt.Println(n4)
	 fmt.Println(name)
	 fmt.Println(n5)

	 n6,height:=6.9,100.6
	 fmt.Println(n6)
	 fmt.Println(height)

	fmt.Println("n7:",n7)
	fmt.Println("n8:",n8)
	fmt.Println("n9:",n9)
	fmt.Println("n10:",n10)
}