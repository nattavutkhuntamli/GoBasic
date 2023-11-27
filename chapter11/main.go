package main

import (
	"fmt"
	"chapter11/calculator"
)

// การนิยามสครัคเจอร์
type Product struct {
	name string
	price float64
	category string
	discount int
}
func main() {
   Result := calculator.Summation(10,20,30)
   fmt.Println(Result)
   product1 := Product{name:"ปากกา", price:10, category:"เครื่องเขียน", discount:10}
   fmt.Println(product1)
   product1.price =5 
   fmt.Println(product1)
}

