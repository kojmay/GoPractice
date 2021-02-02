package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kojmay/GoPractice/1.go-conccurency/3.chatroom_vue/global"
	"github.com/kojmay/GoPractice/1.go-conccurency/3.chatroom_vue/server"
)

var (
	addr   = ":2022"
	banner = `
    ____                  _____
   |      |    |    /\      |
   |      |____|   /  \     | 
   |      |    |  /----\    |
   |____  |    | /      \   |
	ChatRoom，start on：%s
`
)

func init() {
	global.Init()
}

func main() {
	fmt.Printf(banner, addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
