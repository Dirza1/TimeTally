/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/spf13/cobra"
)

// overviewDatesCmd represents the overviewDates command
var OverviewDatesType string
var OverviewDatesFirstDate string
var OverviewDatesSecondDate string

var overviewDatesCmd = &cobra.Command{
	Use:   "overviewDates",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		layout := "01Januray2025"
		firstDate, err := time.Parse(layout, OverviewDatesFirstDate)
		if err != nil {
			log.Fatal("Issue with parsing first date: ", err)
		}

		secondDate, err := time.Parse(layout, OverviewDatesFirstDate)
		if err != nil {
			log.Fatal("Issue with parsing second date: ", err)
		}
		money := database.OverviewTransactionsDateParams{}
		time := database.OverviewTimeDatesParams{}
		switch OverviewDatesType {
		case "Finance":
			money = database.OverviewTransactionsDateParams{
				DateTransaction:   firstDate,
				DateTransaction_2: secondDate,
			}
		case "Time":
			time = database.OverviewTimeDatesParams{
				DateActivity:   firstDate,
				DateActivity_2: secondDate,
			}
		case "All":
			money = database.OverviewTransactionsDateParams{
				DateTransaction:   firstDate,
				DateTransaction_2: secondDate,
			}
			time = database.OverviewTimeDatesParams{
				DateActivity:   firstDate,
				DateActivity_2: secondDate,
			}
		default:
			log.Fatal("Incorrect use of the type flag. Use Finance, Time or All. Pay mind to the capitalation.")
		}

		fmt.Println("overviewMonth called")
		fmt.Printf("%s, %s", money, time)
	},
}

func init() {
	rootCmd.AddCommand(overviewDatesCmd)

	overviewDatesCmd.Flags().StringVarP(&OverviewDatesType, "type", "t", "all", "Flag to specify the database to querry. Use Finance, Time or All after the flag")
	overviewDatesCmd.MarkFlagRequired("type")

	overviewDatesCmd.Flags().StringVarP(&OverviewDatesFirstDate, "First-Date", "f", "", "Flag to specify the first date to querry. Use full date notateion. e.g. 01 January 2025")
	overviewDatesCmd.MarkFlagRequired("month")

	overviewDatesCmd.Flags().StringVarP(&OverviewDatesSecondDate, "Second-Date", "s", "", "Flag to specify the second date to querry. Use full date notateion. e.g. 01 January 2025")
	overviewDatesCmd.MarkFlagRequired("year")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// overviewDatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// overviewDatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
