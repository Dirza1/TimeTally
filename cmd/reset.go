/*
Copyright © 2025 Jasper Olthof-Donker <jasper.olthof@xs4all.nl>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Dirza1/TimeTally/internal/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command

var resetType string
var resetConform string
var resetPassword string

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Removes all data from the finance and or timeregistration ddatabase",
	Long: `This command removes all data from the finance and or timeregistration database.
	NOTE: This is a permenant action. The data will be lost and if there are no backups the data will be lost!
	Execution of this command requires an additional password to protect against axidental use.
	Additionaly this command uses 3 required flags for the database to be deleted, confirming to delete and the password.`,
	Run: func(cmd *cobra.Command, args []string) {
		if resetConform == "" {
			fmt.Printf("\n-c or --confirm flag not set. Please set the confirm flag\n")
			return
		}
		if resetType == "" {
			fmt.Printf("\n-t or --type flag not set. Please set the type flag\n")
			return
		}
		if resetPassword == "" {
			fmt.Printf("\n-p or --pasword flag not set. Please provide the reset pasword\n")
		}
		session, err := utils.LoadSession()
		if err != nil {
			fmt.Printf("\nError loading session. Err:\n%s\n", err)
			return
		}
		currentTime := time.Now()
		if currentTime.Sub(session.LastUsed) > 15*time.Minute {
			fmt.Println("Users session expired. Please use the login command to continue using the system")
			return
		}

		if resetConform != "true" {
			fmt.Println("Confirm flag not set correctly")
			return
		}
		err = godotenv.Load(".env")
		if err != nil {
			fmt.Printf("\nError loading enviromental variables. Err: \n%s\n", err)
			return
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			fmt.Printf("\nPassword flag error. Err: \n%s\n", err)
			return
		}
		setPasword := os.Getenv("reset_password")
		if setPasword != password {
			fmt.Println("Incorrect password supplied")
			return
		}
		queries := utils.DatabaseConnection()

		currentUser, err := utils.LoadSession()
		if err != nil {
			fmt.Printf("\nError retrieving current user from session. Err: \n%s\n", err)
			return
		}
		permissions, err := queries.GetUserPermissions(context.Background(), currentUser.UserName)
		if err != nil {
			fmt.Printf("\nError during retrieval of user permissions from database. Err: \n%s\n", err)
			return
		}
		if permissions.Administrator != true {
			fmt.Println("Current user is not an administrator")
			return
		}

		resetType, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Printf("\nType flag error. Err: \n%s\n", err)
			return
		}

		switch resetType {
		case "Finance":
			err := queries.ResetTransaction(context.Background())
			if err != nil {
				fmt.Printf("\nerror during reset: %s \n", err)
				return
			}
			fmt.Println("reset called on finance")
			utils.UpdateSession()
		case "Time":
			err := queries.ResetTimeRegistration(context.Background())
			if err != nil {
				fmt.Printf("\nerror during reset: %s \n", err)
				return
			}
			fmt.Println("reset called on time")
			utils.UpdateSession()
		case "All":
			err := queries.ResetTimeRegistration(context.Background())
			if err != nil {
				fmt.Printf("\nerror during reset: %s \n", err)
				return
			}
			err = queries.ResetTransaction(context.Background())
			if err != nil {
				fmt.Printf("\nerror during reset: %s \n", err)
				return
			}
			fmt.Println("reset called on all")

		default:
			fmt.Println("Incorrect use of Type flag. Use either Finance, Time or All. Ensure correct capitalisation")
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	resetCmd.Flags().StringVarP(&resetType, "type", "t", "", "Used to diferenciate which database needs to be deleted. Use Finance, Time or all as input. (required)")

	resetCmd.Flags().StringVarP(&resetConform, "confirm", "c", "", "Used to confirm the database delition. Type true after the flag (required)")

	resetCmd.Flags().StringVarP(&resetPassword, "password", "p", "", "Additional password required for delition. (required)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
