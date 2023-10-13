package main

import "fmt"

func main() {
// 输出1-100 所有被6整除的数字
//continue 结束本次循环，但是继续下一次循环。
	for i := 0; i <= 100; i++ {
		if i % 6 != 0 {
			continue
		}
		fmt.Println(i)
	}
}