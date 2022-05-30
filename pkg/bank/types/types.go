package types

//Money представляет собой денежную единицу в минимальных единицах(копейки, дирамы и.т)
type Money int64

//Currency представляет собой код валюты
type Currency string

//Код валюты
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN представляет номер карты
type PAN string

//Card представляет собой информацию о платёжном карте
type Card struct {
	 ID int
	 PAN PAN 
	 Balance Money
	 Currency Currency
	 Color string
	 Name string
	 Active bool
}

//PaymentSource представляет собой информацию об источники оплаты
type PaymentSource struct {
	Type string
	Number string
	Balance Money
}