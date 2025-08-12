/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"time"

	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// LogoutCmd represents the Logout command
var LogoutCmd = &cobra.Command{
	Use:   "Logout",
	Short: "Loginout a session",
	Long:  `Used to logout from the system.`,
	Run: func(cmd *cobra.Command, args []string) {

		s := &utils.Session{
			UserID:   uuid.Nil,
			UserName: "",
			LastUsed: time.Now().Add(-20 * time.Minute),
		}

		err := utils.SaveSession(s)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(LogoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// LogoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// LogoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
