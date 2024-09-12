package main

import (
	"database/sql"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"url-shortener/config"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/repository/postgres"
	"url-shortener/internal/service"
)

func main() {
	useDB := flag.Bool("d", false, "turn on postgres database")
	flag.Parse()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error init env variables: ", err)
	}
	cfg := config.InitConfig()

	var db *sql.DB
	if *useDB {
		var err error
		db, err = postgres.NewPostgresDB(cfg)
		if err != nil {
			log.Fatal("failed to init postgres: ", err)
		}
		log.Print("postgres is working")
	}

	repos, err := repository.NewRepository(*useDB, db)
	if err != nil {
		log.Fatal("failed to init repository", err)
	}
	srv := service.NewService(repos)
	h := handler.NewHandler(srv)

	router := gin.Default()
	h.InitRoutes(router)

	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(router.Run(":" + cfg.Port))
}
