package models

type PurchaseRequest struct {
	PurchaseTicket string `json:"purchase_ticket"`
	Token          string `json:"token"`
	ProductID      string `json:"product_id"`
}

type PurchaseResponse struct {
	Success *PurchaseResponseSuccess `json:"success,omitempty"`
	Error   *PurchaseResponseError   `json:"error,omitempty"`
}

type PurchaseResponseSuccess struct {
	Code    int    `json:"code"` // due to api.nexon.net
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
	//Data    PurchaseRequestData `json:"data,omitempty"`
	Data PurchaseData `json:"data,omitempty"`
}

type PurchaseResponseError struct {
	Code    string `json:"code"` // due to api.nexon.net
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
	//Data    PurchaseRequestData `json:"data,omitempty"`
	Data PurchaseData `json:"data,omitempty"`
}

//
//type PurchaseRequestData struct {
type PurchaseData struct {
	PurchaseTicket string `json:"purchase_ticket,omitempty"`
	TransactionId  string `json:"transaction_id,omitempty"`
	ReceiptId      string `json:"receipt_id,omitempty"`
	PaymentId      string `json:"payment_id,omitempty"` // Nexon payment id
	CallbackAPI    string `json:"callback_api,omitempty"`
}
