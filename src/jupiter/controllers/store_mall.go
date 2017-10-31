package controllers

import (
	"fmt"
	//	"jupiter/libs"
	"jupiter/models"

	"github.com/astaxie/beego"
)

// controller for test purpose
type TestMallController struct {
	beego.Controller
}

func (this *TestMallController) Get() {
	fmt.Printf("in CartController controller.")
	ticket := this.GetString("ticket")

	this.Data["items"] = models.GetTestItems()

	this.Data["ticket"] = ticket
	//	this.Data["OrderURL"] = libs.GetBillingUserOrder(invoice.Ticket)
	//	this.Data["chargeURL"] = models.GetRoutingURL(models.JAPI_PASSPORT)
	//	this.Data["refreshURL"] = models.GetRoutingURL(models.JFRONT_CART)

	this.TplName = "mall.tpl"
	return
}
