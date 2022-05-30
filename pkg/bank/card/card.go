package card

import "bank/pkg/bank/types"

//PaymentSources представляет информацию об источник оплаты
//Отрицательную сумму и сумму в заблокированых картах игнорируються.
func PaymentSources(cards []types.Card) []types.PaymentSource {
	 paymentSource := make([]types.PaymentSource,len(cards))
	index := 0
	for _, card:=range cards {
		if !card.Active {
			continue
		} else if card.Balance < 0 {
			continue
		}
		paymentSource[index].Type = "card"
		paymentSource[index].Balance = card.Balance
		paymentSource[index].Number = string(card.PAN)
		index++
	}
	return paymentSource
}
