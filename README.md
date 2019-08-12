# web_chatroom
简易聊天室，支持多人创建房间聊天  
进入src目录  
go build main.go  
./main  
浏览器输入localhost:8080  

go语言学习部分来自https://github.com/astaxie/build-web-application-with-golang  
此项目主要以beego框架为基础  
采用MVC模式开发 数据传输主要使用websocket技术  
输入用户名进入主页面后输入房间ID若房间已存在自动进入房间否则创建房间  
主页面显示当前所有存在房间  
同时支持单机已存在的房间直接进入聊天室  
用户房间信息和聊天记录保存在mysql中  
介于前端水平过低很多功能暂时无法实现，例如查询聊天记录 房间加入密码设置等。
