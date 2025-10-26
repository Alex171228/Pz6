package main

import (
	"example.com/pz6-gorm/internal/db"
	"example.com/pz6-gorm/internal/http"
	"example.com/pz6-gorm/internal/models"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env файл не найден — используется системное окружение")
	}

	d := db.Connect()

	// Авто-миграция схемы
	if err := d.AutoMigrate(&models.User{}, &models.Note{}, &models.Tag{}); err != nil {
		log.Fatal("migrate:", err)
	}

	r := httpapi.BuildRouter(d)

	log.Println("Сервер слушает на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
