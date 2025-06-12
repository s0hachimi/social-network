package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"social-network/app/utils"
	db "social-network/database"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	Nickname  string `json:"nickname"`
	Profile   string `json:"profile_image"`
	BirthDate string `json:"date"`
	AboutMe   string `json:"about_me"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	CORS(w, r)

	if r.Method == http.MethodPost {

		Access(w)

		var err error
		message := ""

		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		uploadDir := "uploads"
		os.MkdirAll(uploadDir, os.ModePerm)

		filePath := ""

		file, handler, err := r.FormFile("profile_image") //  && err.Error() != "http: no such file"
		if err == nil && handler != nil {

			defer file.Close()

			uniqueFilename := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)

			filePath = filepath.Join(uploadDir, uniqueFilename)
			dst, err := os.Create(filePath)
			if err != nil {
				http.Error(w, "Unable to save the file", http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			_, err = io.Copy(dst, file)
			if err != nil {
				http.Error(w, "Unable to save the file", http.StatusInternalServerError)
				return
			}
			
		}

		jsonData := r.FormValue("info")

		var info User
		errore := json.Unmarshal([]byte(jsonData), &info)
		if errore != nil {
			fmt.Println("hona", errore)
			return
		}

		validatEmail := db.CheckInfo(info.Email, "email")
		if !validatEmail {
			message = "Email already exists"
		}

		validatNikname := db.CheckInfo(info.Nickname, "nickname")
		if !validatNikname {
			message = "Nickname already exists"
		}

		if !validatEmail && !validatNikname {
			message = "Email and nickname already exist"
		}

		if message != "" {
			utils.SendData(w, http.StatusConflict, map[string]any{"success": false, "message": message})
			return
		}

		info.Password, err = utils.HashPassword(info.Password)
		if err != nil {
			utils.SendData(w, http.StatusInternalServerError, map[string]any{"success": false, "message": "Internal server error"})
			return
		}
		err = db.Insertuser(info.FirstName, info.LastName, info.Email, info.Gender, info.Age, info.Nickname, info.Password, filePath, info.BirthDate, info.AboutMe)
		if err != nil {
			utils.SendData(w, http.StatusInternalServerError, map[string]any{"success": false, "message": "Internal server error"})
			return
		}

		utils.SendData(w, http.StatusOK, map[string]any{"success": true, "message": ""})
	}
}

func CORS(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusOK)
		return
	}
}

func Access(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
