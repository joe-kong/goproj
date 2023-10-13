package main

import "fmt"

//数组的初始化方式
func main(){
	
	// 1st
	var arr1 [3] int = [3] int {3,6,9}
	fmt.Println(arr1)//[3 6 9]

	// 2nd
	var arr2 = [3]int{1,4,7}
	fmt.Println(arr2)//[1 4 7]   

	// 3rd
	var arr3 = [...]int{4,5,6,7}
	fmt.Println(arr3)//[4 5 6 7]  

	// 4
	var arr4 = [...]int{2:66,0:33,1:99,3:88}
	fmt.Println(arr4)//[33 99 66 88]


}