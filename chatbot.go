package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		// รับข้อความ
		// ตรวจคำสั่ง
		var text string
		fmt.Print("กรุณากรอกข้อความ: ")
		fmt.Scanln(&text)
		if text == "hello" {
			fmt.Println("HI ")
		} else if text == "time" {
			fmt.Println(time.Now())
		} else if text == "exit" {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("ไม่เข้าใจคำสั่ง")
		}
		return

	}
}
