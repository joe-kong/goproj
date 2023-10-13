package main

import (
	"fmt"
	"strconv"
)

/*
String转为基本数据类型
*/
func main(){
	//string =-->bool
	var s1 string = "true"
	var b bool
	//ParseBool的函数返回值有两个 （value bool，err error）
	//value 就是我们得到的布尔类型的数据，err是出现的错误
	//我们只关注得到的布尔类型的数据，err可以用下划线直接忽略
	b,_=strconv.ParseBool(s1)
	fmt.Printf("b的类型是: %T,　b=%v \n",b,b)

	//string＝＝＝＞int64
	var s2 string  ="19"
	var num1 int64
	//num1,_= strconv.ParseInt(s2,0,0)
	num1,_= strconv.ParseInt(s2,0,10)
	fmt.Printf("num1のタイプは：　%T,num1 = %v \n",num1,num1)

	//string=>float32,float64
	var s3 string="3.14"
	var f1 float64
	f1,_=strconv.ParseFloat(s3,64)

	fmt.Printf("f1のタイプは：　%T,f1=%v \n",f1,f1)


}