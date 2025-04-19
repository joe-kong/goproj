package main

import "fmt"

// Person は人間を表す構造体。
type Person1 struct {
    Name string
    Age  int
}

func main2() {
    // ポインタ型の変数を宣言する
    // pがポインタ型変数
    // var p *Person //初期化しない状態宣言 nil
    // var p *Person = &Person{}

    // var p1 Person //DOの各項目がデフォルト値でObjectを初期化される
    // fmt.Printf("p1 :%+v\n", p1)
    // fmt.Printf("変数p1に格納されているアドレス :%p\n", &p1) //&を付けてアドレスを取得する
    //  aaaa(&p1)
    //  fmt.Printf("Afteraaaa p1 :%+v\n", p1)//aaaaメソッドで値を変えられる
    //  fmt.Printf("Afteraaaa変数p1に格納されているアドレス :%p\n", &p1)

    //  変数p1に格納されているアドレス :0xc000008078
    //  aa-p :&{Name:太郎太郎太郎太郎 Age:2020}
    //  aa-変数pに格納されているアドレス :0xc00000a030      
    // Afteraaaa p1 :{Name:太郎太郎太郎太郎 Age:2020}   ---------------変わった    
    // Afteraaaa変数p1に格納されているアドレス :0xc000008078


    // var p1 Person //DOの各項目がデフォルト値でObjectを初期化される
    // fmt.Printf("p1 :%+v\n", p1)
    // fmt.Printf("変数p1に格納されているアドレス :%p\n", &p1) //&を付けてアドレスを取得する
    //  aaaa(p1) //操作されたくない場合は、ポイント型として宣言する必要がありません。
    //  fmt.Printf("Afteraaaa p1 :%+v\n", p1)//aaaaメソッドで値を変えられる
    //  fmt.Printf("Afteraaaa変数p1に格納されているアドレス :%p\n", &p1)
    //  p1 :{Name: Age:0}
    //  変数p1に格納されているアドレス :0xc000008078
    //   aa-p :{Name:太郎太郎太郎太郎 Age:2020}
    //   aa-変数pに格納されているアドレス :0xc0000080a8
    //  Afteraaaa p1 :{Name: Age:0}---------------変わってない
    //  Afteraaaa変数p1に格納されているアドレス :0xc000008078


   
   
    //  var p2 *Person //ポインタ型変数　初期値が指定されていない場合、Go 言語のポインタは nil (nil ポインタ) に初期化されます。
    // fmt.Printf("p2222 :%+v\n", p2)
    // fmt.Printf("変数p2222に格納されているアドレス :%p\n", &p2) //&を付けてアドレスを取得する
    //  //aaaa(&p2)//nil ポインタだから、コンパイルエラーcannot use &p2 (value of type **Person) as *Person value in argument to aaaa
    //  p2= &Person{}
    //  fmt.Printf("After初期化 p2222 :%+v\n", p2)
    //  fmt.Printf("After初期化変数p2222に格納されているアドレス :%p\n", &p2)
    //  aaaa(p2)//初期化してからOKです。
    //  fmt.Printf("Afteraaaa p2222 :%+v\n", p2)//aaaaメソッドで値を変えられる
    //  fmt.Printf("Afteraaaa変数p2222に格納されているアドレス :%p\n", &p2)
//　------------------------------------------------------------------------------
    //  p2222 :<nil>
    //  変数p2222に格納されているアドレス :0xc00000a028
    //  After初期化 p2222 :&{Name: Age:0}　
    //  After初期化変数p2222に格納されているアドレス :0xc00000a028
    //   aa-p :&{Name:太郎太郎太郎太郎 Age:2020}
    //   aa-変数pに格納されているアドレス :0xc00000a038
    //  Afteraaaa p2222 :&{Name:太郎太郎太郎太郎 Age:2020}
    //  Afteraaaa変数p2222に格納されているアドレス :0xc00000a028
//　------------------------------------------------------------------------------

    var p3 *Person //ポインタ型変数　初期値が指定されていない場合、Go 言語のポインタは nil (nil ポインタ) に初期化されます。
    fmt.Printf("p2222 :%+v\n", p3)
    fmt.Printf("変数p2222に格納されているアドレス :%p\n", &p3) //&を付けてアドレスを取得する
    //  aaaa(p3)
     p3= &Person{}
     aaaa(*p3)
     fmt.Printf("Afteraaaa p3 :%+v\n", p3)//aaaaメソッドで値を変えられる
     fmt.Printf("Afteraaaa変数p3に格納されているアドレス :%p\n", &p3)

//　--------------------------コンパイルエラーがありませんが、実行するとエラー----------------------------------------------------
    //  p2222 :<nil>
    //  変数p2222に格納されているアドレス :0xc00000a028
    //  panic: runtime error: invalid memory address or nil pointer dereference   
    //  [signal 0xc0000005 code=0x1 addr=0x8 pc=0xaffc3e]
     
    //  goroutine 1 [running]:
    //  main.aaaa(0x0)
    //          D:/git/github/koureijun/goproj/testproject01/main/main.go:91 +0x5e
    //  main.main()
    //          D:/git/github/koureijun/goproj/testproject01/main/main.go:71 +0xd4
    //  exit status 2
//　------------------------------------------------------------------------------




    // var p *Person = &Person{}
    // fmt.Printf("p :%+v\n", p)
    // fmt.Printf("変数pに格納されているアドレス :%p\n", &p)
     
    //  aaaa(p)
    //  fmt.Printf("p :%+v\n", p)
    //  fmt.Printf("変数pに格納されているアドレス :%p\n", &p)
    // var p2 *Person 
    // fmt.Printf("p2 :%+v\n", p2)
    // fmt.Printf("変数pに格納されているアドレス :%p2\n", &p2)
}

func aaaa(p Person){ 
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