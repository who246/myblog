package models



import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	
)


func init() {
	
	user := beego.AppConfig.String("db.user")
	pwd := beego.AppConfig.String("db.pwd")
	ip:= beego.AppConfig.String("db.ip")
	port,_ := beego.AppConfig.Int("db.port")
	dbname := beego.AppConfig.String("db.dbname")
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pwd,ip, port, dbname))
    
}
func CreateTable() {
    name := "default"                          //数据库别名
    force := false                             //不强制建数据库
    verbose := true                            //打印建表过程
    err := orm.RunSyncdb(name, force, verbose) //建表
    if err != nil {
        beego.Error(err)
    }
}