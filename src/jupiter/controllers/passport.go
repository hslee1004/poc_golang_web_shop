package controllers

import (
	//	"encoding/json"
	"fmt"
	"jupiter/libs"
	"jupiter/libs/couchbase"
	"jupiter/libs/nxapi"
	"jupiter/models"

	"github.com/astaxie/beego"
)

type PassportController struct {
	beego.Controller
}

func (this *PassportController) Get() {
	fmt.Printf("in invoice controller....\n")
	ticket := this.GetString("ticket")
	invoice := models.Invoice{Ticket: ticket}
	if invoice.GetInvoice(couchbase.Service().Bucket) {
		fmt.Printf("[passport] invoice:%v\n", invoice)
		resp, _ := nxapi.Service().GetPassport(invoice.Token, invoice.UserIP)
		this.Ctx.WriteString(libs.JsonMarshal(resp))
	} else {
		this.Ctx.WriteString("ticket error.")
	}
}

func (this *PassportController) Post() {
	//	var invoice models.Invoice
	//	fmt.Printf("invoice ctrl body:%s\n", this.Ctx.Input.RequestBody)
	//	json.Unmarshal(this.Ctx.Input.RequestBody, &invoice)
	//	fmt.Printf("invoice: %v\n", invoice)

	//	// get invoicy
	//	resp := models.APIResponse{}
}
