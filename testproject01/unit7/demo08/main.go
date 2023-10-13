package main

import "fmt"

//2维数组的遍历
func main(){
	// define

	var arr [3][3]int = [3][3]int{{1,4,7},{2,5,8},{3,6,9}}
	fmt.Println(arr)

	fmt.Println("------------------------------")

	//method1 common for loop
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j],"\t")
		}
		fmt.Println()
	}
	fmt.Println("------------------------------")


	// method2 for-range

	for key,value := range arr {
		for k,v := range value{
			fmt.Printf("arr[%v][%v]の値： %v \t\n",key,k,v)
		}
	}
}

