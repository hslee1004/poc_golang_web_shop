package jsapi

import (
	"fmt"
	"jupiter/libs"
	"jupiter/libs/httplib"

	"jupiter/models"

	"github.com/astaxie/beego"
	"gopkg.in/jmcvetta/napping.v3"
)

const baseURL = "http://localhost"
const product_id = "20000"
const product_key = "23a3fd9d2ecb9d0cb1e592f593e338c71f95ce54"

type APIName int

const (
	API_INVOICE     = "api_invoice"
	API_PURCHASE    = "api_purchase"
	API_TRANSACTION = "api_transaction"
)

var (
	api *JSAPIService
)

type JSAPIService struct {
	http *httplib.HttpService
}

func init() {
	api = newService()
}

func Service() *JSAPIService {
	if api == nil {
		fmt.Println("creating new http connection....")
		api = newService()
	}
	return api
}

func newService() *JSAPIService {
	return &JSAPIService{
		http: httplib.Service(baseURL),
	}
}

func getAPIURL(api string) string {
	return fmt.Sprintf("%s%s", baseURL, beego.AppConfig.String(api))
}

//
// issue
//
func (s *JSAPIService) RegisterInvoice(invoice *models.Invoice) (*models.APIResponseEx, bool) {
	rs := new(models.APIResponseEx)
	url := getAPIURL(API_INVOICE)
	fmt.Println("RegisterInvoice api url: ", url)

	resp, err := s.http.Sling.New().Post(url).BodyJSON(invoice).Receive(rs, rs)
	fmt.Printf("RegisterInvoice status:%s\n", resp.Status)
	if err == nil {
		fmt.Printf("jsapi.RegisterInvoice: response: %s\n", libs.JsonMarshal(rs))
		if rs.Success.Code == 0 {
			return rs, true
		}
		fmt.Println("jsapi.RegisterInvoice: checkout - here...", rs)
	}
	fmt.Println("response is not nil: error:", err)
	return rs, false
}

//
// using napping
//     https://github.com/jmcvetta/napping/tree/master
//
func (s *JSAPIService) RegisterInvoice2(invoice *models.Invoice) (*models.APIResponse, bool) {
	rs := &models.APIResponse{}
	url := getAPIURL(API_INVOICE)
	fmt.Println("RegisterInvoice api url: ", url)
	ns := napping.Session{} // GPL : need to replace
	resp, _ := ns.Post(url, invoice, rs, rs)
	fmt.Printf("[jsapi][RegisterInvoice2] status: %s\n", resp.Status())
	fmt.Printf("[jsapi][RegisterInvoice2] response: %s\n", libs.JsonMarshal(rs))
	return rs, true
}

/* response
{
  "success": {
    "code": 0,
    "data": {
      "purchase_ticket": "606eb564-b2df-46ff-8724-de79742e5d24",
      "receipt": "363878DAAD60479A983F80FE13DA6702"
    }
  }
}
*/
func (s *JSAPIService) CompleteTransaction(invoice *models.Invoice) (*models.PurchaseResponse, bool) {
	url := getAPIURL(API_PURCHASE)
	fmt.Println("[jsapi] CompleteTransaction url: ", url)

	req := &models.PurchaseRequest{PurchaseTicket: invoice.Ticket, ProductID: invoice.ProductId}
	rs := &models.PurchaseResponse{}
	//af := new(models.NXAPIError)
	ns := napping.Session{}
	// billing: invoice --> jupiter
	resp, _ := ns.Post(url, req, rs, rs)

	fmt.Printf("[jsapi][CompleteTransaction] status: %s\n", resp.Status())
	fmt.Printf("[jsapi][CompleteTransaction] response: %s\n", libs.JsonMarshal(rs))

	if rs.Success != nil {
		return rs, true
	}
	return rs, false
}

func (s *JSAPIService) CompleteWalletTrx(trxId string) (*models.PurchaseResponse, bool) {
	req := &models.TransactionRequest{TransactionId: trxId, ProductID: "20000"}
	rs := &models.PurchaseResponse{}
	ns := napping.Session{}
	// billing: invoice --> jupiter
	url := getAPIURL(API_TRANSACTION)
	resp, _ := ns.Post(url, req, rs, rs)
	fmt.Printf("[jsapi][GSCompleteTransaction] status: %s\n", resp.Status())
	fmt.Printf("[jsapi][GSCompleteTransaction] response: %s\n", libs.JsonMarshal(rs))
	if rs.Success != nil {
		return rs, true
	}
	return rs, false
}

func (s *JSAPIService) VerifyReceipt(receipt string) (*models.PurchaseResponse, bool) {
	req := &models.TransactionRequest{Receipt: receipt, ProductID: "20000"}
	rs := &models.PurchaseResponse{}
	ns := napping.Session{}
	// billing: invoice --> jupiter
	url := getAPIURL(API_TRANSACTION)
	resp, _ := ns.Post(url, req, rs, rs)
	fmt.Printf("[jsapi][GSCompleteTransaction] status: %s\n", resp.Status())
	fmt.Printf("[jsapi][GSCompleteTransaction] response: %s\n", libs.JsonMarshal(rs))
	if rs.Success != nil {
		return rs, true
	}
	return rs, false
}
