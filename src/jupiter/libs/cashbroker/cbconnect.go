package cashbroker

import (
	"fmt"
	"jupiter/libs/httplib"
	"net/http"
)

var (
	api *CashBrokerConnect
)

type CashBrokerConnect struct {
	Http *http.Client
}

func init() {
	api = newService()
}

func Service() *CashBrokerConnect {
	if api == nil {
		fmt.Println("creating new http connection....")
		api = newService()
	}
	return api
}

func newService() *CashBrokerConnect {
	return &CashBrokerConnect{
		Http: httplib.GetHttp(),
	}
}
