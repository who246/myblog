package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
 
)

type Article struct {
    Id int `form:"id"`
    CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	ModifyTime time.Time  `orm:"auto_now;type(datetime)"`
	Title   string 	`form:"title";valid:"Required;",orm:"size(60)";`
	Content string  `form:"content";valid:"Required;",orm:"column(content)"` 
    TypeId    int `form:"typeId";valid:"Required;"`
	
}
 
func init() {
	orm.RegisterModelWithPrefix("goblog_", new(Article))
}
func InsertArt(m *Article) {
	o := orm.NewOrm()
	o.Begin();
	_type := Type{Id: m.TypeId}
	if o.Read(&_type) == nil {
    _type.Nums = _type.Nums+1;
    if _, err :=  o.Update(&_type,"Nums"); err != nil {
		o.Rollback();
		return;
	}
	
	}
	if _, err := o.Insert(m); err != nil {
		 o.Rollback();
		return;
	}
	 o.Commit();
}
func   GetArtPageList(pageSize,pageStart int64,extend map[string] string,title string) ([]orm.Params, error)  {
	beego.Info(extend);
	o := orm.NewOrm()
    qs := o.QueryTable("goblog_article")
     for k, v:= range extend{
     qs= qs.Filter(k,v);
    }
	if title != "" {
		qs=qs.Filter("title__contains",title);
	}
	var params []orm.Params
	if _,err := qs.OrderBy( "-modify_time").Limit(pageSize,pageStart).Values(&params); err !=nil {
		return nil,err
	}
    
	return params,nil
}
func GetCount(extend map[string] string,title string) (count int64,err error){
	o := orm.NewOrm()
	qs := o.QueryTable("goblog_article")
     for k, v:= range extend{
     qs= qs.Filter(k,v);
    }
	if title != "" {
		qs=qs.Filter("title__contains",title);
	}
	count, err = qs.Count()
	return count,err;
}
func GetByProperty(art * Article) (err error){
	return orm.NewOrm().Read(art);
}