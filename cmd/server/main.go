package main

import (
	"log"
	"music-app/internal/auth"
	"music-app/internal/playlist"
	"music-app/internal/song"
	"music-app/pkg/config"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)

func main() {
	
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found. Using default environment variables.")
	}


	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	
	song.AddSongs(db)


	http.HandleFunc("/signup", auth.SignupHandler(db))
	http.HandleFunc("/login", auth.LoginHandler(db))
	http.HandleFunc("/songs", song.SongsHandler(db))
	http.HandleFunc("/playlists", playlist.PlaylistsHandler(db))


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
