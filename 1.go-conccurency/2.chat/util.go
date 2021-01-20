package chat

import (
	"fmt"
)

func jugeErr(err error, prompt string) bool {
	if err != nil {
		fmt.Println(prompt, ", err info: ", err)
		return false
	}
	return true
}
