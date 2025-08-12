/*
Copyright © 2025 Jasper Olthof-Donker <jasper.olthof@xs4all.nl>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/spf13/cobra"
)

// UserOverviewCmd represents the UserOverview command
var UserOverviewCmd = &cobra.Command{
	Use:   "UserOverview",
	Short: "Genmerate a user overview",
	Long: `This command generates a overview of the active users in the user database.
	It will return a list of all users, id's and privileges.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		queries := utils.DatabaseConnection()
		users, err := queries.UserOverview(context.Background())
		if err != nil {
			fmt.Printf("\nError retrieving users from database. Err:\n%s\n", err)
			return
		}
		fmt.Println("Current users:")
		for _, user := range users {
			fmt.Printf("Name: %s, ID: %s, Time access: %t, Financial access: %t, Administrator: %t\n",
				user.Name,
				user.ID,
				user.AccessTimeregistration,
				user.AccessFinance,
				user.Administrator)
		}
		utils.UpdateSession()
	},
}

func init() {
	rootCmd.AddCommand(UserOverviewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UserOverviewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UserOverviewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
