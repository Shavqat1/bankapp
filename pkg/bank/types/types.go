package types

//Money представляет собой денежную сумму в минимальных единицах{центы, копейки, дирамы и т.д}
type Money int64

//Currency представляет код валюты
type Currency string

//Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN представляет номер карты
type PAN string

//Card представляет информатцию о платёжной карте
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Nname    string
	Active   bool
}
type Payment struct {
	ID     int
	Amount Money
}
