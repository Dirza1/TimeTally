/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/spf13/cobra"
)

// LoginCmd represents the Login command
var LoginCmd = &cobra.Command{
	Use:   "Login",
	Short: "A command to login to a session",
	Long: `This command will create a session with the suppyed user.
	This session is used to check permissions when executing commands`,
	Run: func(cmd *cobra.Command, args []string) {
		userName, err := cmd.Flags().GetString("username")
		if err != nil {
			fmt.Println("Username flag error")
			return
		}
		password, err := cmd.Flags().GetString("newPassword")
		if err != nil {
			fmt.Println("Password flag error")
			return
		}
		queries := utils.DatabaseConnection()
		user, err := queries.Login(context.Background(), userName)
		if err != nil {
			fmt.Println("Error during retrieval of user from database")
			return
		}
		if !utils.CompairPaswordHash(password, user.HashedPassword) {
			fmt.Println("Incorrect password supplied")
			return
		}
		s := &utils.Session{
			UserID:   user.ID,
			UserName: user.Name,
			LastUsed: time.Now(),
		}

		err = utils.SaveSession(s)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(LoginCmd)

	LoginCmd.Flags().StringP("username", "u", "", "New username. (required)")
	err := LoginCmd.MarkFlagRequired("username")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	LoginCmd.Flags().StringP("newPassword", "n", "", "New password. (required)")
	err = LoginCmd.MarkFlagRequired("newPassword")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// LoginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// LoginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
