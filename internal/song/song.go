package song

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Title  string `json:"title"`
	Artist string `json:"artist"`
	URL    string `json:"url"`
}


func SongsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
			return
		}

		
		var songs []Song
		if err := db.Find(&songs).Error; err != nil {
			http.Error(w, "Failed to fetch songs: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		encoded, err := json.MarshalIndent(songs, "", "    ")
		if err != nil {
			http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
	}
}

func AddSongs(db *gorm.DB) {

	songs := []Song{
		{Title: "Shape of You", Artist: "Ed Sheeran", URL: "https://www.youtube.com/watch?v=JGwWNGJdvx8"},
		{Title: "Blinding Lights", Artist: "The Weeknd", URL: "https://www.youtube.com/watch?v=fHI8X4OXluQ"},
		{Title: "Levitating", Artist: "Dua Lipa", URL: "https://www.youtube.com/watch?v=TUVcZfQe-Kw"},
		{Title: "Someone Like You", Artist: "Adele", URL: "https://www.youtube.com/watch?v=hLQl3WQQoQ0"},
		{Title: "Bad Guy", Artist: "Billie Eilish", URL: "https://www.youtube.com/watch?v=DyDfgMOUjCI"},
	}


	for _, song := range songs {
		var existing Song
		
		if err := db.Where("title = ?", song.Title).First(&existing).Error; err == nil {
			log.Printf("Song already exists: %s by %s", song.Title, song.Artist)
			continue
		}

		if err := db.Create(&song).Error; err != nil {
			log.Printf("Failed to add song: %s - Error: %v", song.Title, err)
		} else {
			log.Printf("Successfully added song: %s by %s", song.Title, song.Artist)
		}
	}

	log.Println("Sample songs addition completed.")
}
