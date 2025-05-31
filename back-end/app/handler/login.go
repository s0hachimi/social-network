package handler

import (
	"net/http"
	"strings"

	"social-network/app/utils"
	db "social-network/database"
)

func Login(w http.ResponseWriter, r *http.Request) {
	CORS(w, r)

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		var boo bool
		var err error
		typ := ""
		var hashedPassword string
		if strings.Contains(email, "@") {
			boo = db.CheckInfo(email, "email")
			typ = "email"
		} else {
			boo = db.CheckInfo(email, "nikname")
			typ = "nikname"
		}

		if !boo {
			hashedPassword, err = db.Getpasswor(typ, email)
		}

		if boo || err != nil || !utils.ComparePassAndHashedPass(hashedPassword, password) {
			utils.SendData(w, http.StatusUnauthorized, map[string]any{
				"error":  "Invalid " + typ + " or password",
				"status": false,
			})
			return
		}
		SessionToken, erre := utils.GenerateSessionToken()
		if erre != nil {
			utils.SendData(w, http.StatusUnauthorized, map[string]any{
				"error":  "Invalid " + typ + " or password",
				"status": false,
			})
			return
		}
		err = db.Updatesession(typ, SessionToken, email) ////email mmkin ikon nikname mmkin ikon email
		if err != nil {
			utils.SendData(w, http.StatusUnauthorized, map[string]any{
				"error":  "Invalid " + typ + " or password",
				"status": false,
			})
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "SessionToken",
			Value: SessionToken,
			Path:  "/",
		})

		utils.SendData(w, http.StatusOK, map[string]any{
			"error":  "Login successful",
			"status": true,
		})

	}
}
