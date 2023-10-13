package main

import "fmt"

//数组的引入
func main(){
	score1 :=95
	score2 :=91
	score3 :=39
	score4 :=60
	score5 :=21

	sum := score1 + score2 + score3 + score4 + score5
	avg := sum/5

	fmt.Printf("成绩的总合为: %v，成绩的平均数为 %v", sum,avg)
}