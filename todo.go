package main

import "fmt"

func main() {

	var WorkName []string
	var name string

	for {
		var menu int

		// แสดงเมนู
		fmt.Println("กรุณาเลือกเมนู:")
		fmt.Println("1 = เพิ่มงาน")
		fmt.Println("2 = แสดงงานทั้งหมด")
		fmt.Println("3 = ลบงานตามหมายเลข")
		fmt.Println("4 = ออกจากโปรแกรม")

		// รับค่าเมนู
		fmt.Print("เลือกเมนู: ")
		fmt.Scan(&menu)
		switch menu {
		case 1:
			fmt.Scan(&name)
			WorkName = append(WorkName, name)

		case 2:
			if len(WorkName) == 0 {
				fmt.Println("ยังไม่มีงาน")
			} else {
				fmt.Println("รายการงานทั้งหมด:")
				for i, work := range WorkName {
					fmt.Printf("%d. %s\n", i+1, work)
				}
			}

		case 3:
			if len(WorkName) == 0 {
				fmt.Println("ยังไม่มีงานให้ลบ")
				break
			}

			var index int
			fmt.Print("กรอกหมายเลขงานที่ต้องการลบ: ")
			fmt.Scan(&index)

			if index < 1 || index > len(WorkName) {
				fmt.Println("หมายเลขไม่ถูกต้อง")
			} else {
				WorkName = append(WorkName[:index-1], WorkName[index:]...)
				fmt.Println("ลบงานเรียบร้อย")
			}

		case 4:
			fmt.Println("ลาก่อน")
			return
		default:
			fmt.Println("กรุณากรอกตัวเลข 1-4")
		}
	}
}
