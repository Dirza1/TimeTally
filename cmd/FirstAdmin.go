/*
Copyright © 2025 Jasper Olthof-Donker <jasper.olthof@xs4all.nl>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/Dirza1/TimeTally/internal/database"
	"github.com/Dirza1/TimeTally/internal/utils"
	"github.com/spf13/cobra"
)

var firstAdminUsername string
var firstAdminPassword string

// FirstAdminCmd represents the FirstAdmin command
var FirstAdminCmd = &cobra.Command{
	Use:   "FirstAdmin",
	Short: "This command adds an administrator if none exist",
	Long: `During first instalation or if no admin exists within the database this command makes a new administrator.
	This command only works when no administrator exists yet in the database.
	Always ensure there is a backup administrator as this command does not work if any administrator accounts exist within the database.
	Should no active administrators be avalible manual intervention in the database is required. AVOID THIS AT ALL COSTS!`,
	Run: func(cmd *cobra.Command, args []string) {
		if firstAdminUsername == "" {
			fmt.Println("-u or --username flag not set. Please set this flag")
		}
		if firstAdminPassword == "" {
			fmt.Println("-p or --newPassword flag not set. Please set this flag")
		}

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

		_, err = queries.GetUserPermissions(context.Background(), firstAdminUsername)
		if err == nil {
			fmt.Println("Username already exists. Please use another user name or delete old user")
			return
		}

		hashedPasword, err := utils.Hashpassword(firstAdminPassword)
		if err != nil {
			fmt.Printf("\nError during pasword hash. Err:\n%s\n", err)
			return
		}
		newAdmin := database.CreateFirstAdministartorParams{
			Name:           firstAdminUsername,
			HashedPassword: hashedPasword,
		}
		created, err := queries.CreateFirstAdministartor(context.Background(), newAdmin)
		if err != nil {
			fmt.Printf("\nError creating a new user:\n%s\n", err)
			return
		}
		fmt.Printf("\n New Administrator created. ID: %s, Name: %s. Ensure admin changes their password ASAP!\n", created.ID, created.Name)
	},
}

func init() {
	rootCmd.AddCommand(FirstAdminCmd)

	FirstAdminCmd.Flags().StringVarP(&firstAdminUsername, "username", "u", "", "New username. (required)")

	FirstAdminCmd.Flags().StringVarP(&firstAdminPassword, "newPassword", "p", "", "New password. (required)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// FirstAdminCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// FirstAdminCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
