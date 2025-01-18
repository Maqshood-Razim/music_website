package playlist

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Name    string `json:"name"`
	SongIDs []int  `json:"song_ids" gorm:"type:int[]"`
}

func PlaylistsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodPost {
			http.Error(w, "Only GET and POST methods are allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.Method == http.MethodPost {
			var playlist Playlist
			if err := json.NewDecoder(r.Body).Decode(&playlist); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := db.Create(&playlist).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
		} else {
			var playlists []Playlist
			if err := db.Find(&playlists).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := json.NewEncoder(w).Encode(playlists); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
