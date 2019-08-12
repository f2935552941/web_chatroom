package controllers

import "github.com/astaxie/beego"

//调用beego中的默认controller模块
type AppController struct {
	beego.Controller
}

// 处理get请求 显示 login.html页面
func (this *AppController) Get() {
	this.TplName = "login.html"
}

// 输入用户名后登录Join请求
func (this *AppController) Join() {
	// Get form value.
	//sess := StartSession()
	uname := this.GetString("uname")

	// Check valid.
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}
	this.Redirect("/index?uname="+uname, 302) //重定向index页面
	return
}
