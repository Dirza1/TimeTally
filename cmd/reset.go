/*
Copyright © 2025 Jasper Olthof-Donker <jasper.olthof@xs4all.nl>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Removes all data from the finance and or timeregistration ddatabase",
	Long: `This command removes all data from the finance and or timeregistration database.
	NOTE: This is a permenant action. The data will be lost and if there are no backups the data will be lost!
	Execution of this command requires an additional password to protect against axidental use.
	Additionaly this command uses 3 required flags for the database to be deleted, confirming to delete and the password.`,
	Run: func(cmd *cobra.Command, args []string) {
		confirm, err := cmd.Flags().GetString("confirm")
		if err != nil || confirm != "true" {
			fmt.Println("Confirm flag not set correctly")
			return
		}
		err = godotenv.Load("/home/jasperolthof/workspace/projects/Time-and-expence-registration/.env")
		if err != nil {
			fmt.Printf("Error loading enviromental variables")
			return
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			fmt.Println("Password flag error")
			return
		}
		setPasword := os.Getenv("reset_password")
		if setPasword != password {
			fmt.Println("Incorrect password supplied")
			return
		}
		queries := utils.DatabaseConnection()

		resetType, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Println("Type flag error")
			return
		}

		switch resetType {
		case "Finance":
			err := queries.ResetTransaction(context.Background())
			if err != nil {
				fmt.Printf("error during reset: %s \n", err)
				return
			}
			fmt.Println("reset called on finance")
		case "Time":
			err := queries.ResetTimeRegistration(context.Background())
			if err != nil {
				fmt.Printf("error during reset: %s \n", err)
				return
			}
			fmt.Println("reset called on time")
		case "All":
			err := queries.ResetTimeRegistration(context.Background())
			if err != nil {
				fmt.Printf("error during reset: %s \n", err)
				return
			}
			err = queries.ResetTransaction(context.Background())
			if err != nil {
				fmt.Printf("error during reset: %s \n", err)
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

	resetCmd.Flags().StringP("type", "t", "", "Used to diferenciate which database needs to be deleted. Use Finance, Time or all as input. (required)")
	err := resetCmd.MarkFlagRequired("type")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	resetCmd.Flags().StringP("confirm", "c", "", "Used to confirm the database delition. Type true after the flag (required)")
	err = resetCmd.MarkFlagRequired("confirm")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	resetCmd.Flags().StringP("password", "p", "", "Additional password required for delition. (required)")
	err = resetCmd.MarkFlagRequired("password")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
