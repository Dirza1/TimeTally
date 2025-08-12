/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// deleteEntryCmd represents the deleteEntry command
var deleteEntryType string
var deleteEntryId string

var deleteEntryCmd = &cobra.Command{
	Use:   "deleteEntry",
	Short: "Deletes a entry from the Time or Financial database",
	Long: `This command deletes a registration from the Time or Financial database.
	Set the ID of the registration to be deleted`,
	Run: func(cmd *cobra.Command, args []string) {
		if deleteEntryType == "" {
			fmt.Println("-t or --type flag not set. Please set this flag")
		}
		if deleteEntryId == "" {
			fmt.Println("-i or --id flag not set. Please set this flag")
		}

		ID, err := uuid.Parse(deleteEntryId)
		if err != nil {
			fmt.Printf("\nerror during parsing of ID: %s \n", err)
			return
		}
		queries := utils.DatabaseConnection()
		currentUser, err := utils.LoadSession()
		if err != nil {
			fmt.Printf("\nError retrieving current user from session. Err:\n%s\n", err)
		}
		permissions, err := queries.GetUserPermissions(context.Background(), currentUser.UserName)
		if err != nil {
			fmt.Printf("\nError during retrieval of user permissions from database. Err:\n%s\n", err)
			return
		}
		if permissions.Administrator != true {
			fmt.Println("Current user is not an administrator")
			return
		}
		switch deleteEntryType {
		case "Financial":
			err := queries.DeleteTransaction(context.Background(), ID)
			if err != nil {
				fmt.Printf("\nerror during deletion: %s \n", err)
				return
			}
			fmt.Println("Entry deleted")
		case "Time":
			err := queries.DeleteTime(context.Background(), ID)
			if err != nil {
				fmt.Printf("\nerror during deletion: %s \n", err)
				return
			}
			fmt.Println("Entry deleted")
		default:
			fmt.Println("Incorrect use of the -t/ --Time flag. Use Finance or Time after the flag. Be mindfull of capitalisation.")
		}
		fmt.Println("deleteEntry called")
	},
}

func init() {
	rootCmd.AddCommand(deleteEntryCmd)

	deleteEntryCmd.Flags().StringVarP(&deleteEntryType, "type", "t", "", "A flag to diferatiate between the databases. Use either Financial or Time after the flag")

	deleteEntryCmd.Flags().StringVarP(&deleteEntryId, "id", "i", "", "A flag to set the ID of the registrations that needs to be deleted")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteEntryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteEntryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
