package main

import "fmt"

func main() {
//双重循环
label2://通过标签定义 来结束指定标签对应的循环 尤其是外部循环
for i := 1; i <= 5; i++ {	
	for  j:= 2; j <=4; j++ {

	if i == 2 && j == 2 {
		// continue
		continue label2
		}
		fmt.Printf("i: %v, j:%v \n",i,j)	
	}
}
fmt.Println("...........")
}