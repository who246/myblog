package utils
import(
	"github.com/astaxie/beego/orm"
	)
	
	
func NewOrm() orm.Ormer{
		return orm.NewOrm()
	}
