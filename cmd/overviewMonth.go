/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/spf13/cobra"
)

// overviewMonthCmd represents the overviewMonth command
var OverviewMonthType string
var OverviewMonthMonth string
var OverviewMonthYear string

var overviewMonthCmd = &cobra.Command{
	Use:   "overviewMonth",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := utils.DatabaseConnection()

		switch OverviewMonthType {
		case "Finance":
			money := database.OverviewTransactionsMonthParams{
				DateTransaction: OverviewMonthMonth,
			}
		case "Time":

		case "All":

		default:
			log.Fatal("Incorrect use of the type flag. Use Finance, Time or All. Pay mind to the capitalation.")
		}

		fmt.Println("overviewMonth called")
	},
}

func init() {
	rootCmd.AddCommand(overviewMonthCmd)

	overviewCmd.Flags().StringVarP(&OverviewMonthType, "type", "t", "all", "Flag to specify the database to querry. Use Finance, Time or All after the flag")
	overviewCmd.MarkFlagRequired("type")

	overviewCmd.Flags().StringVarP(&OverviewMonthMonth, "month", "m", "", "Flag to specify the month to querry. Use mmm notation after the flag")
	overviewCmd.MarkFlagRequired("month")

	overviewCmd.Flags().StringVarP(&OverviewMonthYear, "year", "y", "", "Flag to specify the year to querry. Use mmm notation after the flag")
	overviewCmd.MarkFlagRequired("year")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// overviewMonthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// overviewMonthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
