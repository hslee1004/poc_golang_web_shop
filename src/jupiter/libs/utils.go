package libs

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/google/uuid"
)

func ToInvoiceKey(pid string, uno string) string {
	return fmt.Sprintf("inv:%s-%s", pid, uno)
}

func JsonMarshal(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("{\"error\":\"%s\"}", "marshal error.")
	}
	return fmt.Sprintf("%s", b)
}

func JsonUnmarshal(data []byte, v interface{}) bool {
	if err := json.Unmarshal(data, &v); err == nil {
		fmt.Printf("Unmarshal: %v\n", v)
		return true
	} else {
		fmt.Println("un-marshal error:", err)
		return false
	}
}

func ReplaceXMLHeader(v string) string {
	return strings.Replace(v, `<?xml version="1.0" encoding="utf-16"?>`, "", -1)
}

// format : eaeccaf9-04fe-45e6-a1ee-2de6edf0a1cd
func GetUUID() string {
	return fmt.Sprintf("%s", uuid.New())
}

func GetBillingUserCart(ticket string, invoice string) string {
	return fmt.Sprintf("%s%s?ticket=%s&invoice_id=%s",
		beego.AppConfig.String("shop_server"),
		beego.AppConfig.String("shop_cart"),
		ticket,
		invoice,
	)
}

func GetBillingUserOrder(ticket string, invoice string) string {
	return fmt.Sprintf("%s%s?ticket=%s&invoice_id=%s",
		beego.AppConfig.String("shop_server"),
		beego.AppConfig.String("shop_order"),
		ticket,
		invoice,
	)
}

func GetServiceCode(prodId string) string {
	return beego.AppConfig.String(fmt.Sprintf("service_code_%s", prodId))
}

func GetNow() string {
	now := time.Now()
	return fmt.Sprintf("%s", now.Format("2006-01-02 15:04:05"))
}

func GetMyLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	rs := ""
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		return rs
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
				rs = ipnet.IP.String()
				break
			}
		}
	}
	return rs
}
