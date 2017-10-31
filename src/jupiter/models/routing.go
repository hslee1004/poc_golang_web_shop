package models

import (
	"fmt"
	URL "net/url"

	"github.com/astaxie/beego"
)

const (
	JAPI_INVOICE     = "api_invoice"
	JAPI_PURCHASE    = "api_purchase"
	JAPI_TRANSACTION = "api_transaction"
	JAPI_RECEIPT     = "api_receipt"
	JAPI_PASSPORT    = "api_passport"
	JFRONT_CART      = "shop_cart"
)

func GetRoutingURL(api string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("api_host"), beego.AppConfig.String(api))
}

func GetQueryURL(url string, vs map[string]string) string {
	u, _ := URL.Parse(url)
	q := u.Query()
	for k, v := range vs {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	fmt.Printf("GetRoutingURLEx: %s", u.String())
	return u.String()
}
