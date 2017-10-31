//
// beego tutorial : https://github.com/beego/tutorial
// test:
// http://localhost:8080/api/store/balance?user_id=mantistest3
//
package main

import (
	"fmt"
	ctrls "jupiter/controllers"
	_ "jupiter/routers"

	"flag"
	"jupiter/libs/callisto"
	//"jupiter/libs/cashbroker"
	//"jupiter/libs/mantisdb"

	bg "github.com/astaxie/beego"
)

func main() {
	// test
	//cashbroker.TestRequestPayment()

	// test
	//mantisdb.TestConn()

	// test
	//callisto.TestPostIssue()

	env := flag.String("env", "dev", "env")
	port := flag.String("port", "80", "port")
	flag.Parse()

	bg.Info(bg.BConfig.RunMode, *env)

	// init
	callisto.Service()

	bg.Router("/wallet/balance", &ctrls.BalanceController{})
	bg.Router("/wallet/invoice", &ctrls.InvoiceController{})
	bg.Router("/wallet/purchase", &ctrls.WalletPurchaseController{})
	bg.Router("/wallet/transaction", &ctrls.WalletTrxController{})
	bg.Router("/wallet/receipt", &ctrls.WalletReceiptController{})

	// front
	bg.Router("/front/store/cart", &ctrls.StoreCartController{})
	bg.Router("/front/store/order", &ctrls.StoreOrderController{})
	bg.Router("/front/store/charge", &ctrls.PassportController{})
	bg.Router("/front/store/purchase-nx", &ctrls.PurchaseNXController{})

	// test
	bg.Router("/api/test/async", &ctrls.TestAsyncController{})
	bg.Router("/api/test/api", &ctrls.TestAPIClientController{})
	bg.Router("/api/test/ping", &ctrls.PingController{})

	// simulate game-server for test
	//	bg.Router("/shop/mall", &ctrls.TestMallController{})
	//	bg.Router("/shop/cart", &ctrls.TestServerController{})

	bg.Router("/demo/cart", &ctrls.TestServerController{})
	bg.Router("/shop/cart", &ctrls.TestMallController{}) // can't change name due to demo

	bg.Router("/game/server/redirect", &ctrls.TestServerRedirectController{})

	bg.Run(fmt.Sprintf(":%s", *port))
}
