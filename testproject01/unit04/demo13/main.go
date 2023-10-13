package main

import "fmt"

func main() {
//goto 一般不建议使用 容易造成程序混乱

fmt.Println("hello go1")
fmt.Println("hello go2")
if 1==1 {
	goto label1
}
fmt.Println("hello go3")
fmt.Println("hello go4")
fmt.Println("hello go5")
label1:
fmt.Println("hello go6")
fmt.Println("hello go7")

}