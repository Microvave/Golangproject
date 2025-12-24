package main

import "fmt"

type Store struct {
	Name  string
	Price float64
}

func main() {
	var stores []Store
	var number int
	var total float64
	var cash float64

	total = 0.0

	stores = append(stores, Store{"น้ำเปล่า", 10.0})
	stores = append(stores, Store{"โค้ก", 20.0})
	stores = append(stores, Store{"ขนม", 30.0})

	for number, s := range stores {
		fmt.Println(number, s.Name, s.Price)
	}

	for {
		fmt.Println("กรุณาเลือกสินค้า")
		fmt.Scan(&number)

		if number == -1 {
			fmt.Println("จบการซื้อ")
			break
		}
		if number < 0 || number >= len(stores) {
			fmt.Println("สินค้าไม่ถูกต้อง")
			continue
		}
		fmt.Println("คุณเลือกซื้อ", stores[number].Name, "ราคา", stores[number].Price)

		total += stores[number].Price
		fmt.Println("รวมราคาสินค้า", total)
	}
	fmt.Println("กรุณาใส่เงิน")
	fmt.Scan(&cash)

	if cash < total {
		fmt.Println("เงินไม่พอ")
		fmt.Println("ขอบคุณที่ใช้บริการ")
	} else {
		fmt.Println("เงินทอน", cash-total)
		fmt.Println("ขอบคุณที่ใช้บริการ")

	}
}
