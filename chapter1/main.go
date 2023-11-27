package main


//import packages ที่จำเป็นสำหรับการใช้งาน JSON และการทำงานกับ HTTP ใน Go
import  (
	"encoding/json"
	"net/http"
)
// ประกาศ struct ชื่อ GradeResponse ที่มีฟิลด์ชื่อ Grade แบบ string ซึ่งระบุการแปลงเป็น JSON ด้วย tag json:"grade"
type GradeResponse struct {
	Grade string `json:"grade"`
}

/*
main() เป็นฟังก์ชันหลักที่เริ่มต้นโปรแกรม ที่นี่เรียกใช้ http.HandleFunc
เพื่อที่จะเปิด endpoint /getGrade และสร้าง handler โดยใช้ getGradeHandle
และ http.ListenAndServe เพื่อรันเซิร์ฟเวอร์ที่ port 8080 โดยใช้ nil เพื่อใช้ค่าเริ่มต้นของ HTTP server
*/
func main() {
	http.HandleFunc("/getGrade", getGradeHandler)
	http.ListenAndServe(":8080", nil)
}

/*
getGradeHandler คือ handler ที่จัดการ request 
ที่เข้ามาที่ /getGrade โดยสร้าง GradeResponse จากการเรียก getGrade 
เพื่อรับค่าเกรด และใช้ json.NewEncoder เพื่อ encode 
ข้อมูลเป็น JSON แล้วส่งผลลัพธ์กลับไปยัง client ผ่าน http.ResponseWriter
*/
func getGradeHandler(w http.ResponseWriter, r *http.Request) {
	scrole := 61
	grade := getGrade(scrole)

	response := GradeResponse{Grade: grade}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

/*
getGradeHandler คือ handler ที่จัดการ request ที่เข้ามาที่ /getGrade โดยสร้าง GradeResponse 
จากการเรียก getGrade เพื่อรับค่าเกรด และใช้ json.NewEncoder เพื่อ encode ข้อมูลเป็น JSON แล้วส่งผลลัพธ์กลับไปยัง 
client ผ่าน http.ResponseWriter
*/
func getGrade(scrole int) string {
	if scrole > 80 {
		return "A"
	} else if scrole >= 70 {
		return "B"
	} else if scrole >= 60 {
		return "C"
	} else if scrole >= 50 {
		return "D"
	} else if scrole >= 40 {
		return "E"
	}
	return "F"
}