package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	// Register routers.
	beego.Router("/", &controllers.AppController{})
	// 处理login页面登录请求
	beego.Router("/join", &controllers.AppController{}, "post:Join")

	//处理index基本get
	beego.Router("/index", &controllers.IndexController{})
	//处理index界面post请求
	beego.Router("/create", &controllers.IndexController{}, "post:Join")
	//返回新的index页面
	beego.Router("/index/true", &controllers.IndexController{}, "get:Check")

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

}
