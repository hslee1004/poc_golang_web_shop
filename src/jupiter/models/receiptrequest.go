package models

type ReceiptRequest struct {
	PurchaseTicket string `json:"purchase_ticket"`
	ReceiptId      string `json:"transaction_id"`
	Token          string `json:"token"`
	ProductID      string `json:"product_id"`
}
