package main

import "fmt"
func main() {
	Message("G")
	Message("A")
	Message("B")
	Message("C")

	resultSum,statusSum :=summation(10,20)
	fmt.Println("ยอดรวมทั้งหมด ", resultSum, statusSum)
}

func Message(name string) {
	fmt.Println("Hello Programming", name)
}

func summation (number1, number2 int) (int, string) {
	total  := number1 + number2
	status := ""
	if total % 2 != 0 {
		status ="เลขคู่"
	} else {	
		status ="เลขคี่"
	}
	return total,status
}