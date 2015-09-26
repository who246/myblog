package main

import (
    _ "myblog/routers"
	_ "myblog/models"
	_ "myblog/filter"
	"github.com/astaxie/beego"
	
)

func main() {
	beego.Run()
}
