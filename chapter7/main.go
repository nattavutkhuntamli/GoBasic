package main

import "fmt"

func main() {
	//สร้าง array
	var numbers1 [3]int = [3]int{100,200,300}
	// numbers2 := [3]int{100,200,300,400} //error เพราะ ใน [3] เป็นการกำหนด array ไว้ว่าให้เก็บได้ไม่เกิน 3 ตัว
	numbers3 := [...]int {100,200,300,400} // เป็นการเพิ่มข้อมูลใน array ได้ไม่จำกัด ต่างจ่างnumber1 ที่กำจัดไว้ 3
	fmt.Println(numbers1)
	// fmt.Println(numbers2)
	fmt.Println(numbers1[0])
	numbers1[2] = 450
	fmt.Println(numbers1)

	size := len(numbers1) //นับขนาด array ว่ามีค่าเท่าไร คล้าย .length
	fmt.Println(size)

	fmt.Println(numbers3)
}