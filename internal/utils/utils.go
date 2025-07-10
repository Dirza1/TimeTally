package utils

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DatabaseConnection() database.Queries {
	godotenv.Load("/home/jasperolthof/workspace/projects/Time-and-expence-registration/.env")
	dbURL := os.Getenv("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Datbase error:", err)
	}
	queries := database.New(dbConn)
	return *queries
}

func TimeParse(toParseDate string) time.Time {
	layout := "02-01-2006"
	Date, err := time.Parse(layout, toParseDate)
	if err != nil {
		log.Fatal("error during parsing of dates: ", err)
	}
	return Date
}
