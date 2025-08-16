package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Dirza1/TimeTally/internal/database"
	"github.com/Dirza1/TimeTally/internal/utils"
)

var querry database.Queries

type CLITest struct {
	Name       string
	Args       []string
	WantOutput string
	NotWanted  string
}

func TestMain(m *testing.M) {
	//Tests have to be performed with a fully empty database. No content or users should be in there.
	//Please delete all users prior to starting the tests. After the tests re-create a first admin with FirstAdmin as it will be deleted by the tests
	//User UserOverview to see if there are users in the databases prior to testing

	fmt.Println("Starting tests and database connection")
	//Set up to root directory to ensure relative pathways are still correct
	dir, _ := os.Getwd()
	root := filepath.Join(dir, "..")
	os.Chdir(root)

	//loggin in into the database
	Args := []string{"TimeTally", "FirstAdmin", "-u", "TestAdmin", "-p", "Test"}
	rootCmd.SetArgs(Args[1:])
	rootCmd.Execute()

	Args = []string{"TimeTally", "Login", "-u", "TestAdmin", "-p", "Test"}
	rootCmd.SetArgs(Args[1:])
	rootCmd.Execute()

	// setting up the database connection
	querry = utils.DatabaseConnection()

	//Starting the tests
	exitCode := m.Run()

	fmt.Println("Ending tests and resetting database")
	//Cleaning up both databases
	_ = querry.ResetTransaction(context.Background())
	_ = querry.ResetTimeRegistration(context.Background())

	Args = []string{"TimeTally", "DeleteUser", "-n", "TestAdmin", "-p", "Test"}
	rootCmd.SetArgs(Args[1:])
	rootCmd.Execute()

	os.Exit(exitCode)

}

func runCLI(t *testing.T, args []string) (string, error) {
	t.Helper()
	r, w, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = w

	rootCmd.SetArgs(args)
	err := rootCmd.Execute()

	w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdout

	return buf.String(), err
}

