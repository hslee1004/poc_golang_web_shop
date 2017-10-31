package controllers

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"jupiter/libs"
	"jupiter/libs/cashbroker"
	"jupiter/libs/couchbase"
	"jupiter/libs/nxapi"
	"jupiter/models"

	"github.com/astaxie/beego"
)

type WalletPurchaseController struct {
	beego.Controller
}

func (this *WalletPurchaseController) Post() {
	var trx models.PurchaseRequest
	fmt.Printf("body:%s\n", this.Ctx.Input.RequestBody)
	json.Unmarshal(this.Ctx.Input.RequestBody, &trx)
	fmt.Printf("%v\n", trx)

	// verify token

	// get invoice
	invoice := models.Invoice{Ticket: trx.PurchaseTicket}
	invoice.GetInvoice(couchbase.Service().Bucket)

	// product_id : service code

	// verify invoice
	resp := models.PurchaseResponse{}
	if ok := invoice.Verify(); ok {
		if pay, ok := cashbroker.RequestPayment(libs.GetServiceCode("20000"), &invoice); ok {
			if invoice.OptionUseReceiptFlow {
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
							//PurchaseTicket: trx.PurchaseTicket,
							ReceiptId: receipt.ReceiptId,
							CallbackAPI: models.GetQueryURL(models.GetRoutingURL(models.JAPI_RECEIPT),
								map[string]string{"receipt_id": receipt.ReceiptId}),
						},
					}

				} else {
					// commit error
					resp.Error = &models.PurchaseResponseError{
						Code:    commit.Response.Result,
						Message: "payment error.",
					}
				}

			} else {
				// get new ticket
				//if ticket, ok := nxapi.Service().CreateTicket(invoice.Token); ok {
				if ticket, ok := nxapi.Service().CreateTicket(invoice.Token, nxapi.PRODUCT_ID_BILLING); ok {
					invoice.Ticket = ticket.Ticket
					invoice.Save(couchbase.Service().Bucket) // save trx id
					fmt.Printf("%s", pay)
					// PurchaseResponse
					resp.Success = &models.PurchaseResponseSuccess{
						Code: 0,
						Data: models.PurchaseData{
							TransactionId: invoice.Ticket, // new ticket
							CallbackAPI: models.GetQueryURL(
								models.GetRoutingURL(models.JAPI_TRANSACTION),
								map[string]string{"transaction_id": invoice.Ticket},
							),
						},
					}
				} else {
					resp.Error = &models.PurchaseResponseError{
						Code:    "401",
						Message: "old credential.",
					}
				}
			}

		} else {
			resp.Error = &models.PurchaseResponseError{
				Code:    pay.Response.Result,
				Message: "invalid invoice.",
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
