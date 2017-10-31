package nxapi

import (
	"encoding/json"
	"fmt"
	//	"io/ioutil"
	"jupiter/libs"
	"jupiter/libs/httplib"
	"jupiter/models"

	"github.com/astaxie/beego"
)

const baseURL = "http://api.nexon.net/"
const PRODUCT_ID = "20000"
const product_key = "23a3fd9d2ecb9d0cb1e592f593e338c71f95ce54"
const PRODUCT_ID_BILLING = "30020"

type APIName int

const (
	API_TOKEN = iota
	API_TICKET
	API_USERID
	API_PASSPORT
)

var (
	api *NXAPIService
)

type NXAPIService struct {
	http *httplib.HttpService
}

func init() {
	api = NewNXAPIService()
}

func Service() *NXAPIService {
	if api == nil {
		fmt.Println("creating new http connection....")
		api = NewNXAPIService()
	}
	return api
}

func NewNXAPIService() *NXAPIService {
	return &NXAPIService{
		http: httplib.Service(baseURL),
	}
}

func getAPIURL(n APIName) string {
	host := beego.AppConfig.String("nxapi_server")
	switch n {
	case API_TOKEN:
		return fmt.Sprintf("%s%s", host, "/auth/token")
	case API_TICKET:
		return fmt.Sprintf("%s%s", host, "/auth/ticket")
	case API_USERID:
		return fmt.Sprintf("%s%s", host, "/users/me/userid")
	case API_PASSPORT:
		return fmt.Sprintf("%s%s", host, "/users/me/passport")
	}
	return ""
}

/*
{
  "token": "NX1_2588312M0pMYWJqUURnPQ2",
  "refresh_token": "NX1_258831NWxzPQ2",
  "expires_in": 21591,
  "auth_token": null
}
{
  "error": {
    "code": "6",
    "type": "Bad Request",
    "message": "Parameter is missing.(ticket)"
  }
}
*/
func (s *NXAPIService) GetToken(ticket string) (*models.TokenResponse, bool) {
	p := &models.TokenRequest{Ticket: ticket, ProductID: PRODUCT_ID, SecretKey: product_key}
	fmt.Println("token request: ", *p)
	rs := new(models.TokenResponse)
	url := getAPIURL(API_TOKEN)
	fmt.Println("api url: ", url)

	// test
	s.http.PostEx(url, nil, p, rs)
	/*
		_, err := s.http.Sling.New().Post(url).BodyJSON(p).Receive(rs, af)
		if err == nil {
			fmt.Printf("GetToken response: %v\n", rs)
			if rs.Error.Code == "" {
				return rs, true
			} else {
				return rs, false
			}
		}
		fmt.Println("GetToken: response is not nil:", err)
		return rs, false
	*/

	return rs, true

}

/*
{
  "success": {
    "code": 0,
    "data": {
      "token": "NX1_25883128_RVJvb3VI....{removed}...WtNbVNjPQ2",
      "user_no": 25883128,
      "prod_id": "10800",
      "access_ip": "10.8.25.144",
      "auth_type": "OAUTH",
      "expires_in": 7177
    }
  }
}
*/
func (s *NXAPIService) GetTokenDetail(token string) (*models.TokenResponse, bool) {
	//p := &models.ParamsToken{Token: token}
	p := &models.TokenRequest{Token: token}
	fmt.Println("GetTokenDetail request: ", *p)
	v := &models.TokenResponse{}

	url := getAPIURL(API_TOKEN)
	fmt.Println("api url: ", url)

	if str, ok := s.http.Get(url, p, v); ok {
		fmt.Printf("response: %v\n", str)
		fmt.Printf("json: %v\n", v)
		return v, true
	}
	fmt.Println("response is not nil.")
	return nil, false
}

//
// https://nexonusa.atlassian.net/wiki/display/SUPER/Mantis+Users+APIs#MantisUsersAPIs-GetProfile#MantisUsersAPIs-GetProfile
//
func (s *NXAPIService) GetUserId(token string) (*models.UserIdResponse, bool) {
	p := &models.TokenRequest{Token: token}
	fmt.Println("GetUserId request: ", *p)
	v := &models.UserIdResponse{}

	url := getAPIURL(API_USERID)
	fmt.Println("api url: ", url)

	if str, ok := s.http.Get(url, p, v); ok {
		fmt.Printf("GetUserId response: %v\n", str)
		fmt.Printf("json: %v\n", v)
		return v, true
	}
	fmt.Println("response is not nil.")
	return nil, false
}

// from api.nexon.net
//		{"ticket":"3765f4ea-5c4c-475a-9009-c23928b42b8c"}
//		{"error":{"code":"6","type":"Bad Request","message":"Parameter is missing.(product id)"}}
//
//func (s *NXAPIService) CreateTicket(token string) (*models.TicketResponse, bool) {
func (s *NXAPIService) CreateTicket(token string, pid string) (*models.TicketResponse, bool) {
	p := &models.TicketRequest{Token: token, ProductId: pid}
	fmt.Println("CreateTicket request: ", *p)
	if req, err := s.http.Sling.New().Post(getAPIURL(API_TICKET)).BodyJSON(p).Request(); err == nil {
		resp, _ := httplib.GetHttp().Do(req)
		v := &models.TicketResponse{}
		json.NewDecoder(resp.Body).Decode(v)
		fmt.Printf("CreateTicket response: %s\n", libs.JsonMarshal(v)) // dev-only
		return v, true
	} else {
		return nil, false
	}
}

//
// token to passport string
//     /users/me/passport  GET
//
// response:
// {
//  "user_no": "1000000974",
//  "passport": "NP12:auth01:126:1000000974:roGcpbZE5pgD04yrJtXLOACt~T53cXNc",
//  "auth_token": "TmRYg9tDIzBvYy95g=="
// }
//
func (s *NXAPIService) GetPassport(token string, userIP string) (*models.PassportResponse, bool) {
	p := &models.PassportRequest{Token: token, UserIP: userIP}
	fmt.Println("GetPassport request: ", *p)
	v := &models.PassportResponse{}

	url := getAPIURL(API_PASSPORT)
	fmt.Println("GetPassport url: ", url)

	if str, ok := s.http.Get(url, p, v); ok {
		fmt.Printf("response: %v\n", str)
		fmt.Printf("json: %v\n", v)
		return v, true
	}
	return nil, false
}

/*
func (s *NXAPIService) TestGetToken(ticket string) (*models.NXAPISuccess, *models.NXAPIError) {
	p := &models.TokenRequest{Ticket: ticket, ProductID: product_id, SecretKey: product_key}
	fmt.Println("token request: ", *p)
	url := getAPIURL(API_TOKEN)
	fmt.Println("api url: ", url)

	//	_, err := s.sling.New().Post(url).BodyJSON(p).Receive(as, af)
	if req, err := s.http.Sling.New().Post(url).BodyJSON(p).Request(); err == nil {
		resp, _ := httplib.GetHttp().Do(req)
		d, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response: %s\n", string(d))
		fmt.Println("response is not nil.")
		return nil, af
	} else {
		return nil, af
	}
}
*/
