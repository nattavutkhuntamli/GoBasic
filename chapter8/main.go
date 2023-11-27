package main
import "fmt"
func main() {
	// slices
	/*
	   มีลักษณะคล้าย Array แต่มีความสามารถในการปรับเปลี่ยนขนาดสมาชิกได้ 
	   Dynamic size
	*/
	 numbers := []int{100,200,300,400,500,600,700} //มีค่าเริ่มต้น
	 fmt.Print(numbers[0])
	 fmt.Print(numbers)
	 numbers[0] = 505
	 fmt.Print(numbers)
	 
	 //เพิ่มค่าลงใน numbers ใน array โดยใช้ append
	 numbers = append(numbers,800)
	 numbers = append(numbers,890)
	 fmt.Println(numbers)
	 
	 fmt.Println(numbers[:]) //ดึงข้อมูลมาหมดตั้งแต่เริ่มต้นไปยันตั่วท้าย
	 fmt.Println(numbers[1:]) //index 1 -สุดท้าย
	 fmt.Println(numbers[1:3]) // index 1 -2
	 fmt.Println(numbers[:3]) //ดึง ตั้ง index เริ่มต้นจนถึง index ที่ 3
	 
	 count  := len(numbers)
	 fmt.Println(count)

	//  maps ตัวแปรที่เก็บข้อมูลในรูปแบบตัวแปร key value มีลักษณะคล้าย
	// อาร์เรย์ แต่จะใช้ key เป็น index เพื่อเชื่อมโยงข้อมูล value ที่เก็บใน
	// key นั้น ถ้าทราบ key ก็สามารถเข้าถึง value หรือข้อมูลนั้นได้

	country := [2]string{"Thailand","England"}
	fmt.Println("country length = ", len(country), " response :" ,country)

	coin := map[string]string {}
	coin["Thailand"] = "66"
	coin["England"]  = "55"
	fmt.Println(coin)

	/*
	  break;

	  countinue
	*/
	for i := 0; i<=10; i++ {
		if  i == 5 {
			break;
		}
	}
	fmt.Println("หยุดทำงาน")

	for r := 0; r<=10; r++ {
		if r == 5{
			continue
		}
	}
	fmt.Print("จบ")
}

func countLoop() {
	for i:=0; i<=10; i--{ 
		for r :=0; r<=i; r++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}