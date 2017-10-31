package controllers

import (
	//"encoding/json"
	"fmt"
	"jupiter/libs"
	"jupiter/libs/couchbase"
	//"jupiter/libs/nxapi"
	"jupiter/libs/cashbroker"
	"jupiter/models"

	"github.com/astaxie/beego"
)

type StoreCartController struct {
	beego.Controller
}

//
// http://api-beta.nexon.net/store/front/store/cart?ticket=291a7581-30bd-4260-bac1-3ab39112537d&
// invoice_id=22c10d4f-ef57-44e7-be0a-4369b3542601
//
func (this *StoreCartController) Get() {
	fmt.Printf("in CartController controller....\n")
	// get purchase ticket
	ticket := this.GetString("ticket")

	// get invoice
	invoice := models.Invoice{Ticket: ticket}
	if invoice.GetInvoice(couchbase.Service().Bucket) {
		fmt.Printf("[cart] invoice:%v\n", invoice)
		//		it := invoice.Items
		//		var items []models.Item
		//		libs.JsonUnmarshal([]byte(it), &items)
		//		fmt.Printf("items: %v\n", items)
		bal := cashbroker.GetBalance(invoice.UserId)
		fmt.Printf("cash type:%d\v", invoice.OptionCashType)

		this.Data["invoice"] = &invoice
		//		this.Data["items"] = &items
		this.Data["items"] = invoice.Items
		this.Data["balance"] = &bal
		this.Data["ticket"] = invoice.Ticket
		this.Data["OrderURL"] = libs.GetBillingUserOrder(invoice.Ticket, invoice.Ticket) // ticket, invoice_id
		this.Data["chargeURL"] = models.GetRoutingURL(models.JAPI_PASSPORT)
		this.Data["refreshURL"] = models.GetRoutingURL(models.JFRONT_CART)

		this.TplName = "cart.tpl"
		return
	}
	this.Ctx.WriteString("error..")
}
