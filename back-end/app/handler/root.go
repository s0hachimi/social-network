package handler

import (
	"fmt"
	"net/http"

	"social-network/app/utils"
)

func Root(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("SessionToken")
	if err != nil || session.Value == "" {
		fmt.Println(err)
		utils.SendData(w, http.StatusForbidden, map[string]any{"status": false})
		return
	}
	fmt.Println(session.Value)
}
