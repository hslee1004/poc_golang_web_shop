package controllers

import (
	"encoding/json"
	"fmt"
	"jupiter/libs"
	"jupiter/libs/cashbroker"
	"jupiter/libs/couchbase"
	"jupiter/libs/nxapi"
	"jupiter/models"

	"github.com/astaxie/beego"
)

type InvoiceController struct {
	beego.Controller
}

func (this *InvoiceController) Get() {
	fmt.Printf("in invoice controller....\n")

	// token
	//rs, _ := nxapi.Service().GetToken("1111111111")
	//this.Ctx.WriteString(libs.JsonMarshal(rs))

	ticket := this.GetString("ticket")
	invoice := models.Invoice{Ticket: ticket}
	resp := models.InvoiceResponse{}
	if invoice.GetInvoice(couchbase.Service().Bucket) {
		// get balance
		//bal := cashbroker.GetBalance(invoice.UserId)
		bal := cashbroker.GetWallets(invoice.UserId)
		invoice.Balances = bal
		resp.Success = &models.InvoiceResponseDetail{
			Code:    0,
			Invoice: invoice,
		}
	} else {
		resp.Error = &models.NXAPIErrorDetail{
			Code:    "404",
			Message: "invalid invoice.",
		}
		resp.Success = nil
	}
	this.Ctx.WriteString(libs.JsonMarshal(resp))
}

func (this *InvoiceController) Post() {
	var invoice models.Invoice
	fmt.Printf("invoice ctrl body:%s\n", this.Ctx.Input.RequestBody)
	json.Unmarshal(this.Ctx.Input.RequestBody, &invoice)
	fmt.Printf("invoice: %v\n", invoice)

	// get invoicy
	resp := models.APIResponse{}
	if invoice.ProductId != "" {
		// get user_id
		if user, ok := nxapi.Service().GetUserId(invoice.Token); ok {
			invoice.UserId = user.UserId
			// create a ticket
			// todo : save ticket to
			if ticket, ok := nxapi.Service().CreateTicket(invoice.Token, nxapi.PRODUCT_ID_BILLING); ok {
				invoice.Ticket = ticket.Ticket
				invoice.InvoiceId = libs.GetUUID() // new
				if err := invoice.RegisterInvoice(couchbase.Service().Bucket); err == nil {
					resp.Success = &models.NXAPIDetail{
						Code: 0,
						Data: models.NXAPIData{
							Ticket:     invoice.Ticket,
							InvoiceId:  invoice.InvoiceId,
							BillingURL: libs.GetBillingUserCart(invoice.Ticket, invoice.InvoiceId),
							ForDevInvoiceURL: models.GetQueryURL(models.GetRoutingURL(models.JAPI_INVOICE),
								map[string]string{"ticket": invoice.Ticket, "invoice_id": invoice.InvoiceId}),
						},
					}
					resp.Error = nil
				} else {
					resp.Error = &models.NXAPIErrorDetail{
						Code:    "500",
						Message: err.Error(),
					}
					resp.Success = nil
				}
			}
		}
	} else {
		// https://nexonusa.atlassian.net/wiki/pages/viewpage.action?pageId=19366327
		resp.Success = nil
		resp.Error = &models.NXAPIErrorDetail{
			Code:    "400",
			Message: "no product_id",
		}
	}
	fmt.Printf("%s\n", libs.JsonMarshal(resp))
	this.Ctx.WriteString(libs.JsonMarshal(resp))
}
