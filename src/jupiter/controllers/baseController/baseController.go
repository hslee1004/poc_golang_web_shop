// baseController:
// https://github.com/goinggo/beego-mgo/blob/master/controllers/buoyController.go

package baseController

import (
	"strings"

	"github.com/astaxie/beego"
)

type (
	BaseController struct {
		beego.Controller
	}
)

func (this *BaseController) GetClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
