/*
Copyright © 2025 Jasper Olthof-Donker <jasper.olthof@xs4all.nl>
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

var loginUsername string
var loginPassword string

// LoginCmd represents the Login command
var LoginCmd = &cobra.Command{
	Use:   "Login",
	Short: "A command to login to a session",
	Long: `This command will create a session with the supplied user.
	This session is used to check permissions when executing commands`,
	Run: func(cmd *cobra.Command, args []string) {
		if loginUsername == "" {
			fmt.Println("-u or --username flag not set. Please set this flag")
			return
		}
		if loginPassword == "" {
			fmt.Println("-p or --password flag not set. Please set this flag")
			return
		}

		queries := utils.DatabaseConnection()
		user, err := queries.Login(context.Background(), loginUsername)
		if err != nil {
			fmt.Printf("\nError during retrieval of user from database. Err:\n%s\n", err)
			return
		}
		if !utils.CompairPaswordHash(loginPassword, user.HashedPassword) {
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
		fmt.Println("Login Successful")
	},
}

func init() {
	rootCmd.AddCommand(LoginCmd)

	LoginCmd.Flags().StringVarP(&loginUsername, "username", "u", "", "New username. (required)")

	LoginCmd.Flags().StringVarP(&loginPassword, "password", "p", "", "password. (required)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// LoginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// LoginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
