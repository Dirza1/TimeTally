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

// FirstAdminCmd represents the FirstAdmin command
var FirstAdminCmd = &cobra.Command{
	Use:   "FirstAdmin",
	Short: "This command adds an administrator if none exist",
	Long: `During first instalation or if no admin exists within the database this command makes a new administrator.
	This command only works when no administrator exists yet in the database.
	Always ensure there is a backup administrator as this command does not work if any administrator accounts exist within the database`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := utils.DatabaseConnection()
		administrators, err := queries.CheckOnAdministartor(context.Background())
		if err != nil {
			fmt.Printf("\nError fetching list of current administrators:\n %s\n", err)
			return
		}
		if len(administrators) != 0 {
			fmt.Println("There are administrators in the system. Please ask them to generate accounts as required")
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
		newAdmin := database.CreateFirstAdministartorParams{
			Name:           newUserName,
			HashedPassword: hashedPasword,
		}
		created, err := queries.CreateFirstAdministartor(context.Background(), newAdmin)
		if err != nil {
			fmt.Printf("\nError creating a new user:\n%s\n", err)
			return
		}
		fmt.Printf("\n New Administrator created. ID: %s, Name: %s. Ensure admin changes their password ASAP!", created.ID, created.Name)
	},
}

func init() {
	rootCmd.AddCommand(FirstAdminCmd)

	FirstAdminCmd.Flags().StringP("username", "u", "", "New username. (required)")
	err := FirstAdminCmd.MarkFlagRequired("username")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	FirstAdminCmd.Flags().StringP("newPassword", "n", "", "New password. (required)")
	err = FirstAdminCmd.MarkFlagRequired("newPassword")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// FirstAdminCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// FirstAdminCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
