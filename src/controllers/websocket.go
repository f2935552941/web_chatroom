// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	beego.Controller
}

// 处理get请求
func (this *WebSocketController) Get() {
	// Safe check.
	uname := this.GetString("uname")
	roomid := this.GetString("roomid")
	if len(uname) == 0 || len(roomid) == 0 || len(roomid) > 10 {
		this.Redirect("/", 302)
		return
	}
	this.TplName = "chatRoom.html"
	this.Data["username"] = uname
	this.Data["roomid"] = roomid
}

// 处理Join请求
func (this *WebSocketController) Join() {
	uname := this.GetString("uname")
	roomid, err := this.GetInt64("roomid")
	if len(uname) == 0 || err != nil {
		this.Redirect("/", 302)
		return
	}
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup websocket connection:", err)
		return
	}
	Join(uname, roomid, ws)
	Room[roomid] += 1
	defer Leave(uname, roomid)
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		publish <- newEvent(models.EVENT_MESSAGE, uname, roomid, string(p))
	}
}

// 发送处理事件给前端
func broadcastWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}
