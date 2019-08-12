package main

import (
	_ "./router"
	"github.com/astaxie/beego"
)

const (
	APP_VER = "0.1.1.0227"
)
//运行起后台服务器
func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)
	beego.Run()
}
