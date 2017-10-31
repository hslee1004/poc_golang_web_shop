package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type PingController struct {
	beego.Controller
}

func (this *PingController) Get() {
	rs := fmt.Sprintf("{\"ping\":\"%s\"}", "ok")
	this.Ctx.WriteString(rs)
}
