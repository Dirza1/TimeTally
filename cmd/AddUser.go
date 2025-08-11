/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/spf13/cobra"
)

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
		queries := utils.DatabaseConnection()
		currentUser, err := utils.LoadSession()
		if err != nil {
			fmt.Println("Error retrieving current user from session")
		}
		permissions, err := queries.GetUserPermissions(context.Background(), currentUser.UserName)
		if err != nil {
			fmt.Println("Error during retrieval of user permissions from database")
			return
		}
		if permissions.Administrator != true {
			fmt.Println("Current user is not an administrator")
			return
		}
		newUserName, err := cmd.Flags().GetString("username")
		if err != nil {
			fmt.Println("username flag error")
			return
		}
		_, err = queries.GetUserPermissions(context.Background(), newUserName)
		if err == nil {
			fmt.Println("Username already exists. Please use another user name or delete old user")
			return
		}
		newPassword, err := cmd.Flags().GetString("newPassword")
		if err != nil {
			fmt.Println("Password flag error")
			return
		}
		hashedPasword, err := utils.Hashpassword(newPassword)
		if err != nil {
			fmt.Println("Error during pasword hash")
			return
		}
		finance, err := cmd.Flags().GetBool("AccessFinance")
		if err != nil {
			fmt.Println("Error during finance flag retrieval")
			return
		}
		time, err := cmd.Flags().GetBool("AccessTime")
		if err != nil {
			fmt.Println("Error during time flag retrieval")
			return
		}
		newUser := database.AddUserParams{
			Name:                   newUserName,
			HashedPassword:         hashedPasword,
			AccessFinance:          finance,
			AccessTimeregistration: time,
		}
		createdUser, err := queries.AddUser(context.Background(), newUser)
		if err != nil {
			fmt.Println("Error duing user creation")
			return
		}
		fmt.Printf("New user created. ID: %s, Name: %s, Financial access: %t, Time access: %t. Please update password ASAP.", createdUser.ID,
			createdUser.Name,
			createdUser.AccessFinance,
			createdUser.AccessTimeregistration)
	},
}

func init() {
	rootCmd.AddCommand(AddUserCmd)
	AddUserCmd.Flags().StringP("username", "u", "", "New username. (required)")
	err := AddUserCmd.MarkFlagRequired("username")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	AddUserCmd.Flags().StringP("newPassword", "n", "", "New password. (required)")
	err = AddUserCmd.MarkFlagRequired("newPassword")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	AddUserCmd.Flags().BoolP("AccessFinance", "f", false, "Access to the financial database use true or false")
	err = AddUserCmd.MarkFlagRequired("AccessFinance")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	AddUserCmd.Flags().BoolP("AccessTime", "t", false, "Access to the Time database use true or false")
	err = AddUserCmd.MarkFlagRequired("AccessTime")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