func TestReset(t *testing.T) {
	tests := []CLITest{
		{
			Name:       "Test reset of both databases",
			Args:       []string{"TimeTally", "reset", "-c", "true", "-t", "All", "-p", "Odin2203!"},
			WantOutput: "reset called on all\n",
		},
		{
			Name:       "Test reset of Finance database",
			Args:       []string{"TimeTally", "reset", "-c", "true", "-t", "Finance", "-p", "Odin2203!"},
			WantOutput: "reset called on finance\n",
		},
		{
			Name:       "Test reset of Time database",
			Args:       []string{"TimeTally", "reset", "-c", "true", "-t", "Time", "-p", "Odin2203!"},
			WantOutput: "reset called on time\n",
		},
		{
			Name:       "Wrong password supplied",
			Args:       []string{"TimeTally", "reset", "-c", "true", "-t", "All", "-p", "lol"},
			WantOutput: "Incorrect password supplied\n",
		},
		{
			Name:       "Confirm flag not set correctly",
			Args:       []string{"TimeTally", "reset", "-c", "false", "-t", "All", "-p", "Odin2203!"},
			WantOutput: "Confirm flag not set correctly\n",
		},
		{
			Name:       "Incorrect type flag",
			Args:       []string{"TimeTally", "reset", "-c", "true", "-t", "all", "-p", "Odin2203!"},
			WantOutput: "Incorrect use of Type flag. Use either Finance, Time or All. Ensure correct capitalisation\n",
		},
		{
			Name: "First user adding",
			Args: []string{"TimeTally", "AddUser",
				"-u", "Test-1",
				"-p", "Test-1",
				"-t", "-f",
			},
			WantOutput: "New user created",
			NotWanted:  "Error",
		},
		{
			Name: "Test-1 login",
			Args: []string{"TimeTally", "Login",
				"-u", "Test-1",
				"-p", "Test-1",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name:       "Test reset no admin",
			Args:       []string{"TimeTally", "reset", "-c", "true", "-t", "All", "-p", "Odin2203!"},
			WantOutput: "Current user is not an administrator",
		},
		{
			Name: "Test-1 login",
			Args: []string{"TimeTally", "Login",
				"-u", "TestAdmin",
				"-p", "Test",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name: "delete Test-1 ",
			Args: []string{"TimeTally", "DeleteUser",
				"-n", "Test-1",
			},
			WantOutput: "User deleted.",
			NotWanted:  "Error",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			output, err := runCLI(t, test.Args[1:])

			for _, word := range strings.Split(test.WantOutput, " ") {
				if !strings.Contains(output, word) {
					fmt.Printf("\n%s failed: %s is not in %s", test.Name, word, output)
					t.Fail()
				}

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
			Args: []string{"TimeTally", "registerTime",
				"-d", "15-01-2025",
				"-t", "30",
				"-c", "honney harvest",
				"-e", "Harvest 4,5 kg of honney from hive#1"},
			WantOutput: "15-01-2025 honney harvest 4,5 hive#1",
		},
		{
			Name: "Second time registration",
			Args: []string{"TimeTally", "registerTime",
				"-d", "31-05-2024",
				"-t", "30",
				"-c", "maintenance",
				"-e", "weekly check on hive#2"},
			WantOutput: "31-05-2024 maintenance  hive#2",
		},
		{
			Name: "user adding",
			Args: []string{"TimeTally", "AddUser",
				"-u", "Test-1",
				"-p", "Test-1",
			},
			WantOutput: "New user created",
			NotWanted:  "Error",
		},
		{
			Name: "Login second user",
			Args: []string{"TimeTally", "Login",
				"-u", "Test-1",
				"-p", "Test-1",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name: "Register Withouth privelages",
			Args: []string{"TimeTally", "registerTime",
				"-d", "31-05-2024",
				"-t", "30",
				"-c", "maintenance",
				"-e", "weekly check on hive#2"},
			WantOutput: "Current user is not allowed in the time registartion database",
			NotWanted:  "Hive#2",
		},
		{
			Name: "Admin login",
			Args: []string{"TimeTally", "Login",
				"-u", "TestAdmin",
				"-p", "Test",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name: "delete Test-1 ",
			Args: []string{"TimeTally", "DeleteUser",
				"-n", "Test-1",
			},
			WantOutput: "User deleted.",
			NotWanted:  "Error",
		},
	}

	err := querry.ResetTimeRegistration(context.Background())
	if err != nil {
		fmt.Printf("Error during time database reset prior to test start: %s", err)
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			output, _ := runCLI(t, test.Args[1:])

			for _, word := range strings.Split(test.WantOutput, " ") {
				if !strings.Contains(output, word) {
					fmt.Printf("\n%s failed: %s is not in %s", test.Name, word, output)
					t.Fail()
				}

			}
		})
	}
}

func TestRegisterTransaction(t *testing.T) {
	tests := []CLITest{
		{
			Name: "First transaction registration",
			Args: []string{"TimeTally", "registerTransaction",
				"-d", "15-01-2025",
				"-a", "300",
				"-t", "spent",
				"-c", "glass bottles",
				"-e", "bought glass bottles for honney"},
			WantOutput: "15-01-2025 300 spent glass bottles bought glass bottles for honney",
		},
		{
			Name: "Second transaction registration",
			Args: []string{"TimeTally", "registerTransaction",
				"-d", "16-02-2026",
				"-a", "310",
				"-t", "gained",
				"-c", "bees",
				"-e", "stuff to capture a swarm"},
			WantOutput: "16-02-2026 310 gained bees stuff to capture a swarm",
		},
	}

	err := querry.ResetTransaction(context.Background())
	if err != nil {
		fmt.Printf("Error during time database reset prior to test start: %s", err)
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			output, err := runCLI(t, test.Args[1:])

			for _, word := range strings.Split(test.WantOutput, " ") {
				if !strings.Contains(output, word) && err == nil {
					fmt.Printf("\nTest failed: %s is not in %s", word, output)
					t.Fail()

				} else if err != nil && !strings.Contains(err.Error(), word) {

					fmt.Printf("\nTest failed: %s is not in %s", word, output)
					t.Fail()
				}

			}

		})
	}
}

func TestOverview(t *testing.T) {
	tests := []CLITest{
		{
			Name: "first transaction addition",
			Args: []string{"TimeTally", "registerTransaction",
				"-d", "16-02-2026",
				"-a", "300",
				"-t", "spent",
				"-c", "glass bottles",
				"-e", "bought glass bottles"},
			WantOutput: "16-02-2026. 300, spent glass bottles, bought glass bottles,",
			NotWanted:  "Test",
		},
		{
			Name: "First time registration",
			Args: []string{"TimeTally", "registerTime",
				"-d", "15-01-2025",
				"-t", "30",
				"-c", "honney harvest",
				"-e", "Harvest 4,5 kg of honney from hive#1"},
			WantOutput: "15-01-2025. honney harvest, 4,5 hive#1,",
			NotWanted:  "Test",
		},
		{
			Name: "First time overview",
			Args: []string{"TimeTally", "overview",
				"-t", "Time"},
			WantOutput: "15-01-2025. honney harvest, 4,5 hive#1,",
			NotWanted:  "16-02-2026. 300, glass bottles, bought glass bottles,",
		},
		{
			Name: "First finance overview",
			Args: []string{"TimeTally", "overview",
				"-t", "Finance"},
			WantOutput: "16-02-2026. 3.00 glass bottles, bought glass bottles,",
			NotWanted:  "15-01-2025. honney harvest, 4,5 hive#1,",
		},
		{
			Name: "Incorrect flag used",
			Args: []string{"TimeTally", "overview",
				"-t", "Incorrect"},
			WantOutput: "Incorrect use of the -t/ --Type flag. Use Finance or Time after the flag. Be mindfull of capitalisation.",
			NotWanted:  "15-01-2025. honney harvest, 4,5 hive#1, 16-02-2026. 300, glass bottles, bought glass bottles,",
		},
		{
			Name: "Second time registration",
			Args: []string{"TimeTally", "registerTime",
				"-d", "25-11-3025",
				"-t", "11",
				"-c", "inspection",
				"-e", "inspected hive#2"},
			WantOutput: "25-11-3025. inspection, inspected hive#2,",
			NotWanted:  "Test",
		},
		{
			Name: "second transaction addition",
			Args: []string{"TimeTally", "registerTransaction",
				"-d", "29-07-2126",
				"-a", "254",
				"-t", "gained",
				"-c", "sale",
				"-e", "soled expertice"},
			WantOutput: "29-07-2126. 254, sale, soled expertice,",
			NotWanted:  "Test",
		},
		{
			Name: "second finance overview",
			Args: []string{"TimeTally", "overview",
				"-t", "Finance"},
			WantOutput: "16-02-2026. 3.00 glass bottles, bought glass bottles, 29-07-2126. 2.54  sale, soled expertice,",
			NotWanted:  "15-01-2025. honney harvest, 4,5 hive#1, 25-11-3025. inspection, inspected hive#2,",
		},
		{
			Name: "second time overview",
			Args: []string{"TimeTally", "overview",
				"-t", "Time"},
			WantOutput: "15-01-2025. honney harvest, 4,5 hive#1, 25-11-3025. inspection, inspected hive#2,",
			NotWanted:  "16-02-2026. 300, glass bottles, bought glass bottles, 29-07-2126. 2.54, sale, soled expertice,",
		},
	}
	err := querry.ResetTransaction(context.Background())
	if err != nil {
		fmt.Printf("Error during time database reset prior to test start: %s", err)
	}
	err = querry.ResetTimeRegistration(context.Background())
	if err != nil {
		fmt.Printf("Error during financial database reset prior to test start: %s", err)
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			output, err := runCLI(t, test.Args[1:])

			for _, word := range strings.Split(test.WantOutput, " ") {
				if !strings.Contains(output, word) && err == nil {
					fmt.Printf("Test failed. %s is not in %s", word, output)
					t.Fail()
				} else if err != nil && strings.Contains(err.Error(), word) {
					fmt.Printf("Test failed. %s is not in %s", word, err.Error())
					t.Fail()
				}
			}
			for _, word := range strings.Split(test.NotWanted, " ") {
				if strings.Contains(output, word) && err == nil {
					fmt.Printf("Test failed. %s should not be in %s", word, output)
					t.Fail()
				} else if err != nil && strings.Contains(err.Error(), word) {
					fmt.Printf("Test failed. %s is not in %s", word, err.Error())
					t.Fail()
				}
			}
		})
	}
}

