package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"social-network/app/utils"
	db "social-network/database"
)

type UserLog struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	CORS(w, r)

	if r.Method == "POST" {

		var info UserLog
		errore := json.NewDecoder(r.Body).Decode(&info)
		if errore != nil {
			fmt.Println("hona", errore)
			return
		}

		Access(w)

		fmt.Println(info.Email, info.Password)

		var boo bool
		var err error
		typ := ""
		var hashedPassword string

		if strings.Contains(info.Email, "@") {
			boo = db.CheckInfo(info.Email, "email")
			typ = "email"
		} else {
			boo = db.CheckInfo(info.Email, "nikname")
			typ = "nikname"
		}

		if !boo {
			hashedPassword, err = db.Getpasswor(typ, info.Email)
		}

		if boo || err != nil || !utils.ComparePassAndHashedPass(hashedPassword, info.Password) {
			utils.SendData(w, http.StatusUnauthorized, map[string]any{
				"error":  "Invalid " + typ + " or password",
				"status": false,
			})
			return
		}

		SessionToken, expiration := utils.GenerateSessionToken()

		err = db.Updatesession(typ, SessionToken, info.Email)
		if err != nil {
			utils.SendData(w, http.StatusUnauthorized, map[string]any{
				"error":  "Invalid " + typ + " or password",
				"status": false,
			})
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "SessionToken",
			Value:    SessionToken,
			Expires:  expiration,
			// Path:     "/",
		})

		utils.SendData(w, http.StatusOK, map[string]any{
			"error":  "Login successful",
			"status": true,
		})

	}
}
