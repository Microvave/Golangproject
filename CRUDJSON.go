package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

var students []Student

func saveToFile(students []Student) error {
	data, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("students.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func loadFromFile() ([]Student, error) {
	data, err := os.ReadFile("students.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Student{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []Student{}, nil
	}

	var students []Student
	err = json.Unmarshal(data, &students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

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

func deleteStudent() {
	for {
		if len(students) == 0 {
			fmt.Println("ยังไม่มีนักเรียนในระบบ")
			return
		}

		var index int
		fmt.Print("กรอก index ที่ต้องการลบ (-1 เพื่อออก): ")
		fmt.Scan(&index)

		if index == -1 {
			return
		}

		if index < 0 || index >= len(students) {
			fmt.Println("index ไม่ถูกต้อง")
			continue
		}

		students = append(students[:index], students[index+1:]...)
		fmt.Println("ลบนักเรียนเรียบร้อย")
		return
	}
}

func updateStudent() {
	if len(students) == 0 {
		fmt.Println("ยังไม่มีนักเรียนในระบบ")
		return
	}

	for i, s := range students {
		fmt.Println(i, s.Name, s.Score)
	}

	var index int
	fmt.Print("เลือก index ที่ต้องการแก้ไข: ")
	fmt.Scan(&index)

	if index < 0 || index >= len(students) {
		fmt.Println("index ไม่ถูกต้อง")
		return
	}

	fmt.Println("กำลังแก้ไข:", students[index])

	var choice int
	fmt.Println("1. แก้ชื่อ")
	fmt.Println("2. แก้คะแนน")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("ชื่อใหม่: ")
		fmt.Scan(&students[index].Name)
	case 2:
		fmt.Print("คะแนนใหม่: ")
		fmt.Scan(&students[index].Score)
	default:
		fmt.Println("ตัวเลือกไม่ถูกต้อง")
	}
}

func main() {

	// โหลดไฟล์
	students, err := loadFromFile()
	if err != nil {
		fmt.Println("โหลดไฟล์ไม่สำเร็จ:", err)
		return
	}

	// fmt.Println("ข้อมูลจากไฟล์:")
	// for i, s := range students {
	// 	fmt.Println(i, s.Name, s.Score)
	// }

	// fmt.Println("บันทึกไฟล์สำเร็จ")

	// บันทึกไฟล์

	// err := saveToFile(students)
	// if err != nil {
	// 	fmt.Println("บันทึกไฟล์ไม่สำเร็จ:", err)
	// 	return
	// }

	for {
		fmt.Println("ข้อที่ 17 CRUD ด้วย JSON ")
		fmt.Println("กรุณาเลือกเมนู:")
		fmt.Println("1 = แสดงรายชื่อนักเรียนทั้งหมด")
		fmt.Println("2 = เพิ่มนักเรียน")
		fmt.Println("3 = ลบนักเรียน")
		fmt.Println("4 = แก้ไขนักเรียน")
		fmt.Println("5 = ออก")

		var choice int
		fmt.Print("เลือกเมนู: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if len(students) == 0 {
				fmt.Println("ไม่มีข้อมูล")
			} else {
				for i, s := range students {
					fmt.Println(i, s.Name, s.Score)
				}
			}
		case 2:
			addStudent()
			err := saveToFile(students)
			if err != nil {
				fmt.Println("บันทึกไฟล์ไม่สำเร็จ:", err)
				return
			}
			fmt.Println("บันทึกข้อมูลสำเร็จ")
		case 3:
			if len(students) == 0 {
				fmt.Println("ไม่มีข้อมูล")
			} else {
				for i, s := range students {
					fmt.Println(i, s.Name, s.Score)
				}
			}
			deleteStudent()
			err := saveToFile(students)
			if err != nil {
				fmt.Println("บันทึกไฟล์ไม่สำเร็จ:", err)
				return
			}
			fmt.Println("ลบข้อมูลสำเร็จ")
		case 4:
			updateStudent()
			err := saveToFile(students)
			if err != nil {
				fmt.Println("บันทึกไฟล์ไม่สำเร็จ:", err)
				return
			}
			fmt.Println("แก้ไขข้อมูลสำเร็จ")
		case 5:
			fmt.Println("ลาก่อน")
			return
		default:
			fmt.Println("เมนูไม่ถูกต้อง")
		}
	}

}
