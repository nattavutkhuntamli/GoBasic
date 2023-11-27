package main

import "fmt"

func main() {
	// numbers := []int{10,20,30,40}
	language := map[string]string {"TH":"Thai","En":"Englist"}
	
	// for i:=0; i<= len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	// for index, value := range numbers {
	// 	fmt.Println("index =", index , "Value =", value)
	// }
	// for _, value := range numbers {
	// 	fmt.Println( "Value =", value)
	// }

	for key, value := range language {
		fmt.Println( "key",key ,"Value =", value)
	}
}

