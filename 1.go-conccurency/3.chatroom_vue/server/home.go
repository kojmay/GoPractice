package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/kojmay/GoPractice/1.go-conccurency/3.chatroom_vue/global"
	"github.com/kojmay/GoPractice/1.go-conccurency/3.chatroom_vue/logic"
)

// homeHandleFunc : parse template file
func homeHandleFunc(w http.ResponseWriter, req *http.Request) {
	templ, err := template.ParseFiles(global.RootDir + "/template/home.html")
	if err != nil {
		fmt.Fprint(w, "Parse template file err!")
		return
	}

	err = templ.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, "Execute template error!")
		return
	}
}

//userListHandleFunc : return all users
func userListHandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	userList := logic.Broadcaster.GetUserList()
	b, err := json.Marshal(userList)

	if err != nil {
		fmt.Fprint(w, `[]`)
	} else {
		fmt.Fprint(w, string(b))
	}

}
