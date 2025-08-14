/*
Copyright © 2025 Jasper Olthof-Donker <jasper.olthof@xs4all.nl>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/Dirza1/TimeTally/internal/utils"
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
	This action can only be performed by an administrator.
	Set the ID of the registration to be deleted and which database it needs to be deleted from.`,
	Run: func(cmd *cobra.Command, args []string) {
		if deleteEntryType == "" {
			fmt.Println("-t or --type flag not set. Please set this flag")
		}
		if deleteEntryId == "" {
			fmt.Println("-e or --entry-id flag not set. Please set this flag")
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
			utils.UpdateSession()
		case "Time":
			err := queries.DeleteTime(context.Background(), ID)
			if err != nil {
				fmt.Printf("\nerror during deletion: %s \n", err)
				return
			}
			fmt.Println("Entry deleted")
			utils.UpdateSession()
		default:
			fmt.Println("Incorrect use of the -t/ --Time flag. Use Finance or Time after the flag. Be mindfull of capitalisation.")
		}
		fmt.Println("deleteEntry called")
	},
}

func init() {
	rootCmd.AddCommand(deleteEntryCmd)

	deleteEntryCmd.Flags().StringVarP(&deleteEntryType, "type", "t", "", "A flag to diferatiate between the databases. Use either Financial or Time after the flag")

	deleteEntryCmd.Flags().StringVarP(&deleteEntryId, "entry-id", "e", "", "A flag to set the ID of the registrations that needs to be deleted")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteEntryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteEntryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
