package main

import "fmt"

//数组的遍历
func main(){
	//声明数组
	var scores[5] int
	 for i := 0; i < len(scores); i++ {
		fmt.Printf("请录入第%d个学生的成绩:", i+1)
		fmt.Scanln(&scores[i])//读取录入的值给数组
	 }
	 sum :=0
	 for i := 0; i < len(scores) ; i++ {
		 sum += scores[i]
	 }
	 avg := sum / len(scores)
 
	 fmt.Printf("成绩的总合为 %v，成绩的平均数为 %v\n", sum,avg)

	 // show every students socre
	 //method1 comman for loop

	//  for i := 0; i < len(scores); i++ {
	// 	fmt.Printf("第%d个同学的成绩:  %d \n",i+1,scores[i])
	//  }

	 //method2: for range键值循环
	 //（键值循环）for range 结构是Go语言特有的一种迭代结构，在许多情况下都非常有用，for range可以遍历数组，切片，字符串，map和通道。for range语法上类似其他语言的foreach语句，
	 // 一般形式 for key,val := range coll{....} 
	// coll就是你要遍历的对象，每次遍历得到的索引用key接受，每次遍历得到的索引位置上的值用val keyvalue的名字随便起名 k v key value 都行

	for key,value := range scores{
		fmt.Printf("第%d个同学的成绩:  %d \n",key+1,value)
	}

	//Key值不想接受咋办 像忽略某个值用下划线_·
	for _,value := range scores{
		fmt.Printf("同学的成绩:  %d \n",value)
	}
}