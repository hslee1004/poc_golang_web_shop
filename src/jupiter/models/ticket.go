package models

type Ticket struct {
	Ticket string `json:"ticket"`
}

type TicketRequest struct {
	Token     string `json:"access_token,omitempty", url:"access_token,omitempty"`
	ProductId string `json:"prod_id,omitempty", url:"prod_id,omitempty"`
}

type TicketResponse struct {
	Ticket string       `json:"ticket,omitempty"`
	Error  *NXAPIDetail `json:"error,omitempty"`
}
