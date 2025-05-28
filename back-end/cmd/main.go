package main

import (
	"fmt"
	"net/http"
	"social-network/app/handler"
	"social-network/app/utils"
	db "social-network/database"
)

func main() {
	Db, err := db.Db()
	if err != nil {
		fmt.Println("====Z", err)
		return
	}

	defer Db.Close()

	router := http.NewServeMux()

	router.Handle("/", Middleware(http.HandlerFunc(handler.Root)))

	fmt.Println("✅ Server running on: http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session")
		if err != nil || session.Value == "" {
			utils.SendData(w, http.StatusForbidden, map[string]any{"status": false})
			return
		}

		have := db.HaveToken(session.Value)
		if !have {
			utils.SendData(w, http.StatusForbidden, map[string]any{"status": false})
			return
		}

		next.ServeHTTP(w, r)
	})
}
