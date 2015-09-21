package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
 
	"strconv"
	"strings"
)

type ArticleController struct {
	beego.Controller
}

 
func (this *ArticleController) View(){
	idstr :=this.Ctx.Input.Param("0");
	
	id, err := strconv.Atoi(strings.Split(idstr,".")[0]);
	if err!=nil{
		beego.Error(err);
	}
	art := models.Article{Id: id};
	err = models.GetByProperty(&art);
	if err!=nil{
		beego.Error(err);
	}
	this.Data["article"] = art;
	types := models.GetAllType();	
	this.Data["types"] = types
	this.Data["headStr"] = art.Title
	this.TplNames= "view.tpl";
}