package main

import "fmt"



func main () {
    number1 , number2 := 20 ,10;

	fmt.Println("ผลบวก " ,number1, "+" , number2, " =",  number1 + number2)
	fmt.Println("ผลลบ "  ,number1, "-" , number2, " =",  number1 - number2)
	fmt.Println("ผลคูณ "  ,number1, "*" , number2, " =",  number1 * number2)
	fmt.Println("ผลหาร "  ,number1, "/" , number2, " =",  number1 / number2)
	fmt.Println("ผลหารเอาเศษ "  ,number1, "%" , number2, " =",  number1 % number2)

	fmt.Println(" " ,number1, "เท่ากันหรือไม่" , number2, " =",  number1 == number2)
	fmt.Println(" " ,number1, "ไม่เท่ากันหรือไม่" , number2, " =",  number1 != number2)
	fmt.Println(number1, "มากกว่า" , number2, "=",  number1 > number2)
	fmt.Println(number1, "น้อยกว่า" , number2, "=",  number1 < number2)
	fmt.Println(number1, "มากกว่าเท่ากับ" , number2, "=",  number1 >= number2)
	fmt.Println(number1, "น้อยกว่าเท่ากับ" , number2, "=",  number1 <= number2)
}