func TestAddAdmin(t *testing.T) {
	tests := []CLITest{
		{
			Name: "first Admin adding",
			Args: []string{"TimeTally", "AddAdmin",
				"-u", "Test-1",
				"-p", "Test-1",
			},
			WantOutput: "New Administrator created.",
			NotWanted:  "Error",
		},
		{
			Name: "delete first Admin adding",
			Args: []string{"TimeTally", "DeleteUser",
				"-n", "Test-1",
			},
			WantOutput: "User deleted.",
			NotWanted:  "Error",
		},
		{
			Name:       "logout",
			Args:       []string{"TimeTally", "Logout"},
			WantOutput: "User logged out",
			NotWanted:  "Error",
		},
		{

			Name: "second Admin adding while logged out",
			Args: []string{"TimeTally", "AddAdmin",
				"-u", "Test-1",
				"-p", "Test-1",
			},
			WantOutput: "Users session expired. Please use the login command to continue using the system",
			NotWanted:  "New Administrator created.",
		},
		{
			Name: "Login admin user",
			Args: []string{"TimeTally", "Login",
				"-u", "TestAdmin",
				"-p", "Test",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name: "Add a non admin user",
			Args: []string{"TimeTally", "AddUser",
				"-u", "Test-1",
				"-p", "Test-1",
				"-f", "false",
				"-t", "false",
			},
			WantOutput: "New user created.",
			NotWanted:  "Administrator",
		},
		{
			Name: "Login non admin user",
			Args: []string{"TimeTally", "Login",
				"-u", "Test-1",
				"-p", "Test-1",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name: "Add admin from non admin user",
			Args: []string{"TimeTally", "AddAdmin",
				"-u", "Test-2",
				"-p", "Test-2",
			},
			WantOutput: "Current user is not an administrator",
			NotWanted:  "Test-2",
		},
		{
			Name: "Login admin user",
			Args: []string{"TimeTally", "Login",
				"-u", "TestAdmin",
				"-p", "Test",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name: "Add second user with same name",
			Args: []string{"TimeTally", "AddAdmin",
				"-u", "Test-1",
				"-p", "Test-1",
			},
			WantOutput: "User already exists. Please use a diferent username",
			NotWanted:  "Error",
		},
		{
			Name: "delete first Admin",
			Args: []string{"TimeTally", "DeleteUser",
				"-n", "Test-1",
			},
			WantOutput: "User deleted.",
			NotWanted:  "Error",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			output, err := runCLI(t, test.Args[1:])

			for _, word := range strings.Split(test.WantOutput, " ") {
				if !strings.Contains(output, word) && err == nil {
					fmt.Printf("%s failed. %s is not in %s", test.Name, word, output)
					t.Fail()
				} else if err != nil && strings.Contains(err.Error(), word) {
					fmt.Printf("%s failed. %s is not in %s", test.Name, word, err.Error())
					t.Fail()
				}
			}
			for _, word := range strings.Split(test.NotWanted, " ") {
				if strings.Contains(output, word) && err == nil {
					fmt.Printf("%s failed. %s should not be in %s", test.Name, word, output)
					t.Fail()
				} else if err != nil && strings.Contains(err.Error(), word) {
					fmt.Printf("%s failed. %s is not in %s", test.Name, word, err.Error())
					t.Fail()
				}
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	tests := []CLITest{
		{
			Name: "Admin login",
			Args: []string{"TimeTally", "Login",
				"-u", "TestAdmin",
				"-p", "Test",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name: "First user adding",
			Args: []string{"TimeTally", "AddUser",
				"-u", "Test-1",
				"-p", "Test-1",
				"-t", "-f",
			},
			WantOutput: "New user created",
			NotWanted:  "Error",
		},
		{
			Name: "Second user adding",
			Args: []string{"TimeTally", "AddUser",
				"-u", "Test-2",
				"-p", "Test-2",
				"-f",
			},
			WantOutput: "New user created",
			NotWanted:  "Error",
		},
		{
			Name: "Third user adding",
			Args: []string{"TimeTally", "AddUser",
				"-u", "Test-3",
				"-p", "Test-3",
				"-t",
			},
			WantOutput: "New user created",
			NotWanted:  "Error",
		},
		{
			Name: "Test-1 login",
			Args: []string{"TimeTally", "Login",
				"-u", "Test-1",
				"-p", "Test-1",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name: "Correct overview test",
			Args: []string{"TimeTally", "overview",
				"-t", "Time",
			},
			WantOutput: "Overview of the Timeregistrations:",
			NotWanted:  "Error",
		},
		{
			Name: "Test-2 login",
			Args: []string{"TimeTally", "Login",
				"-u", "Test-2",
				"-p", "Test-2",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name: "Incorrect overview test-2 test",
			Args: []string{"TimeTally", "overview",
				"-t", "Time",
			},
			WantOutput: "Current user is not allowed in the time registration databse",
			NotWanted:  "Overview",
		},
		{
			Name: "Test-3 login",
			Args: []string{"TimeTally", "Login",
				"-u", "Test-3",
				"-p", "Test-3",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name: "Incorrect overview test-3 test",
			Args: []string{"TimeTally", "overview",
				"-t", "Finance",
			},
			WantOutput: "Current user is not allowed in the financial database",
			NotWanted:  "Overview",
		},
		{
			Name: "Admin login",
			Args: []string{"TimeTally", "Login",
				"-u", "TestAdmin",
				"-p", "Test",
			},
			WantOutput: "Login Successful",
			NotWanted:  "Error",
		},
		{
			Name: "delete Test-1 ",
			Args: []string{"TimeTally", "DeleteUser",
				"-n", "Test-1",
			},
			WantOutput: "User deleted.",
			NotWanted:  "Error",
		},
		{
			Name: "delete Test-2 ",
			Args: []string{"TimeTally", "DeleteUser",
				"-n", "Test-2",
			},
			WantOutput: "User deleted.",
			NotWanted:  "Error",
		},
		{
			Name: "delete Test-3 ",
			Args: []string{"TimeTally", "DeleteUser",
				"-n", "Test-3",
			},
			WantOutput: "User deleted.",
			NotWanted:  "Error",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			newUserAccessFinance = false
			newUserAccessTime = false

			output, err := runCLI(t, test.Args[1:])

			for _, word := range strings.Split(test.WantOutput, " ") {
				if !strings.Contains(output, word) && err == nil {
					fmt.Printf("%s failed. %s is not in %s", test.Name, word, output)
					t.Fail()
				} else if err != nil && strings.Contains(err.Error(), word) {
					fmt.Printf("%s failed. %s is not in %s", test.Name, word, err.Error())
					t.Fail()
				}
			}
			for _, word := range strings.Split(test.NotWanted, " ") {
				if strings.Contains(output, word) && err == nil {
					fmt.Printf("%s failed. %s should not be in %s", test.Name, word, output)
					t.Fail()
				} else if err != nil && strings.Contains(err.Error(), word) {
					fmt.Printf("%s failed. %s is not in %s", test.Name, word, err.Error())
					t.Fail()
				}
			}
		})
	}
}
