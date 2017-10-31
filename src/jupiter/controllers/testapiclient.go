package controllers

import (
	"fmt"
	"jupiter/libs"
	"jupiter/libs/nxapi"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type TestAPIClientController struct {
	beego.Controller
}

func (this *TestAPIClientController) Get() {
	fmt.Printf("in TestAPIClientController....\n")

	// call api service
	delay := this.Input().Get("delay")
	if delay != "" {
		d, _ := strconv.Atoi(delay)
		time.Sleep(time.Duration(d))
	}

	rs, _ := nxapi.Service().GetToken("test_token")
	this.Ctx.WriteString(libs.JsonMarshal(rs))
}
