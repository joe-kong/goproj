package main

import "fmt"

// Person は人間を表す構造体。
type Person struct {
    Name string
    Age  int
}

func main() {
    var p3 *Person //ポインタ型変数　初期値が指定されていない場合、Go 言語のポインタは nil (nil ポインタ) に初期化されます。
    fmt.Printf("p2222 :%+v\n", p3)
    fmt.Printf("変数p2222に格納されているアドレス :%p\n", &p3) //&を付けてアドレスを取得する
    //  aaaa(p3)
     p3= &Person{}
     aaaa2(*p3)
     fmt.Printf("Afteraaaa p3 :%+v\n", p3)//aaaaメソッドで値を変えられる
     fmt.Printf("Afteraaaa変数p3に格納されているアドレス :%p\n", &p3)

}

func aaaa2(p Person){ 
   p.Name="太郎太郎太郎太郎"
    p.Age = 2020
    
    fmt.Printf(" aa-p :%+v\n", p)
    fmt.Printf(" aa-変数pに格納されているアドレス :%p\n", &p)
}

 // var p *Person = &Person{}
    // p = Person{
    //     Name: "太郎",
    //     Age:  20,
    // }

    // p.Name="太郎"
    // p.Age = 20

	// p := Person{
    //     Name: "太郎",
    //     Age:  20,
    // }
    // aa(&p)
    // fmt.Printf("p :%+v\n", p)
    // fmt.Printf("変数pに格納されているアドレス :%p\n", &p)

//     p2 := p
//    p2.Age = 30
//    p2.Name= "aaa"

//     fmt.Printf("p :%+v\n", p)
//     fmt.Printf("変数pに格納されているアドレス :%p", &p)