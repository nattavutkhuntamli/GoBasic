package main

import "fmt"

func main() {
	var choie int
	fmt.Println("กรุณาเลือกรายการต่อไปนี้")
	fmt.Println("1 เปิดบัญชี")
	fmt.Println("2 ฝาก - ถอน")
	fmt.Println("เลือกรายการ " )
	fmt.Scanf("%d", &choie)

	switch choie {
		case 1:fmt.Println("เปิดบัญชี") 
		case 2:fmt.Println("ฝากเงิน - ถอนเงิน") 
	    default: fmt.Println("ออกจากระบบ")
		
	}

	printTriangle(10)
	printTrianglet(10)
}

func printTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func printTrianglet(n int) {
	for i := n; i>=1; i-- {
		for j := 1; j<=i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	} 
}