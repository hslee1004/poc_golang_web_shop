package controllers

import (
	//"encoding/json"
	"fmt"
	//	"jupiter/libs"
	"jupiter/libs/couchbase"
	//"jupiter/libs/nxapi"
	"jupiter/libs/cashbroker"
	"jupiter/libs/jsapi"
	"jupiter/models"

	"github.com/astaxie/beego"
)

type StoreOrderController struct {
	beego.Controller
}

func (this *StoreOrderController) Get() {
	ticket := this.GetString("ticket")
	cashType := this.GetString("cash_type")
	// rule id : prepaid, credit
	fmt.Printf("ticket: %s, cash_type:%s\n", ticket, cashType)

	// get invoice
	invoice := models.Invoice{Ticket: ticket}
	if invoice.GetInvoice(couchbase.Service().Bucket) {
		fmt.Printf("[StoreOrder] invoice:%v\n", invoice)
		// TODO : change this part
		//it := invoice.Items
		//		var items []models.Item
		//		libs.JsonUnmarshal([]byte(it), &items)
		//		fmt.Printf("items: %v\n", items)
		//
		bal := cashbroker.GetBalance(invoice.UserId)
		fmt.Printf("balance: %v\n", bal)
		fmt.Printf("invoice: %v\n", invoice)
		// verify balance
		//if models.CASHTYPE_ALL == invoice.CashType {
		//}
		if bal.HasEnoughBalance(&invoice) {
			fmt.Printf("user has enough balance..\n")
			// call complete order on jupiter
			if resp, ok := jsapi.Service().CompleteTransaction(&invoice); ok {
				if resp.Success != nil {
					// return success page
					//this.Ctx.WriteString(libs.JsonMarshal(resp))
					// redirect to callback url
					this.Redirect(
						models.GetQueryURL(
							invoice.RedirectUri,
							map[string]string{
								"receipt_id":     resp.Success.Data.ReceiptId,
								"transaction_id": resp.Success.Data.TransactionId,
								"callback_api":   resp.Success.Data.CallbackAPI,
							}),
						302)
					return

				}
			} else {
				this.Ctx.WriteString("error on CompleteTransaction..")
			}
			return
		} else {
			fmt.Printf("not enough balance..\n")
		}
	}
	this.Ctx.WriteString("can not find invoice.")
}
