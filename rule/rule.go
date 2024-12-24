package Rule

import (
	"fmt"
	"math/rand"
)

func Random() int {
	return rand.Intn(36)
}

func Chan_le(money float32) float32 {
	var x int // so lua chon
	var y float32
	for {
		fmt.Println("Số tiền cược: ")
		fmt.Scan(&y)
		if y > 0 && y <= money {
			break
		}
	}

	fmt.Println("Chọn số chẵn(0) hoặc lẻ(1): ")
	fmt.Scan(&x)
	var point = Random()
	fmt.Print("Số trúng thưởng: ", point)
	if point%2 == 0 && x == 0 {
		fmt.Print("You wwiiinnn")
		money += y
		return money
	} else if point%2 == 1 || x == 1 {
		fmt.Print("You wwiiinnn")
		money += y
		return money
	} else {
		fmt.Print("You lose")
		money -= y
		return money
	}

}

func Dat_cuoc(money float32) float32 {
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
