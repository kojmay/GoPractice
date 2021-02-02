package server

import (
	"net/http"

	"github.com/kojmay/GoPractice/1.go-conccurency/3.chatroom_vue/logic"
)

func RegisterHandle() {
	//handle broadcast message
	go logic.Broadcaster.Start()

	http.HandleFunc("/", homeHandleFunc)
	http.HandleFunc("/user_list", userListHandleFunc)
	http.HandleFunc("/ws", WebSocketHandleFunc)

}
