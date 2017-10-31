package callisto

import (
	//	"encoding/json"
	"fmt"
	//	"io/ioutil"
	"jupiter/libs"
	"jupiter/libs/httplib"
	"jupiter/models"
	"net/http"

	"github.com/astaxie/beego"
)

const (
	RCTAPI_ISSUE   = "receipt_issue"
	RCTAPI_VERIFY  = "receipt_verify"
	RCTAPI_CONSUME = "receipt_consume"
)

var (
	api *ReceiptService
)

type ReceiptService struct {
	http *httplib.HttpService
}

func init() {
	fmt.Printf("callisto init...\n")
	fmt.Printf("receipt base url: %s\n", beego.AppConfig.String("receipt_base_url"))
}

func Service() *ReceiptService {
	if api == nil {
		fmt.Println("creating new http connection....")
		api = newService()
	}
	return api
}

func newService() *ReceiptService {
	return &ReceiptService{
		http: httplib.Service(beego.AppConfig.String("receipt_base_url")),
	}
}

func getAPIURL(api string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("receipt_base_url"), beego.AppConfig.String(api))
}

func (s *ReceiptService) PostIssue(req *models.RVIssueReceiptRequest) (*models.RVReceiptResponse, bool) {
	fmt.Println("token PostIssue: ", req)
	resp := new(models.RVReceiptResponse)
	h := http.Header{"x-api-key": {"OFrY6g2huM38LFUXJLOwm67JhJ5mabCs4OpzXQTQ"}}
	s.http.PostEx(getAPIURL(RCTAPI_ISSUE), h, req, resp)
	return resp, true
}

func (s *ReceiptService) VerifyIssue(req *models.RVIssueReceiptRequest) (*models.RVReceiptResponse, bool) {
	fmt.Println("token VerifyIssue: ", req)
	resp := new(models.RVReceiptResponse)
	h := http.Header{"x-api-key": {"OFrY6g2huM38LFUXJLOwm67JhJ5mabCs4OpzXQTQ"}}
	s.http.PostEx(getAPIURL(RCTAPI_ISSUE), h, req, resp)
	return resp, true
}

func TestPostIssue() {
	req := &models.RVIssueReceiptRequest{
		//TransactionId:        "111",
		TransactionId:        "116", // 111, 112
		ProductId:            "2000",
		UserId:               "user1",
		PurchaseDateInMillis: 1234,
		Price:                1,
		Currency:             "nx",
		PlatformAppId:        "2000",
	}

	fmt.Printf("issue receipt request: %s\n", libs.JsonMarshal(req))
	if resp, ok := Service().PostIssue(req); ok {
		fmt.Printf("issue receipt response: %s\n", libs.JsonMarshal(resp))
	} else {
		fmt.Printf("error on post issue...\n")
	}

}
