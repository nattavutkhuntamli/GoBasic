package main 

import (
	"fmt"
)

func main () {
	var name string 
	var score int
	var grade string
	fmt.Println("ป้อนชื่อนักเรียน =" ,name)
	fmt.Scanf("%s", &name)
	
	fmt.Println("ป้อนคะแนนเกรด" , score)
	fmt.Scanf("%d", &score)

	if score > 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else if score >= 50 {
		grade = "D"
	} else if score >= 40 {
		grade = "E"
	} else {
		grade = "F"
	}

	fmt.Println("นักเรียน" ,name)
	fmt.Println("ได้คะแนน", score ," เกรดที่ได้ " , grade)
}
