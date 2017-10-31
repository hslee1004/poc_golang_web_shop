package controllers

import (
	"fmt"
	//	"jupiter/libs"
	"jupiter/libs/couchbase"
	"jupiter/libs/nxapi"
	"jupiter/models"
	"time"

	"github.com/astaxie/beego"
)

type PurchaseNXController struct {
	beego.Controller
}

func (this *PurchaseNXController) Get() {
	fmt.Printf("in invoice controller....\n")
	ticket := this.GetString("ticket")
	invoice := models.Invoice{Ticket: ticket}

	if invoice.GetInvoice(couchbase.Service().Bucket) {
		fmt.Printf("[passport] invoice:%v\n", invoice)
		resp, ok := nxapi.Service().GetPassport(invoice.Token, invoice.UserIP)
		if ok {
			this.Ctx.SetCookie("NPPv2", resp.Passport)
			this.Ctx.SetCookie("authToken", resp.AuthToken)
			maxAge := 60 * 10
			this.Ctx.ResponseWriter.Header().Add("Set-Cookie",
				fmt.Sprintf("NPPv2=%s; Domain=nexon.net; Path=/; Expires=%s; Max-Age=%d",
					resp.Passport,
					time.Now().Add(time.Duration(maxAge)*time.Second).UTC().Format(time.RFC1123),
					maxAge,
				))
		}
		//this.Ctx.WriteString(libs.JsonMarshal(resp))
	} else {
		//this.Ctx.WriteString("ticket error.")
	}
	this.TplName = "nx.tpl"
}
