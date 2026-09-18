package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 升级器：http -> websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许跨域，开发环境打开
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 处理websocket连接
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 将http连接升级为websocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade err:", err)
		return
	}
	defer conn.Close()

	log.Println("客户端已连接")

	for {
		// 读取客户端消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read err:", err)
			break
		}
		log.Printf("收到客户端: %s\n", msg)

		// 回显消息给客户端
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("write err:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("服务启动 :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
