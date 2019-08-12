$(function(){

    socket=new WebSocket('ws://'+window.location.host+'/index/true');
    
    socket.onopen=function() {
        console.log("webSocket open");
        connected=true;
    }
    socket.onclose=function() {
        console.log("WebSocket close");
        connected=false;
    }
    socket.onmessage=function(event) {
        var data=JSON.parse(event.data)
        console.log("revice:",data);
        var roomid=data.Roomid;
        var num=":"+data.PerNum;
        add(roomid,num)
    };
    var add=function(rid,pnum) {
        var a=$("<div class='c9'></div>").text("");
        a.attr("roomid",""+rid);
        var rrid=$("<a></a>").text("roomid:"+rid);
        var imgg=$("<img src='static/img/3.png'>");
        var c=$("<a></a>").text("person_num:"+pnum);
        rrid.click(function(){
			var id=$(this).parent().attr("roomid");
			Join(id);
		});
        imgg.click(function(){
			var id=$(this).parent().attr("roomid");
			Join(id);
		});
		c.click(function(){
			var id=$(this).parent().attr("roomid");
			Join(id);
		});
		a.append(rrid,imgg,c);
		$("#mp").append(a);
    }
    var Join=function(rid) {
        $('#roomid').val(rid);
        $('#btn').click();
    }
});