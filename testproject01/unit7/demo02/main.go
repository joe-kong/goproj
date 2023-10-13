package main

import "fmt"

func main(){
	// score1 :=95
	// score2 :=91
	// score3 :=39
	// score4 :=60
	// score5 :=21

	// sum := score1 + score2 + score3 + score4 + score5
	// avg := sum/5

	// fmt.Printf("成绩的总合为: %v，成绩的平均数为 %v", sum,avg)

	var score[5] int

	score[0] = 95
	score[1] = 91
	score[2] = 39
	score[3] = 60
	score[4] = 21
	
	sum :=0
	for i := 0; i < len(score) ; i++ {
		sum += score[i]
	}
	avg := sum / len(score)

	fmt.Printf("成绩的总合为 %v，成绩的平均数为 %v", sum,avg)

}