package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSource() {
	cards := []types.Card{
		{
         Balance: 1_000_00,
		 PAN: "5058 xxxx xxxx 8888",
		 Active: true,
		},
		{
			Balance: 2_000_00,
			PAN: "5012 xxxx xxxx 8888",
			Active: false,
		},
		{
			Balance: -2_000_00,
			PAN: "5012 xxxx xxxx 8888",
			Active: false,
		},
		{
			Balance: 2_000_00,
			PAN: "5012 xxxx xxxx 8888",
			Active: true,
		},
		{
			Balance: 3_000_00,
			PAN: "5050 xxxx xxxx 8888",
			Active: true,
		},
		
	}
	result:=PaymentSources(cards)
	for _,card:=range result {
		fmt.Println(card.Number)
	} 
	//Output:
	//5058 xxxx xxxx 8888
	//5012 xxxx xxxx 8888
	//5050 xxxx xxxx 8888
}