package main

import (
	"fmt"

	"game_roulleter.com/betting"
	"game_roulleter.com/connect_wallet"
	"game_roulleter.com/withdraw"
)

type Person struct {
	name   string
	age    int
	wallet map[string]float32
}

func main() {
	var pers1 Person
	var pers2 Person

	//input_data_person
	pers1.name = "name1"
	pers1.age = 20
	pers1.wallet = map[string]float32{"Visa": 200, "ETH": 500}

	pers2.name = "name2"
	pers2.age = 22
	pers2.wallet = map[string]float32{"Visa": 300, "ETH": 700}

	printPerson(pers1)
	printPerson(pers2)

	fmt.Println("--------")

	play(pers1)

}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	for k, v := range pers.wallet {
		fmt.Printf("%v : %v \n", k, v)
	}

}
func play(pers1 Person) {

	var money_betting float32
	//money in wallet game
	var money float32
	for {
		fmt.Print("Mời chọn chức năng:\n")
		fmt.Print("1 --> Connect_Wallet\n2--> Betting\n3 --> Withdraw\n4 --> exit\n")
		var option int
		fmt.Scan(&option)
		//money in betting

		switch option {
		case 1:
			fmt.Println("Connect_Wallet")
			fmt.Print(money)
			var check, money_card, option_wallet = connect_wallet.Connect_wallet(pers1.wallet)
			if check {

				if option_wallet == "Visa" {

					pers1.wallet["Visa"] = pers1.wallet["Visa"] - money_card
					money = money + money_card
					fmt.Printf("Số tiền trong ví game của bạn:%v\n", money)
				} else {

					pers1.wallet["ETH"] = pers1.wallet["ETH"] - money_card
					money = money + money_card
					fmt.Printf("Số tiền trong ví game của bạn:%v\n", money)
				}

			} else {
				fmt.Println("failer connect_wallet")
			}
		case 2:
			fmt.Println("Betting", money)
			money_betting = betting.Betting(money)

		case 3:
			fmt.Println("Withdraw")
			var check, money_card, option_wallet = withdraw.Withdraw(pers1.wallet, money_betting)
			if check {
				if option_wallet == "Visa" {

					//fmt.Println("so tien con lai trong vi la:", pers1.wallet["Visa"]-money)
					pers1.wallet["Visa"] = pers1.wallet["Visa"] + money_card
					fmt.Printf("Rút thành công.Số tiền trong ví Visa của bạn là:%v\n", pers1.wallet["Visa"])
				} else {

					//fmt.Println("so tien con lai trong vi la:", pers1.wallet["ETH"]-money)
					pers1.wallet["ETH"] = pers1.wallet["ETH"] + money_card
					fmt.Printf("Rút thành công.Số tiền trong ví ETH của bạn là:%v\n", pers1.wallet["ETH"])
				}

			} else {
				fmt.Println("failer connect_wallet")
			}
		case 4:
			fmt.Println("exit")
			printPerson(pers1)
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
