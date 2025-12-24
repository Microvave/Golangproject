package main

import "fmt"

// 	สร้างระบบจัดการนักเรียนที่สามารถ:

// เพิ่มนักเรียน

// แสดงรายชื่อนักเรียนทั้งหมด

// ลบนักเรียน

// ใช้ struct + slice

type Student struct {
	Name  string
	Score float64
}

var students []Student

func addStudent() {
	// รับชื่อ + คะแนน
	// append เข้า students
	var Name string
	var Score float64
	fmt.Print("กรอกชื่อนักเรียน: ")
	fmt.Scan(&Name)
	fmt.Print("กรอกคะแนน : ")
	fmt.Scan(&Score)

	students = append(students, Student{Name: Name, Score: Score})
	fmt.Println("เพิ่มนักเรียนเรียบร้อยแล้ว")
}

func listStudents() {
	// แสดงนักเรียนทั้งหมด
	for i, student := range students {
		fmt.Printf("Index %d: %s - %.2f\n", i, student.Name, student.Score)
	}
}

func deleteStudent() {
	if len(students) == 0 {
		fmt.Println("ยังไม่มีนักเรียนในระบบ")
		return
	}

	var index int
	fmt.Print("กรอก index ที่ต้องการลบ: ")
	fmt.Scan(&index)

	if index < 0 || index >= len(students) {
		fmt.Println("index ไม่ถูกต้อง")
		return
	}

	students = append(students[:index], students[index+1:]...)
	fmt.Println("ลบนักเรียนเรียบร้อย")
}

func main() {
	for {
		fmt.Println("ข้อที่ 14 โปรแกรมจัดการนักเรียน")
		fmt.Println("กรุณาเลือกเมนู:")
		fmt.Println("1 = add")
		fmt.Println("2 = list")
		fmt.Println("3 = delete")
		fmt.Println("4 = exit")

		var choice int
		fmt.Print("เลือกเมนู: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addStudent()
		case 2:
			listStudents()
		case 3:
			deleteStudent()
		case 4:
			fmt.Println("ลาก่อน")
			return
		default:
			fmt.Println("เมนูไม่ถูกต้อง")
		}
	}

}
