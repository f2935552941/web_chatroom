$(document).ready(function() {
    var $window=$(window)
    var $messageArea=$('#messageArea');
    var $inputArea=$('#inputArea');
    $inputArea.focus();
    var connected=false;
    function getRandomInt(max) {
        return Math.floor(Math.random() * Math.floor(max));
    }
    function getColor(username) {
        var COLORS = [
            '#e21400', '#91580f', '#f8a700', '#f78b00',
            '#58dc00', '#287b00', '#a8f07a', '#4ae8c4',
            '#3b88eb', '#3824aa', '#a700ff', '#d300e7'
        ];
        index=getRandomInt(10);
        return COLORS[index];
    }
    var nameColor=getColor($("#uname").text());
    $("#uname").css('color',nameColor);

    socket=new WebSocket('ws://'+window.location.host+'/ws/join?uname='+$('#uname').text()+'&roomid='+$('#roomid').text());
    
    socket.onopen=function() {
        console.log("webSocket open");
        connected=true;
    }
    socket.onclose=function() {
        console.log("WebSocket close");
        connected=false;
    }
    socket.onmessage=function(event) {
        console.log(event.data)
        var data=JSON.parse(event.data)
        console.log("revice:",data);
        var name=data.User;
        var type=data.Type;
        var roomid=data.Roomid;
        var msg=":"+data.Content;
        console.log(roomid);
        console.log($('#roomid').text());
        if (roomid!=$('#roomid').text()) return ;
        console.log("sucessful")
        var $messageDiv;
        switch(data.Type) {
            case 0://join
                var $messageBodyDiv = $('<span style="color:#999999;">').text(name+" has joined the chatroom");
                $messageDiv = $('<li style="list-style-type:none;font-size:15px;text-align:center;"/>').append($messageBodyDiv);
                //$messageDiv.innerText=mess;
                break;
            case 1://leave
                var $messageBodyDiv = $('<span style="color:#999999;">').text(name+" has left the chatroom");
                $messageDiv = $('<li style="list-style-type:none;font-size:15px;text-align:center;"/>').append($messageBodyDiv);
                break;
            case 2://send message
                var $usernameDiv = $('<span style="margin-right: 15px;font-weight: 700;overflow: hidden;text-align: right;"/>').text(name);
                if(name==$('#uname').text()) {
                    $usernameDiv.css('color',nameColor);
                } else {
                    $usernameDiv.css('color',getColor(name));
                }
                var $messageBodyDiv = $('<span style="color: gray;"/>')
                    .text(msg);
                $messageDiv = $('<li style="list-style-type:none;font-size:25px;"/>')
                    .data('username', name)
                    .append($usernameDiv, $messageBodyDiv);
        }
        $messageArea.append($messageDiv);
        $messageArea[0].scrollTop = $messageArea[0].scrollHeight;   // 让屏幕滚动
    };
    $("#sendBtn").click(function() {
        sendMessage();
    });
    function sendMessage() {
        var inputMessage=$inputArea.val();
        if(inputMessage&&connected) {
            $inputArea.val('');
            socket.send(inputMessage);
            console.log("send message:"+inputMessage);
        }
    }
});