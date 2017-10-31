package controllers

import (
	"jupiter/libs"
	"jupiter/libs/couchbase"
	"jupiter/models"

	"github.com/astaxie/beego"
)

type WalletReceiptController struct {
	beego.Controller
}

func (this *WalletReceiptController) Get() {
	// verify token

	id := this.GetString("receipt")
	receipt := models.Receipt{ReceiptId: id}
	receipt.Get(couchbase.Service().BucketReceipt)

	resp := models.PurchaseResponse{}
	if ok := receipt.Verify(); ok {
		resp.Success = &models.PurchaseResponseSuccess{
			Code: 0,
			Data: models.PurchaseData{
				ReceiptId: receipt.ReceiptId,
			},
		}
	} else {
		resp.Error = &models.PurchaseResponseError{
			Code:    "404",
			Message: "invalid parameter.",
		}
	}
	this.Ctx.WriteString(libs.JsonMarshal(resp))
}
