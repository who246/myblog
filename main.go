package main

import (
    _ "github.com/who246/myblog/routers"
	 "github.com/who246/myblog/models"
	_ "github.com/who246/myblog/filter"
	_ "github.com/who246/myblog/cache"
	"github.com/astaxie/beego"
	
)

func main() {

    models.CreateTable()
	beego.Run()
}
