package routers

import (
	"github.com/who246/myblog/controllers"
	"github.com/who246/myblog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
  	beego.Router("/", &controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/about",&controllers.AboutController{})
	beego.Router("/admin/logout",&controllers.LogoutController{})
	beego.AutoRouter(&controllers.ArticleController{})
	
	beego.AutoPrefix("/admin",&admin.AticleController{})
    beego.AutoPrefix("/admin",&admin.TypeController{})
    beego.AutoPrefix("/admin",&admin.ConfigController{})
	beego.AutoPrefix("/admin",&admin.ConstructController{})
}
