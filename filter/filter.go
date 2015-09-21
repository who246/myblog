package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var LoginUrlManager = func(ctx *context.Context) {
     username, ok := ctx.Input.Session("username").(string)
	if !ok {
		ctx.Redirect(302, "/login")
	}
	ctx.Input.SetData("username",username);
}

func init() {
  beego.InsertFilter("/admin/*",beego.BeforeRouter,LoginUrlManager)
}

