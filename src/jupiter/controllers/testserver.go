package controllers

import (
	"fmt"
	bc "jupiter/controllers/baseController"
	"jupiter/libs"
	"jupiter/libs/jsapi"
	"jupiter/libs/nxapi"
	"jupiter/models"
	"strconv"
)

type TestServerController struct {
	//beego.Controller
	bc.BaseController
}

//
// http://localhost/shop/cart?item=10001&item=10002&item=10003
//
func (this *TestServerController) Get() {
	fmt.Printf("in TestServerController....\n")
	ticket := this.Input().Get("ticket")
	selected := this.GetStrings("item")
	fmt.Printf("ticket:%s\nselected items:%v\n", ticket, selected)

	// get selected items
	items := models.GetSelectedTestItems(selected)
	fmt.Printf("selected items array:%v\n", items)

	// get token
	api := nxapi.Service()
	token, _ := api.GetToken(ticket)

	// get user profile
	if user, ok := api.GetTokenDetail(token.Token); ok {
		// defind purchase items
		//		items := [...]models.Item{
		//			{Id: "10002", Name: "test_jupiter_item_2", Price: 1, Qty: 1},
		//			{Id: "10003", Name: "test_jupiter_item_3", Price: 2, Qty: 1},
		//		}
		//		fmt.Printf("items: %s\n", libs.JsonMarshal(items))

		var total float32
		for i := range items {
			items[i].Sum(&total)
		}

		payMethods := [...]models.PaymentMethod{{Name: "nx_prepaid"}, {Name: "nx_credit"}}

		invoice := &models.Invoice{
			ProductId:            user.Success.Data.ProdId,
			UserNo:               strconv.Itoa(user.Success.Data.UserNo),
			UserIP:               libs.GetMyLocalIP(), //test:"192.168.14.130"//this.GetClientIp(),
			Date:                 libs.GetNow(),
			Items:                items[:],
			TotalPrice:           total,
			Token:                token.Token,
			RedirectUri:          "http://localhost/game/server/redirect",
			AllowedPaymentMethod: payMethods[:], // nx_prepaid
			OptionUseReceiptFlow: true,
		}

		// call jupiter api server
		resp, _ := jsapi.Service().RegisterInvoice2(invoice)

		// redirect to billing url

		this.Ctx.WriteString(libs.JsonMarshal(resp))
	} else {
		this.Ctx.WriteString(fmt.Sprintf("error - token detai: token:%s", libs.JsonMarshal(token)))
	}
}
