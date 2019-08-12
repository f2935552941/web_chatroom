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

package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type EventType int

const (
	EVENT_JOIN    = 0
	EVENT_LEAVE   = 1
	EVENT_MESSAGE = 2
)

//event结构体
type Event struct {
	Type      EventType // 分别表示3种事件 JOIN, LEAVE, MESSAGE
	User      string    //用户昵称
	Roomid    int64     //用户加入或创建的房间id
	Timestamp int       // 发送消息时的当前时间 Unix timestamp (secs) （未使用）
	Content   string    //发送的消息文本
}

//以下 sql模块
func check(err error) {
	if err != nil {
		panic(err)
	}
}

//用户加入或创建房间
func Join_chatroom(uname string, roomid int64) {
	db, err := sql.Open("mysql", "root:makise@/test?charset=utf8")
	check(err)
	stmt, err := db.Prepare("INSERT INTO user_room SET username=?,roomid=?")
	check(err)
	res, err := stmt.Exec(uname, roomid)
	check(err)
	fmt.Println(res)
	db.Close()
}

//用户退出房间
func Delete_chatroom(uname string, roomid int64) {
	db, err := sql.Open("mysql", "root:makise@/test?charset=utf8")
	check(err)
	stmt, err := db.Prepare("delete from user_room where username=? and roomid=?")
	check(err)
	res, err := stmt.Exec(uname, roomid)
	check(err)
	fmt.Println(res)
	db.Close()
}

//聊天记录
func Join_recode(rid int64, uname string, content string) {
	db, err := sql.Open("mysql", "root:makise@/test?charset=utf8")
	check(err)
	stmt, err := db.Prepare("INSERT INTO record SET roomid=?,username=?,content=?")
	check(err)
	res, err := stmt.Exec(rid, uname, content)
	check(err)
	fmt.Println(res)
	db.Close()
}
