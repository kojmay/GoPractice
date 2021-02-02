package logic

import (
	"expvar"
	"fmt"
	"log"

	"github.com/kojmay/GoPractice/1.go-conccurency/3.chatroom_vue/global"
)

func init() {
	expvar.Publish("message_queue", expvar.Func(calcMessageQueueLen))
}

func calcMessageQueueLen() interface{} {
	fmt.Println("===len=:", len(Broadcaster.messageChannel))
	return len(Broadcaster.messageChannel)
}

// broadcaster
type broadcaster struct {
	// 所有在线用户
	users map[string]*User

	// all channels
	enteringChannel chan *User
	leavingChannel  chan *User
	messageChannel  chan *Message
	// 判断用户是否能进入聊天室
	checkUserChannel         chan string
	checkUserCanEnterChannel chan bool

	// 获取用户列表
	requestUserChannel chan struct{}
	usersChannel       chan []*User
}

// Broadcaster 广播
var Broadcaster = &broadcaster{
	users: make(map[string]*User),

	enteringChannel: make(chan *User),
	leavingChannel:  make(chan *User),
	messageChannel:  make(chan *Message, global.MessageQueueLen),

	checkUserChannel:         make(chan string),
	checkUserCanEnterChannel: make(chan bool),

	requestUserChannel: make(chan struct{}),
	usersChannel:       make(chan []*User),
}

// Start 启动广播器
func (b *broadcaster) Start() {

	for {
		select {
		case user := <-b.enteringChannel:
			b.users[user.NickName] = user

		case user := <-b.leavingChannel:
			delete(b.users, user.NickName)

			user.CloseMessageChannel()
		case msg := <-b.messageChannel:
			for _, user := range b.users {
				if user.UID != msg.User.UID {
					user.MessageChannel <- msg
				}
			}
		case nickname := <-b.checkUserChannel:
			if _, ok := b.users[nickname]; ok {
				b.checkUserCanEnterChannel <- false
			} else {
				b.checkUserCanEnterChannel <- true
			}
		case <-b.requestUserChannel:
			userList := make([]*User, 0, len(b.users))
			for _, user := range b.users {
				userList = append(userList, user)
			}
			b.usersChannel <- userList
		}
	}
}

// UserEntering capture entering users
func (b *broadcaster) UserEntering(u *User) {
	b.enteringChannel <- u
}

// UserLeaving capture leaving users
func (b *broadcaster) UserLeaving(u *User) {
	b.leavingChannel <- u
}

// Broadcast : broadcast msg to all users
func (b *broadcaster) Broadcast(msg *Message) {
	if len(b.messageChannel) >= global.MessageQueueLen {
		log.Println("Attention: Broadcast queue FULL!")
	}
	b.messageChannel <- msg
}

// CanEnterRoom 判断用户的nickname是否有重复
func (b *broadcaster) CanEnterRoom(nickname string) bool {
	b.checkUserChannel <- nickname

	return <-b.checkUserCanEnterChannel
}

func (b *broadcaster) GetUserList() []*User {
	b.requestUserChannel <- struct{}{}
	return <-b.usersChannel
}
