package config

import (
	"fmt"
	"log"
	"os"

	"music-app/internal/auth"
	"music-app/internal/playlist"
	"music-app/internal/song"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnStr() string {
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "postgres"
	}
	dbname := os.Getenv("MUSIC_DB")
	if dbname == "" {
		dbname = "musicdb" 
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "razeem19" 
	}
	sslmode := os.Getenv("SSL_MODE")
	if sslmode == "" {
		sslmode = "disable" 
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1" 
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432" 
	}

	return fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s", host, user, dbname, password, port, sslmode)
}

func ConnectDatabase() (*gorm.DB, error) {
	dsn := GetDBConnStr()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&auth.User{}, &song.Song{}, &playlist.Playlist{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return nil, err
	}

	log.Println("Database connected and migrated successfully.")
	return db, nil
}
