package controllers

import (
	"fmt"
	"jupiter/libs"
	"jupiter/libs/cashbroker"

	"github.com/astaxie/beego"
)

type BalanceController struct {
	beego.Controller
}

func (this *BalanceController) Get() {
	fmt.Printf("in BalanceController....\n")
	// get balance
	user_id := this.Input().Get("user_id")
	fmt.Printf("body:%s", user_id)
	v := cashbroker.GetBalance(user_id)
	//this.Ctx.WriteString(v)
	this.Ctx.WriteString(libs.JsonMarshal(v))
}

/*
func (this *InvoiceController) Post() {
	var invoice models.Invoice
	fmt.Printf("body:%s", this.Ctx.Input.RequestBody)
	json.Unmarshal(this.Ctx.Input.RequestBody, &invoice)
	fmt.Printf("%s\n", invoice)
	if invoice.ProductID != "" {
		invoice.RegisterInvoice(couchbase.Service().Bucket)
		//uid := models.AddUser(user)
		//u.Data["json"] = map[string]string{"uid": uid}
		//u.ServeJSON()
	}
	this.Ctx.WriteString(libs.JsonToString(invoice))
}
*/
