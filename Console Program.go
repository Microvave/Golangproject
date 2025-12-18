package main

import "fmt"

// oldMeter (มิเตอร์เก่า)
// newMeter (มิเตอร์ใหม่)
// pricePerUnit (ราคาต่อหน่วย)
// units = newMeter - oldMeter
// totalPrice = units * pricePerUnit

var oldMeter float64

var newMeter float64

var pricePerUnit float64
var units float64
var totalPrice float64

func main() {
	fmt.Println("ข้อที่ 13 โปรแกรมคำนวณค่าน้ำค่าไฟ")
	fmt.Println("กรุณากรอกข้อมูลมิเตอร์เก่า")
	fmt.Scan(&oldMeter)
	fmt.Println("กรุณากรอกข้อมูลมิเตอร์ใหม่")
	fmt.Scan(&newMeter)
	fmt.Println("กรุณากรอกราคาต่อหน่วย")
	fmt.Scan(&pricePerUnit)
	if newMeter < oldMeter {
		fmt.Println("ข้อมูลมิเตอร์ใหม่ไม่ถูกต้อง")
		return
	}
	units, totalPrice = calculate(oldMeter, newMeter, pricePerUnit)
	fmt.Printf("จำนวนหน่วยที่ใช้: %.2f\n", units)
	fmt.Printf("รวมเป็นเงิน: %.2f\n", totalPrice)
}

func calculate(old, new, price float64) (float64, float64) {
	units = newMeter - oldMeter
	totalPrice = units * pricePerUnit
	return units, totalPrice
	// คืน units และ totalPrice
}
