/*
Copyright © 2025 Jasper Olthof-Donker <jasper.olthof@xs4all.nl>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var toDeleteID string

// DeleteUserCmd represents the DeleteUser command
var DeleteUserCmd = &cobra.Command{
	Use:   "DeleteUser",
	Short: "This command deletes a user",
	Long: `If a user is no longer required to have access to this program their account can be deleted.
	This action can only be performed by an administrator.
	Note this only removed a user from the user database. Entries in the Financial or Time database will still exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		if toDeleteID == "" {
			fmt.Println("-i or --user-id flag not set. Provide the ID of the user to be deleted")
			return
		}
		session, err := utils.LoadSession()
		if err != nil {
			fmt.Printf("\n Error loading user session. Err:\n%s\n", err)
			return
		}
		currentTime := time.Now()
		if currentTime.Sub(session.LastUsed) > 15*time.Minute {
			fmt.Println("Users session expired. Please use the login command to continue using the system")
			return
		}

		queries := utils.DatabaseConnection()
		permissions, err := queries.GetUserPermissions(context.Background(), session.UserName)
		if err != nil {
			fmt.Printf("\nError retrieving user permissions. Err:\n%s\n", err)
			return
		}
		if permissions.Administrator != true {
			fmt.Println("Current user is not an administrator.")
			return
		}
		ID, err := uuid.Parse(toDeleteID)
		if err != nil {
			fmt.Printf("\nerror during parsing of the ID: %s \n", err)
			return
		}

		err = queries.DeleteUser(context.Background(), ID)
		if err != nil {
			fmt.Printf("\nError during deletion of user. Err:\n%s\n", err)
			return
		}
		fmt.Println("User deleted.")
		utils.UpdateSession()
	},
}

func init() {
	rootCmd.AddCommand(DeleteUserCmd)

	deleteEntryCmd.Flags().StringVarP(&toDeleteID, "user-id", "i", "", "The ID of the user to be deleted.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// DeleteUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// DeleteUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
