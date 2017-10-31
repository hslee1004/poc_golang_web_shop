package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type TestAsyncController struct {
	beego.Controller
}

func (this *TestAsyncController) Get() {
	fmt.Printf("in get....\n")
	intVal, _ := this.GetInt64("int")
	var sum int64 = 1
	for sum < intVal {
		sum += 1
		time.Sleep(1)
	}
	rs := fmt.Sprintf("{\"value\":\"%d\"}", intVal)
	this.Ctx.WriteString(rs)
}
