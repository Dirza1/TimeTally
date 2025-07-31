package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
)

var querry database.Queries

type CLITest struct {
	Name       string
	Args       []string
	WantOutput string
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
			Args:       []string{"Time-and-expence-registration", "reset", "-c", "true", "-t", "All", "-p", "Odin2203!"},
			WantOutput: "reset called on all\n",
		},
		{
			Name:       "Test reset of Finance database",
			Args:       []string{"Time-and-expence-registration", "reset", "-c", "true", "-t", "Finance", "-p", "Odin2203!"},
			WantOutput: "reset called on finance\n",
		},
		{
			Name:       "Test reset of Time database",
			Args:       []string{"Time-and-expence-registration", "reset", "-c", "true", "-t", "Time", "-p", "Odin2203!"},
			WantOutput: "reset called on time\n",
		},
		{
			Name:       "Wrong password supplied",
			Args:       []string{"Time-and-expence-registration", "reset", "-c", "true", "-t", "All", "-p", "lol"},
			WantOutput: "Incorrect password supplied\n",
		},
		{
			Name:       "Confirm flag not set correctly",
			Args:       []string{"Time-and-expence-registration", "reset", "-c", "false", "-t", "All", "-p", "Odin2203!"},
			WantOutput: "Confirm flag not set correctly\n",
		},
		{
			Name:       "Incorrect type flag",
			Args:       []string{"Time-and-expence-registration", "reset", "-c", "true", "-t", "all", "-p", "Odin2203!"},
			WantOutput: "Incorrect use of Type flag. Use either Finance, Time or All. Ensure correct capitalisation\n",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			originalStdout := os.Stdout

			os.Stdout = w
			rootCmd.SetArgs(test.Args[1:])
			rootCmd.Execute()

			w.Close()
			var buf bytes.Buffer
			io.Copy(&buf, r)
			os.Stdout = originalStdout
			output := buf.String()

			if output != test.WantOutput {
				fmt.Printf("%s failed:\n", test.Name)
				fmt.Printf("output: %s expected output: %s\n", output, test.WantOutput)
				t.Fail()
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
		})

	}
}

func TestRegisterTime(t *testing.T) {
	tests := []CLITest{
		{
			Name: "First time registration",
			Args: []string{"Time-and-expence-registration", "registerTime",
				"-d", "15-01-2025",
				"-t", "30",
				"-c", "honney harvest",
				"-e", "Harvest 4,5 kg of honney from hive#1"},
			WantOutput: "15-01-2025 honney harvest 4,5 hive#1",
		},
		{
			Name: "Second time registration",
			Args: []string{"Time-and-expence-registration", "registerTime",
				"-d", "31-05-2024",
				"-t", "30",
				"-c", "maintenance",
				"-e", "weekly check on hive#2"},
			WantOutput: "31-05-2024 maintenance  hive#2",
		},
	}

	err := querry.ResetTimeRegistration(context.Background())
	if err != nil {
		fmt.Printf("Error during time database reset prior to test start: %s", err)
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			originalStdout := os.Stdout

			os.Stdout = w
			rootCmd.SetArgs(test.Args[1:])
			err = rootCmd.Execute()

			w.Close()
			var buf bytes.Buffer
			io.Copy(&buf, r)
			os.Stdout = originalStdout
			output := buf.String()
			fmt.Printf("\n Output: %s", output)
			for _, word := range strings.Split(test.WantOutput, " ") {
				if !strings.Contains(output, word) {
					fmt.Printf("\nTest failed: %s is not in %s", word, output)
					t.Fail()
				}

			}
		})
	}
}
