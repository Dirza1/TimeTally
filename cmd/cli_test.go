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

type CLITest struct {
	Name       string
	Command    string
	Args       []string
	WantOutput string
	WantErr    bool
}

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
	tests := []CLITest{
		{
			Name:       "Test reset of both databases",
			Command:    "Time-and-expence-registration",
			Args:       []string{"reset", "-c", "true", "-t", "All", "-p", "Odin2203!"},
			WantOutput: "reset called on all",
			WantErr:    false,
		},
		{
			Name:       "Test reset of Finance database",
			Command:    "Time-and-expence-registration",
			Args:       []string{"reset", "-c", "true", "-t", "Finance", "-p", "Odin2203!"},
			WantOutput: "reset called on all",
			WantErr:    false,
		},
		{
			Name:       "Test reset of Time database",
			Command:    "Time-and-expence-registration",
			Args:       []string{"reset", "-c", "true", "-t", "Time", "-p", "Odin2203!"},
			WantOutput: "reset called on all",
			WantErr:    false,
		},
	}
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
