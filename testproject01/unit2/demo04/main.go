package main

import "fmt"

func main(){
	//定义浮点类型的数据
	var num1 float32 = 3.14
	fmt.Println(num1)

	//可以表示整的浮点类型，也可以表示负的浮点类型
	var num2 float32 = -3.14
	fmt.Println(num2)

	//浮点数可以用十进制表示，也可以用科学计数法表示 E大写和小写都是可以的
	var num3 float32 = 314E-2
	fmt.Println(num3)

	var num4 float32 = 314E+2
	fmt.Println(num4)

	var num5 float32 = 314e+2
	fmt.Println(num5)

	var num6 float64 = 314e+2
	fmt.Println(num6)


	//浮点数通常会有精度的损失，所以通常情况下用float64 golang的默认浮点类型也是·float·64
	var num7 float32 = 256.000000916
    fmt.Println("num7:", num7)

	var num71 float32 = 256.916
    fmt.Println("num71:", num71)

	var num8 float64= 256.000000916
    fmt.Println("num8:", num8)

	//golang默认的浮点类型是·float64
	var num9= 3.17
	fmt.Printf("num9·对应的默认的数据类型为 %T",num9)




}