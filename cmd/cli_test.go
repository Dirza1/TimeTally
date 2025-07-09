package cmd

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
)

var querry database.Queries

func TestMain(m *testing.M) {
	fmt.Println("Starting tests and database connection")
	querry = utils.DatabaseConnection()

	exitCode := m.Run()

	fmt.Println("Ending tests and resetting database")
	_ = querry.ResetTransaction(context.Background())
	_ = querry.ResetTimeRegistration(context.Background())

	os.Exit(exitCode)

}

func TestReset(t *testing.T) {
	os.Args = []string{"Time-and-expence-registration", "reset", "-c", "true", "-t", "All", "-p", "Odin2203!"}
	err := rootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	timeList, err := querry.OverviewAllTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(timeList) != 0 {
		fmt.Println("Test failed, time database not empty")
		t.Fail()
	}
	financeList, err := querry.OverviewAllTransactions(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(financeList) != 0 {
		fmt.Println("Test failed, finance database not empty")
		t.Fail()
	}
}
