package controllers

import (
	"fmt"

	"jupiter/libs"
	"jupiter/libs/jsapi"
	"jupiter/models"

	"github.com/astaxie/beego"
)

type TestServerRedirectController struct {
	beego.Controller
}

func (this *TestServerRedirectController) Get() {
	fmt.Printf("in TestServerController....\n")
	this.Ctx.WriteString("in test game client and server: \n")
	receipt := this.GetString("receipt_id")
	trxId := this.GetString("transaction_id")
	callback := this.GetString("callback_api")

	req := &models.TransactionRequest{Receipt: receipt, TransactionId: trxId, CallbackAPI: callback}
	this.Ctx.WriteString(fmt.Sprintf("req: %s \n", libs.JsonMarshal(req)))
	// call jupiter api server
	if trxId != "" {
		if resp, ok := jsapi.Service().CompleteWalletTrx(trxId); ok {
			// okay
			// show success
			this.Ctx.WriteString("CompleteWalletTrx:\n")
			this.Ctx.WriteString(fmt.Sprintf("%s", libs.JsonMarshal(resp)))
		}
	} else if receipt != "" {
		if resp, ok := jsapi.Service().VerifyReceipt(receipt); ok {
			// okay
			// show success
			this.Ctx.WriteString("VerifyReceipt:\n")
			this.Ctx.WriteString(fmt.Sprintf("%s", libs.JsonMarshal(resp)))
		}
	}
}
