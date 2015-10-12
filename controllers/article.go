package controllers

import (
	"github.com/astaxie/beego"
	"github.com/who246/myblog/models"
    "github.com/who246/myblog/cache"
	 "github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"encoding/json"
	 
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
	art := getArticle(id)
	//计数
	addClickNum(id);
	
	this.Data["article"] = art;
	types := models.GetAllType();	
	this.Data["types"] = types
	this.Data["headStr"] = art.Title
	this.TplNames= "view.tpl";
}
func getArticle(id int)(art models.Article){
	idstr := strconv.Itoa(id);
	art_obj := cache.Redis.Get("article:"+idstr+":view")
	var err error
	beego.Error(art_obj)
	if art_obj == nil{
    art = models.Article{Id: id};
	err = models.GetByProperty(&art);
	if err!=nil{
		beego.Error(err);
	 }
	bytes,_:= json.Marshal(art);
	cache.Redis.Put("article:"+idstr+":view",bytes,60*10) 
	return art
	} 
	bytes,err := redis.Bytes(art_obj,err)
	err = json.Unmarshal(bytes,&art)
	if err != nil {
		beego.Error(err)
	} 
	 
	return art;
}
func addClickNum(id int){
	idstr := strconv.Itoa(id);
	click_num_obj  :=	cache.Redis.Get("article:"+idstr+":click_num")
	var click_num int;
	var err error
    //没有数据获取数据库中的
	if click_num_obj == nil {
		art2 := models.Article{Id: id};
	    err = models.GetByProperty(&art2);
		if(err != nil){
			beego.Error(err);
		}
		click_num = art2.ClickNum;
	}else{
		click_num , _ = redis.Int(click_num_obj,err);
	}
	click_num ++;
	if(click_num%10 == 0){
		if  _,err = models.AddClickNum(id,click_num);err != nil {
			cache.Redis.Delete("article:"+idstr+":click_num")
		}
	}else{
	cache.Redis.Put("article:"+idstr+":click_num",click_num,99999)
	}
}