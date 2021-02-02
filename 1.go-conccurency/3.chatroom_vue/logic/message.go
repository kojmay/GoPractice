package logic

import (
	"time"

	"github.com/spf13/cast"
)

const (
	MsgTypeNormal = iota
	MsgTypeWelcome
	MsgTypeUserEnter
	MsgTypeUserLeave
	MsgTypeError
)

// Message 消息结构体
type Message struct {
	User           *User     `json:"user"`
	Type           int       `json:"type"`
	Content        string    `json:"content"`
	MsgTime        time.Time `json:"msg_time"`
	ClientSendTime time.Time `json:"client_send_time"`

	// @ function
	Ats []string `json:"ats"`
}

//NewMessage gen a new message
func NewMessage(user *User, content, clientTime string) *Message {
	message := &Message{
		User:    user,
		Type:    MsgTypeNormal,
		Content: content,
		MsgTime: time.Now(),
	}
	if clientTime != "" {
		message.ClientSendTime = time.Unix(0, cast.ToInt64(clientTime))
	}
	return message
}

//NewWelcomeMessage welcome
func NewWelcomeMessage(user *User) *Message {
	return &Message{
		User:    user,
		Type:    MsgTypeWelcome,
		Content: user.NickName + " 您好，欢迎加入聊天室！",
		MsgTime: time.Now(),
	}
}

//NewUserEnterMessage enter
func NewUserEnterMessage(user *User) *Message {
	return &Message{
		User:    user,
		Type:    MsgTypeUserEnter,
		Content: user.NickName + " 加入了聊天室",
		MsgTime: time.Now(),
	}
}

//NewUserLeaveMessage leave
func NewUserLeaveMessage(user *User) *Message {
	return &Message{
		User:    user,
		Type:    MsgTypeUserLeave,
		Content: user.NickName + " 离开了聊天室",
		MsgTime: time.Now(),
	}
}

// NewErrorMessage error
func NewErrorMessage(content string) *Message {

	return &Message{
		User:    System,
		Type:    MsgTypeError,
		Content: content,
		MsgTime: time.Now(),
	}
}
