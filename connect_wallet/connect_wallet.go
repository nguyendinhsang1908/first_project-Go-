package connect_wallet

import "fmt"

// connect wallet
func Connect_wallet(wallet map[string]float32) (bool, float32, string) {

	var option int
	fmt.Print("Moi lua chon vi:\n")
	fmt.Print("1 -->Visa\n2--> Coin\n")
	fmt.Scan(&option)
	var money float32 = 0
	if option == 1 {
		for {
			fmt.Print("Nhap so tien ban muon nap:")
			fmt.Scan(&money)
			if money <= wallet["Visa"] && money > 0 {
				fmt.Printf("Ban nap %v tu vi Visa.\n", money)
				return true, money, "Visa"
			} else {
				fmt.Printf("So tien ban nhap khong hop le. Hay nhap lai.\n")
			}
		}
	} else if option == 2 {
		for {
			fmt.Print("Nhap so tien ban muon nap:")
			fmt.Scan(&money)
			if money <= wallet["ETH"] && money > 0 {
				fmt.Printf("Ban nap %v tu vi ETH.", money)
				return true, money, "ETH"
			} else {
				fmt.Println("So tien ban nhap khong hop le. Hay nhap lai.")
			}
		}
	} else {
		return false, money, "failer"
	}

}
