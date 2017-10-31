package models

//{
//  "name": "nx",
//  "payment_method": "nx_prepaid",
//  "amount": 1000,
//  "currency": "NX"
//},

type WalletBalance struct {
	Name          string  `json:"name,omitempty"`
	PaymentMethod string  `json:"payment_method,omitempty"`
	Amount        float32 `json:"amount"`
	Currency      string  `json:"currency,omitempty"`
}
