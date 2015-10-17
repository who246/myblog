package admin

import (
	"github.com/astaxie/beego"
	"github.com/who246/myblog/models"
	"github.com/who246/myblog/layout"
	"github.com/who246/myblog/utils"
	"github.com/astaxie/beego/validation"
    "strings"
	"strconv"
	"os/exec" 
)


type AticleController struct {
	layout.AdminLayoutController
}
func (this *AticleController) Edit() {
	 
    types := models.GetAllType();	
	this.Data["types"] = types
	this.Data["art"] = &models.Article{}
	this.Data["op"]="add"
	//this.TplN""ames= "admin/edit.tpl";
	this.GetLayout("admin/edit.tpl");
}
func (this *AticleController) Add(){
	 
	var article models.Article;
	valid := validation.Validation{};
	title := this.GetString("title","");
	content := this.GetString("content","");
	typeId,_ :=this.GetInt("typeId");
	article.Content = content;
	article.Title = title;
	article.TypeId = typeId;
	this.Data["art"] = article; 
	this.GetLayout("admin/edit.tpl");
	b, err := valid.Valid(&article);
	if err != nil {
       beego.Error(err);
    }
    if !b {         
		this.Data["err"] = valid.Errors;
		return
    }
	
	models.InsertArt(&article);
	types := models.GetAllType();	
	this.Data["types"] = types
	article.Content = "";
	article.Title = "";
	article.TypeId = -1;
	this.Data["art"] = article 
	
}
func (this *AticleController) Manage() {
	this.GetLayout("admin/manage.tpl");
}
func (this *AticleController) Del() {
	idstr := this.GetString("ids","");
	if idstr == "" {
		this.ToJsonFails("id不能为空");
		return;
	}
	var msg string = "";
	ids := strings.Split(idstr,",");
	for _, i := range ids {
		id,err := strconv.Atoi(i);
		if err!=nil {
		    continue;
		}
		 _, err2 := utils.NewOrm().Delete(&models.Article{Id: id})
		if err2 == nil{
			msg = msg+strconv.Itoa(id)+" "
		}
	}
	if msg == "" { 
 	  this.ToJsonFails("删除失败");
      return;
	}
	this.ToJsonSuccess("成功删除Id："+msg);
	 
}
func (this *AticleController) ShowModify() {
	id,err :=this.GetInt("id");
	if err!=nil {
		this.ToJsonFails("请输入正确的id");
		return;
		}
    types := models.GetAllType();	
	art := &models.Article{Id: id};
	models.GetByProperty(art)
	this.Data["types"] = types
	this.Data["art"] = art
	this.Data["op"]="modify"
	this.GetLayout("admin/edit.tpl");
}
func (this *AticleController) Modify() {
	art := models.Article{}
	if err := this.ParseForm(&art); err != nil {
        beego.Error(err)
    
		return
    }
	beego.Info(art);
	if _,err := utils.NewOrm().Update(&art,"Title","Content","TypeId","ModifyTime"); err != nil {
	types := models.GetAllType();	 
	models.GetByProperty(&art)
	this.Data["types"] = types
	this.Data["art"] = art
	this.Data["op"]="modify"
	this.GetLayout("admin/edit.tpl");
	return;
	} 
	this.Manage()
}
type TypeController struct {
	layout.AdminLayoutController
}
func  (this *TypeController)  Edit(){
	this.GetLayout("admin/type_edit.tpl");
}
func  (this *TypeController)  Add(){
	 name := this.GetString("name","");
	if name == "" {
		this.ToJsonFails("添加失败，请输入名称");
		return;
	}
    _,err := utils.NewOrm().Insert(&models.Type{Name:name,Nums:0});
	if err != nil{
		beego.Error(err);
       this.ToJsonFails("添加失败");
		return 
	}
		this.ToJsonSuccess("添加成功");
}
func  (this *TypeController)  Del(){
	idstr := this.GetString("ids","");
	if idstr == "" {
		this.ToJsonFails("id不能为空");
		return;
	}
	var msg string = "";
	ids := strings.Split(idstr,",");
	for _, i := range ids {
		id,err := strconv.Atoi(i);
		if err!=nil {
		    continue;
		}
		 _, err2 := utils.NewOrm().Delete(&models.Type{Id: id})
		if err2 == nil{
			msg = msg+strconv.Itoa(id)+" "
		}
	}
	if msg == "" { 
 	  this.ToJsonFails("删除失败");
      return;
	}
	
	this.ToJsonSuccess("成功删除Id："+msg);
}
type ConfigController struct {
	layout.AdminLayoutController
}
func  (this *ConfigController)  Set(){
	n := this.GetString("n","1");
	url := this.GetString("url","");
	id := this.GetString("id","");
	b_url := "banner"+n+".url";
	b_id := "banner"+n+".id"; 
	if url!="" {
		err := beego.AppConfig.Set(b_url,url);
		if(err != nil)  {
		this.ToJsonSuccess("更改失败");
		return;
		}
	}
	if id!="" {
		err := beego.AppConfig.Set(b_id,id);
		if(err != nil)  {
		this.ToJsonSuccess("更改失败");
		return;
		}
	} 
	err := beego.AppConfig.SaveConfigFile("app.conf");
	if(err == nil)  {
	this.ToJsonFails("更改成功")
	} else{
 	this.ToJsonSuccess("更改失败");
	}
}
//
type ConstructController struct {
	layout.AdminLayoutController
}
func  (this *ConstructController)  Update(){
	cmd := exec.Command("sh", "update.sh")
    err := cmd.Start();
    if err != nil {
        panic(err.Error())
    }
 
   
    if err := cmd.Wait(); err != nil {
        panic(err.Error())
    }
    
	this.Data["json"]="等待重启";
	this.ServeJson();
}
func  (this *ConstructController)  Bulid(){
	cmd := exec.Command("sh", "bulid.sh")
    err := cmd.Start();
    if err != nil {
        panic(err.Error())
    }
 
   
    if err := cmd.Wait(); err != nil {
        panic(err.Error())
    }
	this.Data["json"]="等待重启";
	this.ServeJson();
}
