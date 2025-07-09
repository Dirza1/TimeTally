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
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var ResetType string
var ResetConfirm bool
var ResetPassword string

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Removes all data from the finance and or timeregistration ddatabase",
	Long: `This command removes all data from the finance and or timeregistration database.
	NOTE: This is a permenant action. The data will be lost and if there are no backups the data will be lost!
	Execution of this command requires an additional password to protect against axidental use.
	Additionaly this command uses 3 required flags for the database to be deleted, confirming to delete and the password.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ResetConfirm != true {
			log.Fatal("Confirm flag not set correctly.")
		}
		godotenv.Load("/home/jasperolthof/workspace/github.com/Dirza1/Time-and-expence-registration/.env")

		setPasword := os.Getenv("reset_password")
		if setPasword != ResetPassword {
			log.Fatal("Incorrect password supplied")
		}
		queries := utils.DatabaseConnection()
		switch ResetType {
		case "Finance":
			err := queries.ResetTransaction(context.Background())
			if err != nil {
				log.Fatal("error during reset: ", err)
			}
			fmt.Println("reset called on financee")
		case "Time":
			err := queries.ResetTimeRegistration(context.Background())
			if err != nil {
				log.Fatal("error during reset: ", err)
			}
			fmt.Println("reset called on time")
		case "All":
			err := queries.ResetTimeRegistration(context.Background())
			if err != nil {
				log.Fatal("error during reset: ", err)
			}
			err = queries.ResetTransaction(context.Background())
			if err != nil {
				log.Fatal("error during reset: ", err)
			}
			fmt.Println("reset called on all")
		default:
			fmt.Println("Incorrect use of Type flag. Use either Finance, Time or All. Ensure correct capitalisation")
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	resetCmd.Flags().StringVarP(&ResetType, "type", "t", "all", "Used to diferenciate which database needs to be deleted. Use Finance, Time or all as input. (required)")
	resetCmd.MarkFlagRequired("type")

	resetCmd.Flags().BoolVarP(&ResetConfirm, "confirm", "c", false, "Used to confirm the database delition. Type true after the flag (required)")
	resetCmd.MarkFlagRequired("confirm")

	resetCmd.Flags().StringVarP(&ResetPassword, "password", "p", "", "Additional password required for delition. (required)")
	resetCmd.MarkFlagRequired("password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
