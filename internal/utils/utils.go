package utils

import (
	"database/sql"
	"log"
	"os"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DatabaseConnection() database.Queries {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env:", err)
	}
	dbURL := os.Getenv("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Datbase error:", err)
	}
	queries := database.New(dbConn)
	return *queries
}
