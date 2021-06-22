package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8080", "http Server address")

var upgrader = websocket.Upgrader{}

func echo(resp http.ResponseWriter, req *http.Request) {
	c, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Print("uptrade:", err)
		return
	}
	defer c.Close()

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Print("Read from Client err :", err)
			break
		}
		log.Printf("recv: %s\n", message)
		c.WriteMessage(messageType, []byte("server echo :"+string(message)))
	}

}
func home(resp http.ResponseWriter, req *http.Request) {
	homeTemplate.Execute(resp, "ws://"+req.Host+"/echo")
}

func main() {

	flag.Parse()
	log.SetFlags(0)
	// 设置路由
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	// 启动 http 服务器
	log.Println(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
