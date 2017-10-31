package cashbroker

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"jupiter/libs"
	"jupiter/models"
	"net/http"

	"github.com/astaxie/beego"
)

const (
	CB_RESPCODE_SUCCESS = "50000"
)

func SetXMLHeader(req *http.Request) {
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
}

func Send(data []byte, v interface{}) error {
	// request
	req, err := http.NewRequest("POST", beego.AppConfig.String("cashbroker"), bytes.NewBuffer(data))
	SetXMLHeader(req)
	resp, err := Service().Http.Do(req)
	if err != nil {
		fmt.Println("cashbroker, error...")
	}
	defer resp.Body.Close()
	fmt.Println("response status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("response: %s\n", string(body))
	if err := xml.Unmarshal([]byte(libs.ReplaceXMLHeader(string(body))), &v); err == nil {
		fmt.Printf("Send: %v\n", v)
	} else {
		fmt.Printf("Send: %v, error:%v\n", v, err)
	}
	return err
}

func GetBalance(user_id string) *models.UserBalance {
	st := models.NewBalanceReq(user_id)
	data, _ := st.XMLEncode(nil)
	v := models.XmlCashBrokerSection{}
	bal := &models.UserBalance{}
	if err := Send(data, &v); err == nil {
		getBalance(&v.CheckBalanceResponse, bal)
		//return libs.JsonMarshal(&v.CheckBalanceResponse)
	}
	return bal
}

//type WalletBalance struct {
//	Name          string  `json:"name,omitempty"`
//	PaymentMethod string  `json:"payment_method,omitempty"`
//	Amount        float32 `json:"amount,omitempty"`
//	Currency      string  `json:"currency,omitempty"`
//}
//
// new design
//
func GetWallets(user_id string) []models.WalletBalance {
	st := models.NewBalanceReq(user_id)
	data, _ := st.XMLEncode(nil)
	v := models.XmlCashBrokerSection{}
	bal := &models.UserBalance{}
	bals := []models.WalletBalance{} // new
	if err := Send(data, &v); err == nil {
		getBalance(&v.CheckBalanceResponse, bal)
		//return libs.JsonMarshal(&v.CheckBalanceResponse)
		bals = append(bals,
			models.WalletBalance{
				Name:          "nx_prepaid",
				PaymentMethod: "nx_prepaid",
				Amount:        bal.Prepaid,
				Currency:      "NX",
			})
		bals = append(bals,
			models.WalletBalance{
				Name:          "nx_credit",
				PaymentMethod: "nx_credit",
				Amount:        bal.Credit,
				Currency:      "NX",
			})
	}
	return bals[:]
}

func getBalance(resp *models.CheckBalanceResponse, bal *models.UserBalance) {
	bal.Total = resp.Balance
	if beego.AppConfig.String("ruleid_nx_prepaid") == "WSR2" {
		// nx prepaid balance
		bal.Prepaid = resp.BalanceByRule
		bal.Credit = resp.Balance - resp.BalanceByRule
	} else {
		bal.Credit = resp.BalanceByRule
		bal.Prepaid = resp.Balance - resp.BalanceByRule
	}
}

/*
<xmlCashBrokerSection type="Nexon.CashBroker.Universal.Entity.RequestPaymentResponse, Nexon.CashBroker.Universal, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null">
  <RequestPaymentResponse xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <Result>50000</Result>
  <DetailMessage/>
  <TransactionId>F4EA9616CD464225BF0B96962428B0A1</TransactionId>
  </RequestPaymentResponse>
</xmlCashBrokerSection>
*/
func RequestPayment(serviceCode string, i *models.Invoice) (*models.RequestPaymentSection, bool) {
	req := models.NewPayReq(serviceCode, i)
	data, _ := req.XMLEncode(nil)
	fmt.Printf("[cashbroker.RequestPayment]: request: \n%s\n", []byte(data))
	v := models.RequestPaymentSection{Response: &models.RequestPaymentResponse{}} //change to point
	Send(data, &v)
	fmt.Printf("[cashbroker.RequestPayment]:\n%v\n", libs.JsonMarshal(v))
	if v.Response.Result == CB_RESPCODE_SUCCESS {
		i.TransactionId = v.Response.TransactionId
		fmt.Printf("invoice, transactionId: \n%v\n", libs.JsonMarshal(i))
		return &v, true
	}
	return &v, false
}

//
// test
//
func TestRequestPayment() {
	// test
	items := [...]models.Item{
		{Id: "10002", Name: "test_jupiter_item_2", Price: 1, Qty: 1},
		{Id: "10003", Name: "test_jupiter_item_3", Price: 2, Qty: 1},
	}

	invoice := models.Invoice{Ticket: "trx_test_111", Items: items[:]}
	req := models.NewPayReq("test", &invoice)
	data, _ := req.XMLEncode(nil)
	fmt.Printf("TestRequestPayment: %s\n", []byte(data))
}

/*
<xmlCashBrokerSection type="Nexon.CashBroker.Universal.Entity.CommitPaymentResponse, Nexon.CashBroker.Universal, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null">
  <CommitPaymentResponse xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <Result>50000</Result>
  <DetailMessage/>
  <PaymentNo>8328963</PaymentNo>
  </CommitPaymentResponse>
</xmlCashBrokerSection>
*/
func CommitPayment(i *models.Invoice) (*models.CommitPaymentSection, bool) {
	req := models.NewCommitPayReq(i.TransactionId, i)
	data, _ := req.XMLEncode(nil)
	fmt.Printf("CommitPayment: \n%s\n", []byte(data))
	v := models.CommitPaymentSection{Response: &models.CommitPaymentResponse{}}
	Send(data, &v)
	fmt.Printf("CommitPayment: \n%v\n", libs.JsonMarshal(v))
	if v.Response.Result == CB_RESPCODE_SUCCESS {
		return &v, true
	}
	return &v, false
}

// CompletePaymentSection
func CompletePayment(i *models.Invoice) (*models.CompletePaymentSection, bool) {
	req := models.NewCommitPayReq(i.TransactionId, i)
	data, _ := req.XMLEncode(nil)
	fmt.Printf("CompletePayment: \n%s\n", []byte(data))
	v := models.CompletePaymentSection{Response: &models.CompletePaymentResponse{}} // todo
	Send(data, &v)
	fmt.Printf("CompletePayment: \n%v\n", libs.JsonMarshal(v))
	if v.Response.Result == CB_RESPCODE_SUCCESS {
		return &v, true
	}
	return &v, false
}
