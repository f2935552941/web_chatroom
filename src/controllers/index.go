package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"

	"github.com/astaxie/beego"
)

//默认beego的controller
type IndexController struct {
	beego.Controller
}

//储存当前每个房间的人数
var Room = make(map[int64]int)

//记录房间id以及人数 方便给前端传递数据
type Roomlist struct {
	Roomid int64
	PerNum int
}

//处理基本get请求
func (this *IndexController) Get() {
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}
	this.TplName = "index.html"
	this.Data["username"] = uname
}

//用户输入或者点击房间后加入房间
func (this *IndexController) Join() {
	//roomid=this.GetString("roomid")
	uname := this.GetString("uname")
	roomid, _ := this.GetInt64("roomid")
	fmt.Println(uname)
	fmt.Println(roomid)
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	url := "/ws?uname=" + uname + "&roomid=" + strconv.FormatInt(roomid, 10)
	this.Redirect(url, 302)
	return
}

//处理前端websocket发送的数据
func (this *IndexController) Check() {
	fmt.Printf("error")
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup websocket connection:", err)
		return
	}
	//发送当前所有房间信息给前端
	for k, v := range Room {
		fmt.Printf("%d %d", k, v)
		if v == 0 {
			continue
		}
		data, err := json.Marshal(Roomlist{k, v})
		if err != nil {
			beego.Error("Fail to marshal event:", err)
			return
		}
		ws.WriteMessage(websocket.TextMessage, data)
	}
}
