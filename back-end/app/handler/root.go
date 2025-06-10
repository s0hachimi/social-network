package handler

import (
	"fmt"
	"net/http"

	"social-network/app/utils"
)

func Root(w http.ResponseWriter, r *http.Request) {
	CORS(w, r)

	

	if r.Method == http.MethodGet {
		Access(w)
		
		fmt.Println(r.Cookies())

		session, err := r.Cookie("SessionToken")
		if err != nil || session.Value == "" {
			fmt.Println(err)
			utils.SendData(w, http.StatusForbidden, map[string]any{"status": false})
			return
		}

		utils.SendData(w, http.StatusOK, map[string]any{"cookie": session.Value, "status": true})
	}
}
