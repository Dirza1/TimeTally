/*
Copyright © 2025 Jasper Olthof-Donker <jasper.olthof@xs4all.nl>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// resetFinanceCmd represents the resetFinance command
var resetFinanceCmd = &cobra.Command{
	Use:   "resetFinance",
	Short: "Removes all data from the finance ddatabase",
	Long: `This command removes all data from the finance database.
	NOTE: This is a permenant action. The data will be lost and if there are no backups the data will be lost!
	Execution of this command requires an additional password to protect against axidental use.`,
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load(".env")
		setPasword := os.Getenv("reset_password")
		var input string
		fmt.Println(`This is a PERMANENT action that deletes all data.
		If this was a axidental execute, please press enter.
		Else provide your delete password.`)
		_, _ = fmt.Scanln(&input)

		if input != setPasword {
			fmt.Println("Wrong password")
			return
		} else {
			queries := utils.DatabaseConnection()
			err := queries.ResetTransaction(context.Background())
			if err != nil {
				log.Fatal("error during reset: ", err)
			}
			fmt.Println("Financial data deleted")

		}
	},
}

func init() {
	rootCmd.AddCommand(resetFinanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetFinanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetFinanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
