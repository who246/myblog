package cache
import (  
    "fmt"
    "github.com/astaxie/beego/cache"  
    _"github.com/astaxie/beego/cache/redis" 
	"github.com/astaxie/beego"  
)
 var Redis cache.Cache
func init(){  
var err error;  
	pwd := beego.AppConfig.String("redis.pwd")
	ip:= beego.AppConfig.String("redis.ip")
	port,_ := beego.AppConfig.Int("redis.port")
	key:= beego.AppConfig.String("redis.key")

 Redis, err = cache.NewCache("redis", fmt.Sprintf(`{"conn":"%s:%d", "key":"%s","password":"%s"}`, ip, port,key, pwd))  
if err != nil {  
    beego.Error(err) 
}
 
}
 