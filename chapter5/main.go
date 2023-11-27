package main

import "fmt"
func main() {
	var score int;
	fmt.Println("ป้อนคะแนน =")
	fmt.Scanf("%d", &score)
	if score >= 80 {
	 fmt.Println("A")
	} else if score >=70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else if score >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("f")
	}
}