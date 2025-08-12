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

var addadminUsername string
var addadminPassword string

// AddAdminCmd represents the AddAdmin command
var AddAdminCmd = &cobra.Command{
	Use:   "AddAdmin",
	Short: "Add a user with administrator priveligaes",
	Long: `This command regesters a new user with administrator privelages. This is the highest level of access avalible.
	This level shouldd be restricted to as few people as possible to avoid mistakes.
	
	
	This level can:
	- Add, Update and remove entries from both databases.
	- Reset both databases.
	- Generate overviews of both databases
	- Create new users`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := utils.DatabaseConnection()
		currentUser, err := utils.LoadSession()
		if err != nil {
			fmt.Printf("\nError retrieving current user from session. Err: \n%s\n", err)
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
		_, err = queries.GetUserPermissions(context.Background(), addadminUsername)
		if err == nil {

			fmt.Printf("\n User already exists. Please use a diferent username\n")
			return
		}

		hashedPasword, err := utils.Hashpassword(addadminPassword)
		if err != nil {
			fmt.Printf("\nError during pasword hash. Err: \n%s\n", err)
			return
		}
		newAdmin := database.AddAdminParams{
			Name:           addadminUsername,
			HashedPassword: hashedPasword,
		}
		created, err := queries.AddAdmin(context.Background(), newAdmin)
		if err != nil {
			fmt.Printf("\nError creating a new user. Err: \n%s\n", err)
			return
		}
		fmt.Printf("\n New Administrator created. ID: %s, Name: %s. Ensure admin changes their password ASAP!\n", created.ID, created.Name)
	},
}

func init() {
	rootCmd.AddCommand(AddAdminCmd)

	AddAdminCmd.Flags().StringVarP(&addadminUsername, "username", "u", "", "New username. (required)")

	AddAdminCmd.Flags().StringVarP(&addadminPassword, "newPassword", "n", "", "New password. (required)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddAdminCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddAdminCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
