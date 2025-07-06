package utils

import (
	"database/sql"
	"log"
	"os"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
)

func DatabaseConnection() database.Queries {
	dbURL := os.Getenv("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Datbase error:", err)
	}
	queries := database.New(dbConn)
	return *queries
}
