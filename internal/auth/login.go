package auth

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func LoginHandler(db *gorm.DB) http.HandlerFunc {
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

		var storedUser User
		if err := db.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		if user.Password != storedUser.Password {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := map[string]string{
			"message": "login successfull", 
            "welcome": user.Username,
		}
		json.NewEncoder(w).Encode(response)

	}
}
