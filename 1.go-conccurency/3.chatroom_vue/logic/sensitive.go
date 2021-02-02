package logic

import (
	"strings"

	"github.com/kojmay/GoPractice/1.go-conccurency/3.chatroom_vue/global"
)

// FilterSensitive replace sensitive words with **
func FilterSensitive(content string) string {
	for _, word := range global.SensitiveWords {
		content = strings.ReplaceAll(content, word, "**")
	}
	return content
}
