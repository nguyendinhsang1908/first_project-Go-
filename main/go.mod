module game_roulleter.com/main

go 1.23.4

replace game_roulleter.com/connect_wallet => ../connect_wallet

replace game_roulleter.com/betting => ../betting

replace game_roulleter.com/rule => ../rule

replace game_roulleter.com/withdraw => ../withdraw

require (
	game_roulleter.com/betting v0.0.0-00010101000000-000000000000
	game_roulleter.com/connect_wallet v0.0.0-00010101000000-000000000000
	game_roulleter.com/withdraw v0.0.0-00010101000000-000000000000
)

replace game_roulleter/withdraw => ../withdraw
