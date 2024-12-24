package withdraw

import "fmt"

// connect wallet
func Withdraw(wallet map[string]float32, money_betting float32) (bool, float32, string) {

	var option int
	var money float32
	fmt.Print("Moi lua chon vi:\n")
	fmt.Print("1 -->Visa\n2--> Coin\n")
	fmt.Scan(&option)
	if option == 1 {
		for {
			fmt.Print("Nhap so tien ban muon nap:")
			fmt.Scan(&money)
			if money <= money_betting && money > 0 {
				fmt.Printf("Ban rut %v vao vi Visa.\n", money)
				return true, money, "Visa"
			} else {
				fmt.Printf("So tien ban nhap khong hop le. Hay nhap lai.\n")
			}
		}
	} else if option == 2 {
		for {
			fmt.Print("Nhap so tien ban muon nap:")
			fmt.Scan(&money)
			if money <= money_betting && money > 0 {
				fmt.Printf("Ban nap %v vao vi ETH.", money)
				return true, money, "ETH"
			} else {
				fmt.Println("So tien ban nhap khong hop le. Hay nhap lai.")
			}
		}
	} else {
		return false, money, "failer"
	}

}
