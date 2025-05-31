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

	router.HandleFunc("/", handler.Root)

	router.HandleFunc("/register", handler.Register)
	router.HandleFunc("/login", handler.Login)

	fmt.Println("✅ Server running on: http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        // w.Header().Set("Content-Type", "application/json")
        // w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        // w.Header().Set("Access-Control-Allow-Credentials", "true")

        // if r.Method == "OPTIONS" {
        //     w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        //     w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        //     w.WriteHeader(http.StatusOK)
        //     return
        // }

		session, err := r.Cookie("SessionToken")
		if err != nil || session.Value == "" {
			fmt.Println(err)
			utils.SendData(w, http.StatusForbidden, map[string]any{"status": false})
			return
		}
		fmt.Println(session.Value)

		have := db.HaveToken(session.Value)
		if !have {
			utils.SendData(w, http.StatusForbidden, map[string]any{"status": false})
			return
		}

		next.ServeHTTP(w, r)
	})
}
