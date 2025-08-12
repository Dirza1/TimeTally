/*
Copyright © 2025 Jasper Olthof-Donker <jasper.olthof@xs4all.nl>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/spf13/cobra"
)

// overviewDatesCmd represents the overviewDates command
var OverviewDatesType string
var OverviewDatesFirstDate string
var OverviewDatesSecondDate string

var overviewDatesCmd = &cobra.Command{
	Use:   "overviewDates",
	Short: "Gives an overview of either or both of the databases between two specified dates.",
	Long: `This command returns an overview of one or both of the databaases entries between the two specified dates.
	One flag sets the datase to be querries and the other two flags specify the dates to querry.`,
	Run: func(cmd *cobra.Command, args []string) {
		if OverviewDatesType == "" {
			fmt.Println("-t or --type flag not set. Please use this flag")
			return
		}
		if OverviewDatesFirstDate == "" {
			fmt.Println("-f or --first-date flag not set. Please use this flag")
			return
		}
		if OverviewDatesSecondDate == "" {
			fmt.Println("-s or --second-date flag not set. Please use this flag")
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

		firstDate := utils.TimeParse(OverviewDatesFirstDate)
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

		secondDate := utils.TimeParse(OverviewDatesSecondDate)

		money := database.OverviewTransactionsDateParams{}
		time := database.OverviewTimeDatesParams{}
		switch OverviewDatesType {
		case "Finance":
			if permissions.AccessFinance != true {
				fmt.Println("Current user is not allowed in the financial database")
				return
			}
			money = database.OverviewTransactionsDateParams{
				DateTransaction:   firstDate,
				DateTransaction_2: secondDate,
			}
			fmt.Println("Overview op financial databse:")
			fmt.Println(queries.OverviewTransactionsDate(context.Background(), money))
			utils.UpdateSession()
		case "Time":
			if permissions.AccessTimeregistration != true {
				fmt.Println("Current user is not allowed in the time registration databse")
				return
			}
			time = database.OverviewTimeDatesParams{
				DateActivity:   firstDate,
				DateActivity_2: secondDate,
			}
			fmt.Println("Overview op timeregistration databse:")
			fmt.Println(queries.OverviewTimeDates(context.Background(), time))
			utils.UpdateSession()
		case "All":
			if permissions.AccessTimeregistration != true || permissions.AccessFinance != true {
				fmt.Println("Current user is missing either time or financial access.")
				return
			}
			money = database.OverviewTransactionsDateParams{
				DateTransaction:   firstDate,
				DateTransaction_2: secondDate,
			}
			fmt.Println("Overview op financial databse:")
			fmt.Println(queries.OverviewTransactionsDate(context.Background(), money))
			time = database.OverviewTimeDatesParams{
				DateActivity:   firstDate,
				DateActivity_2: secondDate,
			}
			fmt.Println("Overview op timeregistration databse:")
			fmt.Println(queries.OverviewTimeDates(context.Background(), time))
			utils.UpdateSession()
		default:
			fmt.Println("Incorrect use of the type flag. Use Finance, Time or All. Pay mind to the capitalation.")
		}
	},
}

func init() {
	rootCmd.AddCommand(overviewDatesCmd)

	overviewDatesCmd.Flags().StringVarP(&OverviewDatesType, "type", "t", "all", "Flag to specify the database to querry. Use Finance, Time or All after the flag")

	overviewDatesCmd.Flags().StringVarP(&OverviewDatesFirstDate, "first-date", "f", "", "Flag to specify the first date to querry. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")

	overviewDatesCmd.Flags().StringVarP(&OverviewDatesSecondDate, "second-date", "s", "", "Flag to specify the second date to querry. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// overviewDatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// overviewDatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
