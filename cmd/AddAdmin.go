/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// AddAdminCmd represents the AddAdmin command
var AddAdminCmd = &cobra.Command{
	Use:   "AddAdmin",
	Short: "Add a user with administrator priveligaes",
	Long: `This command regesters a new user with administrator privelages. This is the highest level of access avalible.
	This level shouldd be restricted to as few people as possible to avoid mistakes.
	
	This level can:
	- Add, Update and remove entries from both databases.
	- Reset both databases.
	- Generate overviews of both databases`,
	Run: func(cmd *cobra.Command, args []string) {
		confirm, err := cmd.Flags().GetString("confirm")
		if err != nil || confirm != "true" {
			fmt.Println("Confirm flag not set correctly")
			return
		}
		err = godotenv.Load("../.env")
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
		currentUser, err := utils.LoadSession()
		if err != nil {
			fmt.Println("Error retrieving current user from session")
		}

		if currentUser.Administrator != true {
			fmt.Println("Current user is not an administrator")
			return
		}
		newUserName, err := cmd.Flags().GetString("username")
		if err != nil {
			fmt.Println("Password flag error")
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
		newAdmin := database.AddAdminParams{
			Name:           newUserName,
			HashedPassword: hashedPasword,
		}
		created, err := queries.AddAdmin(context.Background(), newAdmin)
		if err != nil {
			fmt.Println("Error creating a new user")
			return
		}
		fmt.Printf("\n New Administrator created. ID: %s, Name: %s", created.ID, created.Name)
	},
}

func init() {
	rootCmd.AddCommand(AddAdminCmd)

	AddAdminCmd.Flags().StringP("confirm", "c", "", "Used to confirm the database delition. Type true after the flag (required)")
	err := AddAdminCmd.MarkFlagRequired("confirm")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	AddAdminCmd.Flags().StringP("password", "p", "", "Additional password required for delition. (required)")
	err = AddAdminCmd.MarkFlagRequired("password")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	AddAdminCmd.Flags().StringP("username", "u", "", "New username. (required)")
	err = AddAdminCmd.MarkFlagRequired("username")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	AddAdminCmd.Flags().StringP("newPassword", "n", "", "New password. (required)")
	err = AddAdminCmd.MarkFlagRequired("newPassword")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddAdminCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddAdminCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
