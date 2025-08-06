package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type LoggedInUser struct {
	Name                    string
	Access_Finance          bool
	Access_Timeregistration bool
	Administrator           bool
}

func DatabaseConnection() database.Queries {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading enviromental variables")
		return database.Queries{}
	}

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

func ReturnLoggedInUser() LoggedInUser {
	user := LoggedInUser{}
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading enviromental variables")
		return user
	}
	user.Name = os.Getenv("Name_logged_in")
	user.Access_Finance, _ = strconv.ParseBool(os.Getenv("Access_finance"))
	user.Access_Timeregistration, _ = strconv.ParseBool(os.Getenv("Access_timeregistration"))
	user.Administrator, _ = strconv.ParseBool(os.Getenv("Administrator"))

	return user
}
