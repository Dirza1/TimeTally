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

var newUserUsername string
var newUserPassword string
var newUserAccessFinance bool
var newUserAccessTime bool

// AddUserCmd represents the AddUser command
var AddUserCmd = &cobra.Command{
	Use:   "AddUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if newUserUsername == "" {
			fmt.Println("-u or --username flag not set. Please set this flag.")
			return
		}
		if newUserPassword == "" {
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

		_, err = queries.GetUserPermissions(context.Background(), newUserUsername)
		if err == nil {
			fmt.Println("Username already exists. Please use another user name or delete old user")
			return
		}

		hashedPasword, err := utils.Hashpassword(newUserPassword)
		if err != nil {
			fmt.Printf("\nError during pasword hash. Err:\n%s\n", err)
			return
		}

		newUser := database.AddUserParams{
			Name:                   newUserUsername,
			HashedPassword:         hashedPasword,
			AccessFinance:          newUserAccessFinance,
			AccessTimeregistration: newUserAccessTime,
		}
		createdUser, err := queries.AddUser(context.Background(), newUser)
		if err != nil {
			fmt.Printf("\nError duing user creation. Err:\n%s\n", err)
			return
		}
		fmt.Printf("New user created. ID: %s, Name: %s, Financial access: %t, Time access: %t. Please update password ASAP.", createdUser.ID,
			createdUser.Name,
			createdUser.AccessFinance,
			createdUser.AccessTimeregistration)
		utils.UpdateSession()
	},
}

func init() {
	rootCmd.AddCommand(AddUserCmd)

	AddUserCmd.Flags().StringVarP(&newUserUsername, "username", "u", "", "New username. (required)")

	AddUserCmd.Flags().StringVarP(&newUserPassword, "newPassword", "p", "", "New password. (required)")

	AddUserCmd.Flags().BoolVarP(&newUserAccessFinance, "AccessFinance", "f", false, "Access to the financial database use true or false")

	AddUserCmd.Flags().BoolVarP(&newUserAccessTime, "AccessTime", "t", false, "Access to the Time database use true or false")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
