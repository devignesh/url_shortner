package main

import (
	"log"

	"url_shortener/internal/handler"
	models "url_shortener/internal/model"
	"url_shortener/internal/repository"
	"url_shortener/internal/service"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	//db connection
	db := Dbcon()

	defer db.Close()

	// Run database migrations
	db.AutoMigrate(&models.URL{})

	// Create repository and service instances
	urlRepo := repository.NewURLShortenerRepository(db)
	urlService := service.NewURLShortenerService(urlRepo)

	log.Println("\n urlrepo", urlRepo)
	log.Println("\n urlservice", urlService)

	// Create Gin router
	// router := gin.Default()

	router := handler.SetupRouter(urlService)
	log.Println("\n router", router)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func Dbcon() *gorm.DB {
	//db connection
	dsn := "host=localhost port=5432 user=postgres dbname=url_shortners sslmode=disable password=postgres"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("db connected successfully")

	return db
}
