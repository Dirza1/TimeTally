/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/spf13/cobra"
)

// overviewCmd represents the overview command
var OverviewType string

var overviewCmd = &cobra.Command{
	Use:   "overview",
	Short: "Returns an overview of all entries in the financial or timeregistration database",
	Long: `This command returns all the entries in either the Financial or the Timeregistration database.
	It returns the full database from the start of recording till the end. There are other commands avalible to narrow down the search for particular years or months.
	This command requires one flag to diferantiate between the diferent databases that can be querried.`,
	Run: func(cmd *cobra.Command, args []string) {
		if OverviewType == "" {
			fmt.Printf("\n-t --type flag not set. Please set this flag.\n")
			return
		}
		session, err := utils.LoadSession()
		if err != nil {
			fmt.Printf("\nError loading session. Err:\n%s\n", err)
		}
		currentTime := time.Now()
		if currentTime.Sub(session.LastUsed) > 15*time.Minute {
			fmt.Println("Users session expired. Please use the login command to continue using the system")
		}

		layout := "02-01-2006"
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
		switch OverviewType {
		case "Finance":
			if permissions.AccessFinance != true {
				fmt.Println("Current user is not allowed in the financial database")
				return
			}
			fmt.Println("Overfiew of the Financial database:")
			entries, err := queries.OverviewAllTransactions(context.Background())
			if err != nil {
				fmt.Printf("\nerror during fetching of data: %s \n", err)
				return
			}
			for _, entry := range entries {
				fmt.Printf("Entry ID: %s. Transaction date: %s. Category: %s, Description: %s, Total ammount(Euro): %.2f \n",
					entry.ID, entry.DateTransaction.Format(layout), entry.Catagory, entry.Description, entry.Amount)
			}
			utils.UpdateSession()
		case "Time":
			if permissions.AccessTimeregistration != true {
				fmt.Println("Current user is not allowed in the time registration databse")
				return
			}
			fmt.Println("Overview of the Timeregistrations:")
			entries, err := queries.OverviewAllTime(context.Background())
			if err != nil {
				fmt.Printf("\nerror during fetching of data: %s \n", err)
				return
			}
			for _, entry := range entries {
				fmt.Printf("Entry ID: %s. Activity date: %s. Category: %s, Description: %s, Time spent(Hours): %.2f \n",
					entry.ID, entry.DateActivity.Format(layout), entry.Catagory, entry.Description, entry.TimeHours)
			}
			utils.UpdateSession()
		default:
			fmt.Println("Incorrect use of the -t/ --Type flag. Use Finance or Time after the flag. Be mindfull of capitalisation.")
		}
	},
}

func init() {
	rootCmd.AddCommand(overviewCmd)

	overviewCmd.Flags().StringVarP(&OverviewType, "type", "t", "", "A flag to diferatiate between the databases. Use either Financial or Time after the flag")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// overviewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// overviewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
