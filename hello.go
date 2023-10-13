package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("hello \rworld!")

	var aa string= "aaaa"  
	fmt.Println(aa)

	var s1 string ="true"
	var b bool
	//返回值有两个 value bool，err error
	b,_=strconv.ParseBool(s1)
	fmt.Printf("type of b is: %T\n,b=%v\n",b,b)

	//string to int64
	var s2 string ="19"
	var num1 int64
	num1,_= strconv.ParseInt(s2,10,64)
	fmt.Printf("num1のタイプは %T,num1=%v \n",num1,num1)
}

