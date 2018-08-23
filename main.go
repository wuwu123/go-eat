package main

import (
	_ "eat/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"eat/common"
)

func init()  {
	common.RegistInit()
}
func main() {
	orm.Debug = true
	beego.Run()
}

