package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Load file .env gagal")
	}

	log.Println("Load file .env berhasil")
}
