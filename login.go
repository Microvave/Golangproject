package main

import "fmt"

var user = make(map[string]string)
var username string
var password string

func main() {

	for {
		var menu int

		// แสดงเมนู
		fmt.Println("กรุณาเลือกเมนู:")
		fmt.Println("1 = สมัครสมาชิก")
		fmt.Println("2 = เข้าสู่ระบบ")
		fmt.Println("3 = แสดงผู้ใช้ทั้งหมด")
		fmt.Println("4 = ออกจากโปรแกรม")

		// รับค่าเมนู
		fmt.Print("เลือกเมนู: ")
		fmt.Scan(&menu)
		switch menu {
		case 1:
			fmt.Print("กรอก username: ")
			fmt.Scan(&username)
			fmt.Print("กรอก password: ")
			fmt.Scan(&password)

			_, ok := user[username]

			if ok {

				fmt.Println("username นี้ถูกใช้งานแล้ว")
			} else {

				user[username] = password
				fmt.Println("สมัครสมาชิกสำเร็จ")
			}

		case 2:
			fmt.Print("กรอก username: ")
			fmt.Scan(&username)
			fmt.Print("กรอก password: ")
			fmt.Scan(&password)

			storedPassword, ok := user[username]

			if !ok {
				fmt.Println("ไม่พบผู้ใช้")
			} else {
				if password == storedPassword {
					fmt.Println("Login สำเร็จ")
				} else {
					fmt.Println("รหัสผ่านผิด")
				}
			}

		case 3:
			if len(user) == 0 {
				fmt.Println("ยังไม่มีผู้ใช้")
			} else {
				fmt.Println("รายชื่อผู้ใช้ทั้งหมด:")
				for username := range user {
					fmt.Println("-", username)
				}
			}

		case 4:
			fmt.Println("ลาก่อน")
			return
		default:
			fmt.Println("กรุณากรอกตัวเลข 1-4")
		}
	}
}
