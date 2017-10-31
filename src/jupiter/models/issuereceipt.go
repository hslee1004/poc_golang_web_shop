package models

type RVIssueReceiptRequest struct {
	TransactionId        string `json:"transactionId"`
	ProductId            string `json:"productId"`
	UserId               string `json:"userId"`
	PurchaseDateInMillis int    `json:"purchaseDateInMillis"`
	Price                int    `json:"price"`
	Currency             string `json:"currency"`
	PlatformAppId        string `json:"platformAppId"` // ??
}

type RVReceiptResponse struct {
	Payload      *RVReceiptPayload `json:"payload,omitempty"`
	ResponseCode int               `json:"responseCode,omitempty"`
}

type RVReceiptPayload struct {
	ReceiptData *RVReceiptData `json:"receiptData,omitempty"`
}

type RVReceiptData struct {
	Contents *RVReceiptContents `json:"contents,omitempty"`
	Sha256   string             `json:"sha256,omitempty"`
	Version  int                `json:"version,omitempty"`
}

type RVReceiptContents struct {
	Currency                  string `json:"currency,omitempty"`
	PlatformAppId             string `json:"platformAppId,omitempty"`
	Price                     int    `json:"price,omitempty"`
	ProductId                 string `json:"productId,omitempty"`
	ProfileName               string `json:"profileName,omitempty"`
	PurchaseDateInMillis      int    `json:"purchaseDateInMillis,omitempty"`
	ReceiptCreateTimeInMillis int    `json:"receiptCreateTimeInMillis,omitempty"`
	TransactionId             string `json:"transactionId,omitempty"`
	UserId                    string `json:"userId,omitempty"`
	Version                   int    `json:"version,omitempty"`
}
