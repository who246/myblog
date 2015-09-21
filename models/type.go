package models

import(
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
)
type Type struct{
    Id int
    CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	ModifyTime time.Time  `orm:"auto_now;type(datetime)"`
	Name string
	Nums int
}
func init(){
	orm.RegisterModelWithPrefix("goblog_", new(Type))
}
func GetAllType() (t1 []Type) {
	var t []Type
	_,err := orm.NewOrm().QueryTable("goblog_type").OrderBy("-CreateTime").All(&t);
	beego.Error(err);
	return t;
}
 
