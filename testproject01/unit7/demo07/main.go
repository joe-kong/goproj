package main

import "fmt"

//二维数组
func main(){
	var arr [2][3]int16
	fmt.Println(arr)

	fmt.Printf("arrのアドレスは：　%p\n", &arr)
	fmt.Printf("arr[0]のアドレスは：　%p \n", &arr[0])
	fmt.Printf("arr[0][0]のアドレスは：　%p\n", &arr[0][0])

	fmt.Printf("arr1のアドレスは：　%p\n", &arr[1])
	fmt.Printf("arr1[0]のアドレスは：　%p\n", &arr[1][0])
	fmt.Printf("arr1[1]のアドレスは：　%p\n", &arr[1][1])


	//fu zhi
	arr[0][1] = 47
	arr[0][0] = 82
	arr[1][1] = 25
	fmt.Println(arr)


	//init operation
	var arr1 [2][3]int = [2][3]int {{1,4,7},{2,5,8}}
	fmt.Println(arr1)
}
