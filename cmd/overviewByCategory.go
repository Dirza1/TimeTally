/*
Copyright © 2025 Jasper Olthof-Donker <jasper.olthof@xs4all.nl>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/Dirza1/TimeTally/internal/utils"
	"github.com/spf13/cobra"
)

// overviewByCategoryCmd represents the overviewByCategory command
var overviewByCategoryCategory string
var overviewByCategoryType string

var overviewByCategoryCmd = &cobra.Command{
	Use:   "overviewByCategory",
	Short: "Overview of all entries related to a specific category",
	Long: `This command returns all entries of a spcific database that is registered under a specifc catagory.
	This comand requires two flags. One to specify the database to querry and one to specify the catagory being looked for.`,
	Run: func(cmd *cobra.Command, args []string) {
		if overviewByCategoryCategory == "" {
			fmt.Printf("\n-c --category flag not set. Please set this flag\n")
			return
		}
		if overviewByCategoryType == "" {
			fmt.Printf("\n-t --type flag not set. Please set this flag\n")
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

		layout := "02-01-2006"
		queries := utils.DatabaseConnection()
		currentUser, err := utils.LoadSession()
		if err != nil {
			fmt.Printf("\nError retrieving current user from session. Err:\n%s\n", err)
		}
		permissions, err := queries.GetUserPermissions(context.Background(), currentUser.UserName)
		if err != nil {
			fmt.Printf("\nError during retrieval of user permissions from database. Err:\n%s\n", err)
			return
		}
		switch overviewByCategoryType {
		case "Financial":
			if permissions.AccessFinance != true {
				fmt.Println("Current user is not allowed in the financial database")
				return
			}
			entries, err := queries.OverviewTransactionByCatagory(context.Background(), overviewByCategoryCategory)
			if err != nil {
				fmt.Printf("\nerror during fetching of data: %s \n", err)
				return
			}
			fmt.Printf("Overview of the Financial database of the catagroy %s\n", overviewByCategoryCategory)
			for _, entry := range entries {
				fmt.Printf("Entry ID: %s. Transaction date: %s. Category: %s, Description: %s, Total ammount(Euro): %.2f \n",
					entry.ID, entry.DateTransaction.Format(layout), entry.Catagory, entry.Description, entry.Amount)
			}
			utils.UpdateSession()
		case "Time":
			if permissions.AccessTimeregistration != true {
				fmt.Println("Current user is allowed in the time registration databse")
				return
			}
			entries, err := queries.OverviewTimeByCatagory(context.Background(), overviewByCategoryCategory)
			if err != nil {
				fmt.Printf("\nerror during fetching of data: %s \n", err)
				return
			}
			fmt.Printf("Overview of the Time database of the catagroy %s\n", overviewByCategoryCategory)
			for _, entry := range entries {
				fmt.Printf("Entry ID: %s. Activity date: %s. Category: %s, Description: %s, Time spent(Hours): %.2f \n",
					entry.ID, entry.DateActivity.Format(layout), entry.Catagory, entry.Description, entry.TimeHours)
			}
			utils.UpdateSession()
		default:
			fmt.Println("Incorrect use of the -t/ --Time flag. Use Finance or Time after the flag. Be mindfull of capitalisation.")
		}

	},
}

func init() {
	rootCmd.AddCommand(overviewByCategoryCmd)

	overviewByCategoryCmd.Flags().StringVarP(&overviewByCategoryCategory, "category", "c", "", "A flag to specify the category you aare looking for")

	overviewByCategoryCmd.Flags().StringVarP(&overviewByCategoryType, "type", "t", "", "A flag to specify the database you want to querry. Use Financial, Time or all after the flag")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// overviewByCategoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// overviewByCategoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
