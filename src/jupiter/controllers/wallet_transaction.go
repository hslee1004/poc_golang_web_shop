package controllers

import (
	//"encoding/json"
	"fmt"
	"jupiter/libs"
	"jupiter/libs/couchbase"
	//"jupiter/libs/nxapi"
	"encoding/json"
	"jupiter/libs/cashbroker"
	"jupiter/models"

	"github.com/astaxie/beego"
)

type WalletTrxController struct {
	beego.Controller
}

func (this *WalletTrxController) Post() {
	var trxReq models.TransactionRequest
	fmt.Printf("body:%s\n", this.Ctx.Input.RequestBody)
	json.Unmarshal(this.Ctx.Input.RequestBody, &trxReq)
	fmt.Printf("%v\n", trxReq)

	// get invoice
	//invoice := models.Invoice{Ticket: trxReq.PurchaseTicket}
	invoice := models.Invoice{Ticket: trxReq.TransactionId}
	invoice.GetInvoice(couchbase.Service().Bucket)

	resp := models.PurchaseResponse{}
	if ok := invoice.Verify(); ok {
		if commit, ok := cashbroker.CommitPayment(&invoice); ok {
			receipt := &models.Receipt{
				ReceiptId: commit.Response.PaymentNo,
				Invoice:   &invoice,
			}
			// save receipt to persistence
			receipt.Save(couchbase.Service().BucketReceipt)
			resp.Success = &models.PurchaseResponseSuccess{
				Code: 0,
				Data: models.PurchaseData{
					PurchaseTicket: trxReq.PurchaseTicket, // TODO : gen new ticket
					ReceiptId:      receipt.ReceiptId,     // cashbroker PaymentNo
				},
			}
		}
	} else {
		resp.Error = &models.PurchaseResponseError{
			Code:    "404",
			Message: "invalid invoice.",
		}
	}
	this.Ctx.WriteString(libs.JsonMarshal(resp))

}
