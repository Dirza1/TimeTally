package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/Dirza1/TimeTally/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Session struct {
	UserName string    `json:"user_name"`
	UserID   uuid.UUID `json:"user_id"`
	LastUsed time.Time `json:"last_used"`
}

type Backup struct {
	WeeklyBackUp    time.Time `json:"weekly_back_up"`
	QuarterlyBackUp time.Time `json:"quarterly_back_up"`
}

func DatabaseConnection() database.Queries {
	err := godotenv.Load(".env")
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

func Database() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading enviromental variables")
		return nil
	}

	dbURL := os.Getenv("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Datbase error:", err)
	}
	return dbConn
}

func TimeParse(toParseDate string) time.Time {
	layout := "02-01-2006"
	Date, err := time.Parse(layout, toParseDate)
	if err != nil {
		log.Fatal("error during parsing of dates: ", err)
	}
	return Date
}

func Hashpassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CompairPaswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SaveSession(s *Session) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(".session.json", data, 0600)
}

func LoadSession() (*Session, error) {
	data, err := os.ReadFile(".session.json")
	if err != nil {
		return nil, err
	}
	var s Session
	err = json.Unmarshal(data, &s)
	return &s, err
}

func UpdateSession() {
	session, err := LoadSession()
	if err != nil {
		fmt.Printf("\nError loading current session. Err:\n%s\n", err)
		return
	}
	newSession := Session{
		UserName: session.UserName,
		UserID:   session.UserID,
		LastUsed: time.Now(),
	}
	err = SaveSession(&newSession)
	if err != nil {
		fmt.Printf("\nError saving new session. Err:\n%s\n", err)
		return
	}
}

func LoadBackupTimes() (*Backup, error) {
	data, err := os.ReadFile(".backups.json")
	if err != nil {
		return nil, err
	}
	var b Backup
	err = json.Unmarshal(data, &b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func BackupProcess() {
	CurrentTime := time.Now()
	BackupTimes, err := LoadBackupTimes()
	if err != nil {
		fmt.Println("There was an error loading your latest backup times. To ensure the backup process runs as expected are made new backups will now be made")
		BackupTimes = &Backup{
			WeeklyBackUp:    CurrentTime.AddDate(0, 0, -14),
			QuarterlyBackUp: CurrentTime.AddDate(0, -3, 0),
		}
	}
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading enviromental variables")
		return
	}
	dbURL := os.Getenv("DB_URL")

	if CurrentTime.Sub(BackupTimes.WeeklyBackUp) > 7*24*time.Hour {

		backupFile := fmt.Sprintf("backups/weekly/db_%s.sql", time.Now().Format("2006-01-02_150405"))
		cmd := exec.Command("pg_dump", dbURL, "-F", "p", "-f", backupFile)
		err = cmd.Run()
		if err != nil {
			fmt.Println("Backup failed:", err)
			return
		} else {
			fmt.Println("Backup saved to", backupFile)
		}
		BackupTimes.WeeklyBackUp = time.Now()
	}
	if CurrentTime.Sub(BackupTimes.QuarterlyBackUp) > 91*24*time.Hour {

		tables := []string{"users", "timeregistration", "Finances"}
		for _, table := range tables {
			csvFile := fmt.Sprintf("backups/Quarterly/Boekhouding_%s_%s.csv", table, time.Now().Format("2006-01-02"))
			cmd := exec.Command("psql", dbURL, "-c", fmt.Sprintf("\\copy Boekhouding.%s TO '%s' CSV HEADER", table, csvFile))
			err := cmd.Run()
			if err != nil {
				fmt.Println("CSV backup failed for table", table, ":", err)
				return
			} else {
				fmt.Println("CSV backup created:", csvFile)
			}

		}
		BackupTimes.QuarterlyBackUp = time.Now()

	}
	SaveNewBackupTimes(BackupTimes)
}

func SaveNewBackupTimes(b *Backup) error {
	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(".backups.json", data, 0600)
}
