/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/spf13/cobra"
)

// registerTimeCmd represents the registerTime command
var registerTimeDate string
var registerTimeMinutes int32
var registerTimeCategory string
var registerTimeDescription string

var registerTimeCmd = &cobra.Command{
	Use:   "registerTime",
	Short: "The command to register time spent",
	Long: `This command will allow you to register time spent.
	It will require the date of the activity, the time spent in minutes and what it was spent on.
	Later this entry is modifiable and deletable.`,
	Run: func(cmd *cobra.Command, args []string) {

		time := database.AddTimeRegistrationParams{
			DateActivity:  utils.TimeParse(registerTimeDate),
			LengthMinutes: registerTimeMinutes,
			Description:   registerTimeDescription,
			Catagory:      registerTimeCategory,
		}
		query := utils.DatabaseConnection()
		entry, err := query.AddTimeRegistration(context.Background(), time)
		if err != nil {
			fmt.Printf("error during inserting data into the database: %s \n", err)
			return
		}
		fmt.Printf("Databse entry created!")
		fmt.Println(entry)

	},
}

func init() {
	rootCmd.AddCommand(registerTimeCmd)

	registerTimeCmd.Flags().StringVarP(&registerTimeDate, "date", "d", "", "Flag to specify the date worked on a project. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")
	registerTimeCmd.MarkFlagRequired("date")

	registerTimeCmd.Flags().Int32VarP(&registerTimeMinutes, "time", "t", 0, "Flag to specify the amount of time worked on a project in minutes.")
	registerTimeCmd.MarkFlagRequired("time")

	registerTimeCmd.Flags().StringVarP(&registerTimeCategory, "category", "c", "", "Flag to specify the category/project name of the project.")
	registerTimeCmd.MarkFlagRequired("category")

	registerTimeCmd.Flags().StringVarP(&registerTimeDescription, "description", "e", "", "Flag to specify the description of the work performed.")
	registerTimeCmd.MarkFlagRequired("description")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerTimeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerTimeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
