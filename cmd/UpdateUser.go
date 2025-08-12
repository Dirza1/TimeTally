/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/spf13/cobra"
)

var updateUserUsername string
var updateUserPassword string
var updateUserAccessFinance bool
var updateUserAccessTime bool
var updateUserAdministrator bool

// UpdateUserCmd represents the UpdateUser command
var UpdateUserCmd = &cobra.Command{
	Use:   "UpdateUser",
	Short: "This command updates user data.",
	Long: `Use this command to:
	- Update a users name
	- Reset a users password
	- Update a users privaleges
	This command can also be used to make a user administrator
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if updateUserUsername == "" {
			fmt.Println("-u or --username flag not set. Please set this flag.")
			return
		}
		if updateUserPassword == "" {
			fmt.Println("-p or --newPassword flag not set. Please set this flag.")
			return
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

		queries := utils.DatabaseConnection()
		currentUser, err := utils.LoadSession()
		if err != nil {
			fmt.Printf("\nError retrieving current user from session. Err:\n%s\n", err)
			return
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

		_, err = queries.GetUserPermissions(context.Background(), updateUserUsername)
		if err == nil {
			fmt.Println("Username already exists. Please use another user name or delete old user")
			return
		}

		hashedPasword, err := utils.Hashpassword(updateUserPassword)
		if err != nil {
			fmt.Printf("\nError during pasword hash. Err:\n%s\n", err)
			return
		}
		newUser := database.UpdateUserParams{
			Name:                   newUserUsername,
			HashedPassword:         hashedPasword,
			AccessFinance:          updateUserAccessFinance,
			AccessTimeregistration: updateUserAccessTime,
			Administrator:          updateUserAdministrator,
		}
		createdUser, err := queries.UpdateUser(context.Background(), newUser)
		if err != nil {
			fmt.Printf("\nError duing user creation. Err:\n%s\n", err)
			return
		}
		fmt.Printf("User updates. created. Name: %s, Financial access: %t, Time access: %t. Administrator access: %t Please update password ASAP.",
			createdUser.Name,
			createdUser.AccessFinance,
			createdUser.AccessTimeregistration,
			createdUser.Administrator)
		utils.UpdateSession()
	},
}

func init() {
	rootCmd.AddCommand(UpdateUserCmd)

	UpdateUserCmd.Flags().StringVarP(&updateUserUsername, "username", "u", "", "New username. (required)")

	UpdateUserCmd.Flags().StringVarP(&updateUserPassword, "newPassword", "p", "", "New password. (required)")

	UpdateUserCmd.Flags().BoolVarP(&updateUserAccessFinance, "AccessFinance", "f", false, "Access to the financial database use true or false")

	UpdateUserCmd.Flags().BoolVarP(&updateUserAccessTime, "AccessTime", "t", false, "Access to the Time database use true or false")

	UpdateUserCmd.Flags().BoolVarP(&updateUserAdministrator, "administrator", "a", false, "Administrator rights for this person use true or false")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UpdateUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UpdateUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
