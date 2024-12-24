package betting

import (
	"fmt"
	"math/rand"
)

func Betting(money float32) float32 {
	var x int // so lua chon
	var y float32
	for {
		fmt.Println("Chọn số từ 0 đến 36: ")
		fmt.Scan(&x)
		if x >= 0 && x <= 36 {
			break
		}
	}
	for {
		fmt.Println("Số tiền cược: ")
		fmt.Scan(&y)
		if y > 0 && y <= money {
			break
		}
	}
	var a int = rand.Intn(36) // so trung thuong
	fmt.Printf("Số trúng thưởng : %v\n", a)
	if a == x {
		fmt.Println("You win")
		money += y
	} else {
		fmt.Println("You lose")
		money -= y
	}
	return money
}
