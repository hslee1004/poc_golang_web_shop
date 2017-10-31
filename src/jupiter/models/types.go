package models

type Params struct {
	Count int `url:"count,omitempty"`
}

type ParamsToken struct {
	Token string `url:"access_token,omitempty`
}

type TokenRequest struct {
	Ticket    string `json:"ticket,omitempty" url:"ticket,omitempty"`
	Token     string `json:"token,omitempty" url:"access_token,omitempty"`
	ProductID string `json:"product_id,omitempty" url:"product_id,omitempty"`
	SecretKey string `json:"secret_key,omitempty" url:"secret_key,omitempty"`
}

type TokenResponse struct {
	Token        string           `json:"token,omitempty"`
	RefreshToken string           `json:"refresh_token,omitempty"`
	ExpiresIn    int              `json:"expires_in,omitempty"`
	AuthToken    string           `json:"auth_token,omitempty"`
	Success      NXAPIDetail      `json:"success,omitempty"`
	Error        NXAPIErrorDetail `json:"error,omitempty"`
}

type UserIdResponse struct {
	UserNo string `json:"user_no,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

type APIResponse struct {
	Success *NXAPIDetail      `json:"success,omitempty"`
	Error   *NXAPIErrorDetail `json:"error,omitempty"`
}

type PassportRequest struct {
	Token  string `json:"access_token,omitempty" url:"access_token,omitempty"`
	UserIP string `json:"login_ip,omitempty" url:"login_ip,omitempty"`
}

type PassportResponse struct {
	UserNo    string `json:"user_no,omitempty"`
	Passport  string `json:"passport,omitempty"`
	AuthToken string `json:"auth_token,omitempty"`
}

type APIResponseEx struct {
	Success NXAPIDetail `json:"success,omitempty"`
	//Error   NXAPIDetail `json:"error,omitempty"`
}

//type NXAPIError struct {
//	Error NXAPIDetail
//}

type NXAPISuccess struct {
	Success NXAPIDetail
}

type NXAPIDetail struct {
	//Code    int       `json:"code,omitempty"`
	Code    int       `json:"code"`
	Type    string    `json:"type,omitempty"`
	Message string    `json:"message,omitempty"`
	Data    NXAPIData `json:"data,omitempty"`
}

type NXAPIErrorDetail struct {
	Code    string    `json:"code"`
	Type    string    `json:"type,omitempty"`
	Message string    `json:"message,omitempty"`
	Data    NXAPIData `json:"data,omitempty"`
}

type NXAPIData struct {
	Token            string `json:"token,omitempty"`
	Ticket           string `json:"ticket,omitempty"`
	InvoiceId        string `json:"invoice_id,omitempty"`
	UserNo           int    `json:"user_no,omitempty"`
	ProdId           string `json:"prod_id,omitempty"`
	AccessIp         string `json:"access_ip,omitempty"`
	AuthType         string `json:"auth_type,omitempty"`
	ExpiresIn        int    `json:"expires_in,omitempty"`
	BillingURL       string `json:"billing_url,omitempty"`
	ForDevInvoiceURL string `json:"for_dev_invoice_url,omitempty"` // for dev purpose
}

//{
//  "success": {
//    "code": 0,
//    "data": {
//      "token": "NX1_25883128_RVJvb3VI....{removed}...WtNbVNjPQ2",
//      "user_no": 25883128,
//      "prod_id": "10800",
//      "access_ip": "10.8.25.144",
//      "auth_type": "OAUTH",
//      "expires_in": 7177
//    }
//  }
//}
