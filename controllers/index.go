package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/who246/myblog/models"
)

type MainController struct {
	ConfigController
}

func (this *MainController) Get() {
	page,err := this.GetInt64("p",1);
	if err != nil {
	beego.Error(err);
	}
	
	extend := map[string]string{};
    tid := this.GetString("tid","");
	key := this.GetString("key");
	if tid != ""{
		extend["type_id"] = tid;
	}
	 
	beego.Info(extend);
	maps,err := models.GetArtPageList(10,(page-1)*10,extend,key);
	
	if err != nil {
	beego.Error(err);
	}
	count, err := models.GetCount(extend,key);
	if(err != nil){
		beego.Error(err);
	}
	p := pagination.NewPaginator(this.Ctx.Request, 10, count)
	beego.Info(p);
	this.Data["paginator"] = p
	types := models.GetAllType();
	this.Data["articles"] = maps
	this.Data["types"] = types
	this.Data["key"] = key
	beego.AppConfig.String("banner1.name");
	
	this.TplNames= "index.tpl";
}
type AboutController struct {
	beego.Controller
}
func (this *AboutController) Get(){
	this.TplNames= "about.tpl";
}
type LoginController struct {
	ConfigController
}
func (this *LoginController) Get() {
   username,ok := this.GetSecureCookie("myblog","username");
   if ok {
	this.Data["username"] = username;
	   beego.Info(username);
   }
   password,ok := this.GetSecureCookie("myblog","password");
   if ok {
	this.Data["password"] = password;
	 beego.Info(password);
   }
   this.TplNames= "login.tpl";
}
func (this *LoginController) Post() {
   username := this.GetString("username","");
   password := this.GetString("password","");
   isRemeber:=this.GetString("isRemeber","");
   if username != beego.AppConfig.String("username") || password != beego.AppConfig.String("password") {
	this.Data["error"] = "用户名或密码❌";
	this.Get();
	return;
   }
    
   this.SetSession("username",username);
   if isRemeber == "on" {
   this.SetSecureCookie("myblog","username",username);
   this.SetSecureCookie("myblog","password",password);
   }
   //this.TplNames= "admin/edit.tpl";
   this.Redirect("/admin/aticle/edit",302)
}
type LogoutController struct {
	beego.Controller
}
func (this *LogoutController) Get() {
   this.DelSession("username");
   beego.Info(this.GetSession("username"))
   this.Redirect("/",302)
}
type ConfigController struct {
	beego.Controller
}
func (this *ConfigController ) Prepare(){
	banner1_url := beego.AppConfig.String("banner1.url")
    banner1_id,_:= beego.AppConfig.Int("banner1.id")
	art := models.Article{Id: banner1_id};
	err := models.GetByProperty(&art);
	if err!=nil{
		beego.Error(err);
	}
	this.Data["banner1_url"]=banner1_url
	this.Data["banner1_art"]=art
	
	banner2_url := beego.AppConfig.String("banner2.url")
    banner2_id,_:= beego.AppConfig.Int("banner2.id")
	art2 := models.Article{Id: banner2_id};
	err2 := models.GetByProperty(&art2);
	if err2 !=nil{
		beego.Error(err);
	}
	this.Data["banner2_url"]=banner2_url
	this.Data["banner2_art"]=art2
 
	banner3_url := beego.AppConfig.String("banner3.url")
    banner3_id,_:= beego.AppConfig.Int("banner3.id")
	art3 := models.Article{Id: banner3_id};
	err3 := models.GetByProperty(&art3);
	if err3 !=nil{
		beego.Error(err);
	}
	this.Data["banner3_url"]=banner3_url
	this.Data["banner3_art"]=art3
}
