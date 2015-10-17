package models

import (
	"github.com/astaxie/beego/orm"
	"time"
  
)

type Article struct {
    Id int `form:"id" redis:"id"`
    CreateTime time.Time `orm:"auto_now_add;type(datetime)" redis:"createTime"`
	ModifyTime time.Time  `orm:"auto_now;type(datetime)" redis:"modifyTime"`
	Title   string 	`form:"title" valid:"Required";orm:"size(60)" redis:"title"`
	//Title   string 	`orm:"size(60)" valid:"Required"`
	Content string  `form:"content" valid:"Required" orm:"type(text);null" redis:"content"` 
    TypeId    int `form:"typeId" valid:"Required" redis:"typeId"`
	ClickNum  int `redis:"clickNum"`
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
func AddClickNum(id ,num int) (count int64,err error){
	return orm.NewOrm().Update(&Article{Id: id,ClickNum:num},"ClickNum")
}