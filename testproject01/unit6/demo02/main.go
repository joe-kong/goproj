package main

import (
	"errors"
	"fmt"
)

func main(){
	err := test()

	if err != nil {
		fmt.Println("自定义错误: err:",err)
		panic(err)//如果不想继续执行后面的程序，可以使用builtin包下面的panic函数
	}

	fmt.Println("上面的出发操作执行成功。。。")
	fmt.Println("正常执行下面的逻辑。。。")

}


func test() (err error){
	num1 := 10;
	// num2 := 1
	num2 := 0

	if num2 == 0 {
		// 借助Errors抛出自定义的错误
		return errors.New("除数不能为零哦")
	}else{
		
		result := num1/num2
		fmt.Println(result)
		//如果没有出错，返回零值。 
		return nil

	}
}