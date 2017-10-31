package models

type TransactionRequest struct {
	PurchaseTicket string `json:"purchase_ticket"`
	TransactionId  string `json:"transaction_id"`
	Receipt        string `json:"receipt"`
	Token          string `json:"token"`
	ProductID      string `json:"product_id"`
	CallbackAPI    string `json:"callback_api"`
}
