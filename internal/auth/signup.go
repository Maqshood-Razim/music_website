package auth

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func SignupHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.Create(&user).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		response := map[string]string{
			"message": "Signup successful", "username": user.Username,
		}
		json.NewEncoder(w).Encode(response)

	}
}
