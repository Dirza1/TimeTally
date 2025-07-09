/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
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
		layout := "02-01-2006"
		query := utils.DatabaseConnection()
		switch overviewByCategoryType {
		case "Financial":
			entries, err := query.OverviewTransactionByCatagory(context.Background(), overviewByCategoryCategory)
			if err != nil {
				log.Fatal("error during record retrieval: ", err)
			}
			fmt.Printf("Overview of the Financial database of the catagroy %s\n", overviewByCategoryCategory)
			for _, entry := range entries {
				fmt.Printf("Entry ID: %s. Transaction date: %s. Category: %s, Description: %s, Total ammount(Euro): %.2f \n",
					entry.ID, entry.DateTransaction.Format(layout), entry.Catagory, entry.Description, entry.Amount)
			}
		case "Time":
			entries, err := query.OverviewTimeByCatagory(context.Background(), overviewByCategoryCategory)
			if err != nil {
				log.Fatal("error during record retrieval: ", err)
			}
			fmt.Printf("Overview of the Time database of the catagroy %s\n", overviewByCategoryCategory)
			for _, entry := range entries {
				fmt.Printf("Entry ID: %s. Activity date: %s. Category: %s, Description: %s, Time spent(Hours): %.2f \n",
					entry.ID, entry.DateActivity.Format(layout), entry.Catagory, entry.Description, entry.TimeHours)
			}
		default:
			fmt.Println("Incorrect use of the -t/ --Time flag. Use Finance or Time after the flag. Be mindfull of capitalisation.")
		}

	},
}

func init() {
	rootCmd.AddCommand(overviewByCategoryCmd)

	overviewByCategoryCmd.Flags().StringVarP(&overviewByCategoryCategory, "category", "c", "", "A flag to specify the category you aare looking for")
	overviewByCategoryCmd.MarkFlagRequired("category")

	overviewByCategoryCmd.Flags().StringVarP(&overviewByCategoryType, "type", "t", "", "A flag to specify the database you want to querry. Use Financial, Time or all after the flag")
	overviewByCategoryCmd.MarkFlagRequired("type")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// overviewByCategoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// overviewByCategoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
