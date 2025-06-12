package handler

import (
	"fmt"
	"net/http"

	"social-network/app/utils"
	db "social-network/database"
)

func Root(w http.ResponseWriter, r *http.Request) {
	CORS(w, r)

	if r.Method == http.MethodGet {

		Access(w)

		session, err := r.Cookie("SessionToken")
		if err != nil || session.Value == "" {
			fmt.Println(err)
			utils.SendData(w, http.StatusForbidden, map[string]any{"status": false})
			return
		}

		have := db.HaveToken(session.Value)
		if !have {
			utils.SendData(w, http.StatusForbidden, map[string]any{"status": false})
			return
		}

		utils.SendData(w, http.StatusOK, map[string]any{"cookie": session.Value, "status": true})
	}
}
