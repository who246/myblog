package layout
import (
     
	"github.com/astaxie/beego"
	
)
type AdminLayoutController struct {
	beego.Controller
}
func (this *AdminLayoutController)GetLayout(cotentTpl string ){
	this.Layout = "admin/layout.tpl"
    this.TplNames = cotentTpl
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["head"] = "admin/head.tpl"
    this.LayoutSections["left"] = "admin/left.tpl"
}

func (this *AdminLayoutController)toJson(msg ,status string){
	this.Data["json"] = map[string]string{"msg":msg,"status":status} 
	this.ServeJson();
}

func (this *AdminLayoutController)ToJsonSuccess(msg  string){
	this.toJson(msg,"0");
}
func (this *AdminLayoutController)ToJsonFails(msg  string){
	this.toJson(msg,"1");
}