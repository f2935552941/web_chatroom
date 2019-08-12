package controllers

import (
	"container/list"
	"time"

	"../models"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

//储存每个用户信息以及所拥有的websocket
type Subscriber struct {
	Name   string
	Roomid int64
	Conn   *websocket.Conn // Only for WebSocket users; otherwise nil.
}

var (
	//新用户
	subscribe = make(chan Subscriber, 10)
	//退出的用户
	unsubscribe = make(chan string, 10)
	//发送事件及请求
	publish = make(chan models.Event, 10)
	//储存所有当前用户
	subscribers = list.New()
)

func newEvent(ep models.EventType, user string, roomid int64, msg string) models.Event {
	return models.Event{ep, user, roomid, int(time.Now().Unix()), msg}
}

//新用户加入
func Join(user string, roomid int64, ws *websocket.Conn) {
	subscribe <- Subscriber{Name: user, Roomid: roomid, Conn: ws}
}
//用户离开
func Leave(user string, roomid int64) {
	Room[roomid] = Room[roomid] - 1
	unsubscribe <- user
	models.Delete_chatroom(user, roomid)
}

func init() {
	go chatroom()
}
//检查用户是否已经存在
func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}
//检查是否有时间发生
func chatroom() {
	for {
		select {
		case sub := <-subscribe://有新人加入聊天室
			if !isUserExist(subscribers, sub.Name) {
				subscribers.PushBack(sub)
				publish <- newEvent(models.EVENT_JOIN, sub.Name, sub.Roomid, "")//加入对应事件
			}
			models.Join_chatroom(sub.Name, sub.Roomid)
		case event := <-publish://读取到新的事件
			broadcastWebSocket(event) //websocket处理事件发送给前端
			if event.Type == models.EVENT_MESSAGE {
				models.Join_recode(event.Roomid, event.User, event.Content)//储存聊天记录mysql
			}
		case unsub := <-unsubscribe://用户离开
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					rid := sub.Value.(Subscriber).Roomid
					subscribers.Remove(sub)
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("websocket closed", unsub)
					}
					publish <- newEvent(models.EVENT_LEAVE, unsub, rid, "")
					break
				}
			}
		}
	}
}
